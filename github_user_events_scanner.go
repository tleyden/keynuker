// Licensed to the Apache Software Foundation (ASF) under one or more contributor license agreements;
// and to You under the Apache License, Version 2.0.  See LICENSE in project root for full license + copyright.

package keynuker

import (
	"context"
	"time"

	"log"

	"strings"

	"strconv"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/google/go-github/github"
	"github.com/pkg/errors"
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

func (s *ScanResult) SetCheckpointIfMostRecent(user *github.User, latestEventScanned *github.Event) {

	if latestEventScanned == nil {
		// Ignore nil events
		return
	}

	// If there is no checkpoint event yet whatsoever, set as current event no matter what it is
	if s.CheckpointEvent == nil {
		log.Printf("SetCheckpointIfMostRecent setting checkpoint from nil -> to %v for user: %v", *latestEventScanned.ID, user)
		s.CheckpointEvent = latestEventScanned
	}

	// Otherwise only set the checkpoint if current event happened after or has a larger ID than current checkpoint
	if EventMoreRecent(s.CheckpointEvent, latestEventScanned) {
		log.Printf("SetCheckpointIfMostRecent setting checkpoint from %v -> to %v for user: %v", *s.CheckpointEvent.ID, *latestEventScanned.ID, user)
		s.CheckpointEvent = latestEventScanned
	}

}

// Is the incoming event more recent than the checkpoint event?
func EventMoreRecent(checkpoint, incoming *github.Event) bool {

	if checkpoint == nil || incoming == nil {
		return false
	}

	// Compare by date, if it's after, then we're done
	if (*incoming.CreatedAt).After(*checkpoint.CreatedAt) {
		return true
	}

	// If the date is equal, then fallback to comparing by the checkpoint ID
	if (*incoming.CreatedAt).Equal(*checkpoint.CreatedAt) {

		incomingID, err := strconv.Atoi(*incoming.ID)
		if err != nil {
			log.Printf("Warning: error converting event id to int.  Event id: %v.  Err: %v", *incoming.ID, err)
			return false
		}

		checkpointID, err := strconv.Atoi(*checkpoint.ID)
		if err != nil {
			log.Printf("Warning: error converting event id to int.  Event id: %v.  Err: %v", *checkpoint.ID, err)
			return false
		}

		return incomingID > checkpointID

	}

	return false

}

func (s *ScanResult) SetDefaultResultCheckpoint(user *github.User, checkpoints GithubEventCheckpoints) {

	checkpoint, ok := checkpoints.CheckpointForUser(user)
	if ok {
		log.Printf("SetDefaultResultCheckpoint to %v for user %v", *checkpoint.ID, user)
		s.SetCheckpointIfMostRecent(user, checkpoint)
	}

}

// Return a compact (stripped) version of the checkpoint event that has the minimal
// fields to still be useful
func (s ScanResult) CompactCheckpointEvent() *github.Event {

	if s.CheckpointEvent == nil {
		return nil
	}

	resultEvent := &github.Event{}
	resultEvent.CreatedAt = s.CheckpointEvent.CreatedAt
	resultEvent.ID = s.CheckpointEvent.ID
	resultEvent.Type = s.CheckpointEvent.Type

	return resultEvent

}

func NewGithubUserEventsScanner(fetcher GithubUserEventFetcher) *GithubUserEventsScanner {

	return &GithubUserEventsScanner{
		fetcher: fetcher,
	}

}

// For each github user, scan their event feed for leaked aws keys
// This purposely does not make concurrent requests due to: https://developer.github.com/guides/best-practices-for-integrators/#dealing-with-abuse-rate-limits
func (gues GithubUserEventsScanner) ScanAwsKeys(params ParamsScanGithubUserEventsForAwsKeys) (docWrapper DocumentScanGithubUserEventsForAwsKeys, err error) {

	ctx := context.Background()

	// Accumulate any leaked key events here
	leakedKeyEvents := []LeakedKeyEvent{}

	// Resulting checkpoints after processing.  Initialize them with a copy of the starting checkpoints.
	// Even if the process exits early, this must return _all_ of the checkpoints, since any missing checkpoints
	// will effectively erase them in the database (until a smarter checkpoint DB updater/merger is added)
	githubEventCheckpoints := params.CopyGithubEventCheckpoints()

	startTime := time.Now()

	for _, user := range params.GithubUsers {

		scanResult, err := gues.scanAwsKeysForUser(
			ctx,
			user,
			params,
		)

		if err != nil {
			log.Printf("Warning: scanAwsKeysForUser() returned err: %v", err)
		}

		// Update the checkpoint with the scanresult.  This is done regardless of whether or not there was an error,
		// since scanAwsKeysForUser() will only move it forward in the case there were no errors
		githubEventCheckpoints[*scanResult.User.Login] = scanResult.CompactCheckpointEvent()

		// ErrorInjection: only append leaks if no errors, then run test with temp github error on commit after finding a leak
		if len(scanResult.LeakedKeyEvents) > 0 {
			leakedKeyEvents = append(leakedKeyEvents, scanResult.LeakedKeyEvents...)
		}

		if scanResult.Error != nil {
			log.Printf("Warning: Got error trying to scan github user events: %+v", scanResult.Error)
		}

		// if we've gone over 270 seconds, return early and leave remaining for next polling loop
		if time.Since(startTime) > keynuker_go_common.HighWatermarkExecutionSeconds {
			log.Printf("Warning: over high watermark for action exuction time limit.  Returning current results so far")
			break
		}

	}

	// Create result doc
	doc := DocumentScanGithubUserEventsForAwsKeys{
		LeakedKeyEvents:        leakedKeyEvents,
		GithubEventCheckpoints: githubEventCheckpoints,
	}

	return doc, nil

}

func (gues GithubUserEventsScanner) scanAwsKeysForUser(ctx context.Context, user *github.User,
	params ParamsScanGithubUserEventsForAwsKeys) (scanResult ScanResult, err error) {

	scanResult.User = user

	// It's better to return the existing checkpoint rather than an empty checkpoint,
	// since an empty checkpoint will clobber what's in the database and cause it to revert to
	// a time-based checkpoint 24 hours ago.  So set the scanResult checkpoint to the current
	// checkpoint for that user.
	scanResult.SetDefaultResultCheckpoint(user, params.GithubEventCheckpoints)

	fetchUserEventsInput := params.CreateFetchUserEventsInput(user)

	log.Printf("scanAwsKeysForUser for user: %v.  fetchUserEventsInput: %+v", *user.Login, fetchUserEventsInput)

	userEvents, err := gues.fetcher.FetchUserEvents(ctx, fetchUserEventsInput)
	if err != nil {
		scanResult.Error = errors.Wrapf(err, "Failed to fetch user events.  ParamsScanGithubUserEventsForAwsKeys: %+v", fetchUserEventsInput)
		return scanResult, scanResult.Error
	}

	// Track the latest event processed during this scan
	// var checkpointEvent *github.Event

	log.Printf("Scanning %d events for user: %v", len(userEvents), *user.Login)

	// Create a logger that will only log at most 5 messages about older checkpoints, to prevent spamming the logs
	boundedLogger := keynuker_go_common.CreateBoundedLogger(5)

	for _, userEvent := range userEvents {

		// Make sure that it's _after_ the stored checkpoint, otherwise skip it since it's already been scanned
		// The reason it can't just check Before() || Equal() is that if two or more events had thesame timestamp,
		// and the tool polled the api and only get one of the events, then the others would be ignored by
		// the checkpointing algorithm
		if (*userEvent.CreatedAt).Before(*fetchUserEventsInput.SinceEventTimestamp) {
			msg := "Skipping event since before checkpoint date. " +
				"User: %v. Event id: %v  Event created at: %v Stored checkpoint: %v"
			boundedLogger.Printf(msg, *user.Login, *userEvent.ID, *userEvent.CreatedAt, *fetchUserEventsInput.SinceEventTimestamp)

			// Update the checkpoint, if it's the most recent event, despite the fact that we are skipping this event.
			// This covers the edge case where _all_ events returned by the user are older than the checkpoint -- in
			// that case it's still necessary to bump the checkpoint to the most recent event in that event set.
			scanResult.SetCheckpointIfMostRecent(user, userEvent)

			continue
		}

		// If the event has the exact same ID as the checkpoint event ID, then skip it since it's already been scanned.
		// Warning: before doing any numerical comparison on checkpoint event ID's, will have to address the fact
		// that it's currently storing checkpoints with ArtificialCheckPointID
		if fetchUserEventsInput.MatchesCheckpointID(userEvent) {
			msg := "Skipping event since it has the same event ID as the checkpoint. " +
				"User: %v. Event id: %v  Event created at: %v  Checkpoint timestamp: %v Checkpoint ID: %v"
			boundedLogger.Printf(msg, *user.Login, *userEvent.ID, *userEvent.CreatedAt,
				*fetchUserEventsInput.SinceEventTimestamp, fetchUserEventsInput.CheckpointID)

			scanResult.SetCheckpointIfMostRecent(user, userEvent)

			continue
		}

		// Logging
		msg := "Fetching downstream content for event. " +
			"User: %v. Event id: %v  Event created at: %v Stored checkpoint: %v Checkpoint ID: %v"
		log.Printf(msg, *user.Login, *userEvent.ID, *userEvent.CreatedAt,
			*fetchUserEventsInput.SinceEventTimestamp, fetchUserEventsInput.CheckpointID)

		// Logging
		msg = "Scanning event. User: %v. Event id: %v  Event created at: %v Stored checkpoint: %v Checkpoint ID: %v"
		log.Printf(msg, *user.Login, *userEvent.ID, *userEvent.CreatedAt,
			*fetchUserEventsInput.SinceEventTimestamp, fetchUserEventsInput.CheckpointID)

		leakedKeys, err := gues.fetcher.ScanDownstreamContent(ctx, userEvent, params.AccessKeyMetadata)

		// Before checking the error, update the scan result with any leaked keys.  It's possible that
		// it found leaks, and returned an error.
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
				LeakerEmail:       leakerEmail,
			}

			scanResult.LeakedKeyEvents = append(scanResult.LeakedKeyEvents, leakedKeyEvent)
		}

		// Check the error
		if err != nil {

			// If it's a rate limit error, treat this as temporary / retryable.  Abort the current
			// operation and return an error, which will prevent the checkpoint from advancing, which will cause a retry later.
			if keynuker_go_common.IsTemporaryGithubError(err) {
				scanResult.Error = errors.Wrapf(err, "WARNING: Failed to fetch user event content due to temporary error.  Event: %+v", userEvent)
				return scanResult, scanResult.Error
			} else {
				// Otherwise, treat this as a permanent error and log a warning and skip this event (which is bad, since now
				// that event's content will never be scanned)
				scanResult.SetCheckpointIfMostRecent(user, userEvent)
				log.Printf("WARNING: Failed to fetch user event content due to unexpected error.  Permanently skipping Event: %+v Error: %v", userEvent, err)
				continue
			}

		}

		// Update checkpoint.
		scanResult.SetCheckpointIfMostRecent(user, userEvent)

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

	// Github API URL and access token
	GithubConnectionParams

	// This is the name of the KeyNuker "org/tenant".  Defaults to "default", but allows to be extended multi-tenant.
	KeyNukerOrg string

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
		return errors.Errorf("You must supply the GithubAccessToken")
	}
	return nil
}

func (p ParamsScanGithubUserEventsForAwsKeys) CopyGithubEventCheckpoints() GithubEventCheckpoints {
	result := GithubEventCheckpoints{}
	for userLogin, startingCheckpoint := range p.GithubEventCheckpoints {
		checkpoint := *startingCheckpoint
		result[userLogin] = &checkpoint
	}
	return result
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
			if githubCheckpointEvent.ID != nil {
				fetchUserEventsInput.CheckpointID = *githubCheckpointEvent.ID
			}
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
				ID:        aws.String("ArtificialCheckpointId"),
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
