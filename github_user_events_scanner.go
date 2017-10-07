// Licensed to the Apache Software Foundation (ASF) under one or more contributor license agreements;
// and to You under the Apache License, Version 2.0.  See LICENSE in project root for full license + copyright.

package keynuker

import (
	"context"
	"fmt"
	"time"

	"log"
	"sync"

	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/google/go-github/github"
	"github.com/tleyden/keynuker/keynuker-go-common"
)

type GithubUserEventsScanner struct {

	// The github user events scanner uses an underlying fetcher, which can be easily mocked out for testing
	fetcher GithubUserEventFetcher
}

// The result of scanning a user's github events
type ScanResult struct {
	CheckpointEvent *github.Event // Latest event scanned
	User            *github.User
	LeakedKeyEvents []LeakedKeyEvent
	Error           error
}

func NewGithubUserEventsScanner(fetcher GithubUserEventFetcher) *GithubUserEventsScanner {

	return &GithubUserEventsScanner{
		fetcher: fetcher,
	}

}

func (gues GithubUserEventsScanner) ScanAwsKeys(params ParamsScanGithubUserEventsForAwsKeys) (docWrapper DocumentScanGithubUserEventsForAwsKeys, err error) {

	ctx := context.Background()

	// TODO: this code needs review.  It works, and it's relatively efficient, but it's unnecessarily complicated

	chUsersToProcess := make(chan *github.User)

	chScanResults := make(chan ScanResult)

	// Start goroutines for scanning github users
	processUsersDone := sync.WaitGroup{}
	processUsersDone.Add(50)
	for i := 0; i < 50; i++ {
		go func(id int) {
			for {
				user, ok := <-chUsersToProcess

				if !ok {
					// channel is closed, we're done
					processUsersDone.Done()
					return
				}
				scanResult, _ := gues.scanAwsKeysForUser(
					ctx,
					user,
					params,
				)
				chScanResults <- scanResult
			}
		}(i)
	}

	// Send users down the chUsersToProcess channel
	sentUsersWaitGroup := sync.WaitGroup{}
	sentUsersWaitGroup.Add(1)
	go func() {
		defer sentUsersWaitGroup.Done()
		for _, user := range params.GithubUsers {
			chUsersToProcess <- user
		}
	}()

	// Accumulate any leaked key events here
	leakedKeyEvents := []LeakedKeyEvent{}

	// Resulting checkpoints after processing
	githubEventCheckpoints := GithubEventCheckpoints{}

	collectedResultsWaitGroup := sync.WaitGroup{}
	collectedResultsWaitGroup.Add(1)
	go func() {
		defer collectedResultsWaitGroup.Done()
		for scanResult := range chScanResults {

			// TODO: partial errors are being absorbed/ignored here.  They should somehow be propagated back to the caller
			if scanResult.Error != nil {
				log.Printf("Warning: Got error trying to scan github user events: %+v", scanResult)
				continue
			}

			githubEventCheckpoints[*scanResult.User.Login] = scanResult.CheckpointEvent

			leakedKeyEvents = append(leakedKeyEvents, scanResult.LeakedKeyEvents...)
		}
	}()

	// Wait until all users have been sent down chUsersToProcess
	sentUsersWaitGroup.Wait()

	// Close chUsersToProcess to signal that no more users need to be processed
	close(chUsersToProcess)

	// Wait until all user processing goroutines are done
	processUsersDone.Wait()

	// Close chScanResults to signal no more results will be coming
	close(chScanResults)

	// Wait until all results are collected
	collectedResultsWaitGroup.Wait()

	// Create result doc
	doc := DocumentScanGithubUserEventsForAwsKeys{
		LeakedKeyEvents:        leakedKeyEvents,
		GithubEventCheckpoints: githubEventCheckpoints,
	}

	return doc, nil

}

