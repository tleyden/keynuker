// Licensed to the Apache Software Foundation (ASF) under one or more contributor license agreements;
// and to You under the Apache License, Version 2.0.  See LICENSE in project root for full license + copyright.

package keynuker

import (
	"context"
	"io/ioutil"
	"net/http"
	"time"

	"bytes"

	"log"

	"io"
	"strings"

	"github.com/google/go-github/github"
	"github.com/tleyden/keynuker/keynuker-go-common"
)

//go:generate goautomock -template=testify -o "github_user_event_fetcher_mock.go" GithubUserEventFetcher

// Abstract the calls to the github API in an interface for dependency injection / mocking purposes
type GithubUserEventFetcher interface {

	// Given a github username and filtering parameters, fetch events from the user event stream
	FetchUserEvents(ctx context.Context, fetchUserEventsInput FetchUserEventsInput) ([]*github.Event, error)

	// Given a specific github event (eg, a commit), get the actual content for that event to be scanned for aws keys
	FetchDownstreamContent(ctx context.Context, userEvent *github.Event) (content []byte, err error)
}

type  GoGithubUserEventFetcher struct {
	*GithubClientWrapper
}

// Input parameters for the github user event fetcher, which include filtering params such as checkpoint filtering
type FetchUserEventsInput struct {

	// The github username
	Username string

	// For checkpointing purposes, only consider events that are _after_ this timestamp
	SinceEventTimestamp *time.Time
}

func NewGoGithubUserEventFetcher(accessToken string) *GoGithubUserEventFetcher {
	return &GoGithubUserEventFetcher{
		GithubClientWrapper: NewGithubClientWrapper(accessToken),
	}
}

func eventCreatedAtBefore(event *github.Event, sinceEventTimestamp *time.Time) bool {
	return (*event.CreatedAt).Before(*sinceEventTimestamp)
}

func (guef GoGithubUserEventFetcher) FetchUserEvents(ctx context.Context, fetchUserEventsInput FetchUserEventsInput) ([]*github.Event, error) {

	publicOnly := false
	curApiResultPage := 0
	events := []*github.Event{}

	// Loop over all pages returned by API and accumulate events
	// TODO: #1 needs to also collect github gists
	// TODO: #2 filter out any events that aren't in EventTypesToInclude
	// TODO: #3 should skip any events that are before SinceEventId and SinceEventTimestamp
	// TODO: #4 the code to loop over all pages could be extracted out into a re-usable wrappr function
	// TODO: #5 what should publicOnly be set to?  Seems like it depends on the permissions of the accesstoken
	// TODO:    and it would be good to grab private events if the access token gives enough permissions

	for {

		listOptions := &github.ListOptions{
			PerPage: MaxPerPage,
			Page:    curApiResultPage,
		}

		eventsPerPage, response, err := guef.ApiClient.Activity.ListEventsPerformedByUser(
			ctx,
			fetchUserEventsInput.Username,
			publicOnly,
			listOptions,
		)

		if err != nil {
			return nil, err
		}

		// Create a logger that will only log at most 5 messages about older checkpoints, to prevent spamming logs
		boundedLogger := keynuker_go_common.CreateBoundedLogger(5)

		// Loop over events and append to result
		for _, event := range eventsPerPage {

			// If the event is older than our checkpoint, skip it.  This filtering also happens in GithubUserEventsScanner.scanAwsKeysForUser()
			if fetchUserEventsInput.SinceEventTimestamp != nil && eventCreatedAtBefore(event, fetchUserEventsInput.SinceEventTimestamp) {
				msg := "Event %v created_at %v, but that is older than checkpoint at %v, so skipping.  User: %v"
				boundedLogger.Printf(msg, *event.ID, *event.CreatedAt, *fetchUserEventsInput.SinceEventTimestamp, fetchUserEventsInput.Username)
				continue
			}
			events = append(events, event)

		}

		if response.NextPage <= curApiResultPage {
			// Last page, we're done
			break
		}

		curApiResultPage += 1

	}

	return events, nil

}