func (gues GithubUserEventsScanner) scanAwsKeysForUser(ctx context.Context, user *github.User,
	params ParamsScanGithubUserEventsForAwsKeys) (scanResult ScanResult, err error) {

	log.Printf("ScanGithubUserEventsForAwsKeys for user: %v", *user.Login)

	scanResult.User = user

	fetchUserEventsInput := params.CreateFetchUserEventsInput(user)

	userEvents, err := gues.fetcher.FetchUserEvents(ctx, fetchUserEventsInput)
	if err != nil {
		msg := "Failed to fetch user events.  ParamsScanGithubUserEventsForAwsKeys: %+v Error: %v"
		scanResult.Error = fmt.Errorf(msg, fetchUserEventsInput, err)
		return scanResult, scanResult.Error
	}

	// Track the latest event processed during this scan
	var checkpointEvent *github.Event

	log.Printf("Scanning %d events for user: %v", len(userEvents), *user.Login)

	for _, userEvent := range userEvents {

		// Make sure that it's _after_ the stored checkpoint, otherwise skip it since it's already been scanned
		// NOTE: this code is a "double check" with the code in FetchUserEvents()
		if (*userEvent.CreatedAt).Before(*fetchUserEventsInput.SinceEventTimestamp) {
			msg := "Skipping event since after stored checkpoint. User: %v. Event id: %v  Event created at: %v Stored checkpoint: %v"
			log.Printf(msg, *user.Login, *userEvent.ID, *userEvent.CreatedAt, *fetchUserEventsInput.SinceEventTimestamp)
			continue
		}

		downstreamEventContent, err := gues.fetcher.FetchDownstreamContent(ctx, userEvent)
		if err != nil {
			msg := "Failed to fetch user event content.  Event: %+v Error: %v"
			scanResult.Error = fmt.Errorf(msg, userEvent, err)
			return scanResult, scanResult.Error
		}

		// Scan for leaked keys
		log.Printf("User: %v. Scanning %d bytes of content for event: %v", *user.Login, len(downstreamEventContent), *userEvent.ID)

		leakedKeys, nearbyContent, err := Scan(params.AccessKeyMetadata, downstreamEventContent)
		if err != nil {
			scanResult.Error = fmt.Errorf("Failed to scan event content.  Event: %+v Error: %v", userEvent, err)
			return scanResult, scanResult.Error
		}

		leakerEmail := ""
		if len(leakedKeys) > 0 {
			log.Printf("Found %d leaked keys in event: %v", len(leakedKeys), *userEvent.ID)
			leakerEmail, err = gues.discoverLeakerEmail(userEvent)
			if err != nil {
				log.Printf("Warning: error discovering leaker email: %v", err)
			}
		}

		// Create LeakedKeyEvents from leaked keys, append to result
		for _, leakedKey := range leakedKeys {
			leakedKeyEvent := LeakedKeyEvent{
				AccessKeyMetadata: leakedKey,
				GithubUser:        user,
				GithubEvent:       userEvent,
				NearbyContent:     nearbyContent,
				LeakerEmail:       leakerEmail,
			}

			scanResult.LeakedKeyEvents = append(scanResult.LeakedKeyEvents, leakedKeyEvent)
		}

		// Update checkpoint.  If there is no checkpoint event yet, set as current event
		if checkpointEvent == nil {
			checkpointEvent = userEvent
		} else {
			// Otherwise only set the checkpoint if current event happened after checkpoint
			if (*userEvent.CreatedAt).After(*checkpointEvent.CreatedAt) {
				checkpointEvent = userEvent
			}
		}
		scanResult.CheckpointEvent = checkpointEvent

	}

	return scanResult, nil

}

func (gues GithubUserEventsScanner) discoverLeakerEmail(userEvent *github.Event) (email string, err error) {

	payload, err := userEvent.ParsePayload()
	if err != nil {
		return "", err
	}

	switch v := payload.(type) {
	case *github.PushEvent:

		if v.Ref != nil && strings.Contains(*v.Ref, keynuker_go_common.KeyNukerIntegrationTestBranch) {
			// skip this since as an experiment
			log.Printf("Skipping push event %v on %v branch", *v.PushID, keynuker_go_common.KeyNukerIntegrationTestBranch)
			return "", nil
		}

		commits := v.Commits
		for _, commit := range commits {
			if commit.Author != nil && commit.Author.Email != nil && *commit.Author.Email != "" {
				return *commit.Author.Email, nil
			}
		}

		return "", nil

	default:
		return "", nil

	}

}

type ParamsScanGithubUserEventsForAwsKeys struct {

	// This is the name of the KeyNuker "org/tenant".  Defaults to "default", but allows to be extended multi-tenant.
	KeyNukerOrg string

	// The github access token, which needs "read:org" permissions in order to read the concealed "non-public"
	// members of the orgs
	GithubAccessToken string

	// A list of github users
	GithubUsers []*github.User

	// AWS access keys to scan for
	AccessKeyMetadata []FetchedAwsAccessKey

	// Track the latest event processed for each user in GithubUsers by keeping per-user checkpoints
	GithubEventCheckpoints GithubEventCheckpoints
}

func (p ParamsScanGithubUserEventsForAwsKeys) Validate() error {
	// Must have a github access token
	if p.GithubAccessToken == "" {
		return fmt.Errorf("You must supply the GithubAccessToken")
	}
	return nil
}

func (p ParamsScanGithubUserEventsForAwsKeys) CreateFetchUserEventsInput(user *github.User) FetchUserEventsInput {

	fetchUserEventsInput := FetchUserEventsInput{
		Username: *user.Login,
	}

	// If there is a stored checkpoint for that user, use it
	githubCheckpointEvent, ok := p.GithubEventCheckpoints[*user.Login]
	if ok {

		if githubCheckpointEvent == nil {
			fetchUserEventsInput.SinceEventTimestamp = aws.Time(time.Now().Add(keynuker_go_common.DefaultCheckpointEventTimeWindow))
		} else {
			fetchUserEventsInput.SinceEventTimestamp = githubCheckpointEvent.CreatedAt
		}
	}

	return fetchUserEventsInput
}

func (p ParamsScanGithubUserEventsForAwsKeys) WithDefaultKeynukerOrg() ParamsScanGithubUserEventsForAwsKeys {

	returnParams := p

	// If no keynuker org is specified, use "default"
	if returnParams.KeyNukerOrg == "" {
		returnParams.KeyNukerOrg = keynuker_go_common.DefaultKeyNukerOrg
	}

	return returnParams

}

// Update params to set the github event checkpoints to have a recent time window so that it doesn't scan every single
// user event (which can take too long) in the absence of actual stored checkpoints, which aren't fully working yet
// as of the time of this writing.
//
// Example recentTimeWindow: time.Duration(time.Hour * -12)
func (p ParamsScanGithubUserEventsForAwsKeys) SetDefaultCheckpointsForMissing(recentTimeWindow time.Duration) ParamsScanGithubUserEventsForAwsKeys {

	returnParams := p

	// If GithubEventCheckpoints is nil, create an empty map
	if returnParams.GithubEventCheckpoints == nil {
		returnParams.GithubEventCheckpoints = GithubEventCheckpoints{}
	}

	// For each github user that doesn't have a checkpoint event, create an artificial one
	// that is 24 hours old, so that only events from last 24 hours are processed, in order to
	// limit the amount of work this call ScanGithubUserEventsForAwsKeys() will perform.
	for _, githubUser := range returnParams.GithubUsers {

		checkpoint, ok := returnParams.GithubEventCheckpoints.CheckpointForUser(githubUser)
		if !ok || checkpoint == nil {
			githubCheckpointEvent := &github.Event{
				CreatedAt: aws.Time(time.Now().Add(recentTimeWindow)), // eg, time.Hour * -12
			}
			returnParams.GithubEventCheckpoints[*githubUser.Login] = githubCheckpointEvent
		}
	}

	return returnParams

}

type DocumentScanGithubUserEventsForAwsKeys struct {
	Id                     string `json:"_id"`
	LeakedKeyEvents        []LeakedKeyEvent
	GithubEventCheckpoints GithubEventCheckpoints
}

type DocumentWrapperScanGithubUserEventsForAwsKeys struct {
	// Serialize into a form that the cloudant db adapter expects
	Doc   DocumentScanGithubUserEventsForAwsKeys `json:"doc"`
	DocId string                                 `json:"docid"`
}