func (guef GoGithubUserEventFetcher) FetchDownstreamContent(ctx context.Context, userEvent *github.Event) (content []byte, err error) {

	payload, err := userEvent.ParsePayload()
	if err != nil {
		return nil, err
	}

	switch v := payload.(type) {
	case *github.PullRequestEvent:

		content, err := guef.FetchUrlContent(ctx, v.PullRequest.GetPatchURL())
		if err != nil {
			return nil, err
		}
		return content, nil

	case *github.PushEvent:

		buffer := bytes.Buffer{}

		maxCommitsPerPushEvent := 20

		if *v.Size > maxCommitsPerPushEvent {
			log.Printf("WARNING: PushEvent %v has > 20 commits, but only 20 commtis will be scanned.", *v.PushID)
		}

		if strings.Contains(*v.Ref, keynuker_go_common.KeyNukerIntegrationTestBranch) {
			// skip this since as an experiment
			log.Printf("Skipping push event %v since it's on %v testing branch", *v.PushID, keynuker_go_common.KeyNukerIntegrationTestBranch)
			return []byte(""), nil
		}

		commits := v.Commits
		for _, commit := range commits {
			log.Printf("Getting content for commit: %v url: %v", *commit.SHA, commit.GetURL())
			content, err := guef.FetchUrlContent(ctx, commit.GetURL())
			if err != nil {
				return nil, err
			}
			buffer.Write(content)
		}

		// If more than 20 commits for this PushEvent, fetch content for the remaining commits.
		// Example PushEvent w/ more than 20 commits: https://gist.github.com/tleyden/68d972b02b2b9306fa6e2eb26310b751
		if *v.Size > maxCommitsPerPushEvent {

			// Fetch the rest of the commits for this push event and append downstream content to buffer
			err := guef.FetchCommitsForPushEvent(ctx, userEvent, v, &buffer)
			if err != nil {
				return nil, err
			}

		}

		return buffer.Bytes(), nil

	default:
		// TODO: add case statements for all events that should be scanned
		return []byte(*userEvent.RawPayload), nil

	}

	return nil, nil
}

// Since PushEvents only contain 20 commits max, this fetches the remaining commits and writes the content to the
// writer passed in.  For example, pushEvent.Size might indicate that there were 100 commits in the push events,
// and so the remaining 80 commits will need to be scanned.
//
// Github API: https://developer.github.com/v3/repos/commits/
func (guef GoGithubUserEventFetcher) FetchCommitsForPushEvent(ctx context.Context, userEvent *github.Event, pushEvent *github.PushEvent, w io.Writer) error {

	numCommitsToScan := *pushEvent.Size
	numCommitsScanned := 0
	resultPage := 0

	// The inline commits in the push event don't need to be re-scanned, so build a map of their sha's
	inlineCommits := map[string]struct{}{}
	for _, commit := range pushEvent.Commits {
		inlineCommits[*commit.SHA] = struct{}{}
	}

	// Keep listing commits on that branch until we go back far enough to reach the last commit in pushEvent.size

	for {

		if numCommitsScanned >= numCommitsToScan {
			// done
			return nil
		}

		commitListOptions := &github.CommitsListOptions{
			SHA: *pushEvent.Head,
			ListOptions: github.ListOptions{
				PerPage: MaxPerPage,
				Page:    resultPage,
			},
		}
		log.Printf("Listing additional commits: %+v", commitListOptions)

		// "tleyden/keynuker" -> ["tleyden", "keynuker"] -> "keynuker"
		repoNameAndUsername := *userEvent.Repo.Name
		repoNameAndUsernameComponents := strings.Split(repoNameAndUsername, "/")
		repoName := repoNameAndUsernameComponents[1]

		additionalCommits, resp, err := guef.ApiClient.Repositories.ListCommits(
			ctx,
			*userEvent.Actor.Login,
			repoName,
			commitListOptions,
		)
		if err != nil {
			return err
		}

		for _, additionalCommit := range additionalCommits {

			_, foundInlineCommit := inlineCommits[*additionalCommit.SHA]
			if foundInlineCommit {
				numCommitsScanned += 1
				continue
			}

			log.Printf("Getting content for additional commit: %v url: %v", *additionalCommit.SHA, additionalCommit.GetURL())
			content, err := guef.FetchUrlContent(ctx, additionalCommit.GetURL())
			if err != nil {
				return err
			}
			w.Write(content)

			numCommitsScanned += 1

			if numCommitsScanned >= numCommitsToScan {
				// done
				return nil
			}

		}

		log.Printf("resp %+v.  NextPage: %v.  LastPage: %v", resp, resp.NextPage, resp.LastPage)

		resultPage = resp.NextPage

	}

	return nil
}

func (guef GoGithubUserEventFetcher) FetchUrlContent(ctx context.Context, url string) (content []byte, err error) {

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	q := req.URL.Query()
	q.Add("access_token", guef.AccessToken)
	req.URL.RawQuery = q.Encode()

	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	resp_body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return resp_body, nil

}
