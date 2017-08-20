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

	"github.com/google/go-github/github"
	"github.com/tleyden/keynuker/keynuker-go-common"
	"strings"
)

//go:generate goautomock -template=testify -o "github_user_event_fetcher_mock.go" GithubUserEventFetcher

type GithubUserEventFetcher interface {
	FetchUserEvents(ctx context.Context, fetchUserEventsInput FetchUserEventsInput) ([]*github.Event, error)
	FetchDownstreamContent(ctx context.Context, userEvent *github.Event) (content []byte, err error)
}

type GoGithubUserEventFetcher struct {
	*GithubClientWrapper
}

type FetchUserEventsInput struct {
	Username string

	EventTypesToInclude []string

	SinceEventId uint64

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

		// Loop over events and append to result
		for _, event := range eventsPerPage {

			// If the event is older than our checkpoint, skip it
			if fetchUserEventsInput.SinceEventTimestamp != nil && eventCreatedAtBefore(event, fetchUserEventsInput.SinceEventTimestamp) {
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

		if *v.Size > 20 {
			log.Printf("WARNING: PushEvent %v has > 20 commits, but only 20 commtis will be scanned.", *v.PushID)
		}

		if strings.Contains(*v.Ref, keynuker_go_common.KeyNukerIntegrationTestBranch) {
			// skip this since as an experiment
			log.Printf("Skipping push event %v on %v branch", *v.PushID, keynuker_go_common.KeyNukerIntegrationTestBranch)
			return []byte(""), nil
		}

		// TODO: If there are 20 commits in this push, there is a good chance there are more commits that didn't
		// TODO: make it in due to limitations mentioned in https://developer.github.com/v3/activity/events/types/#pushevent
		// TODO: and so there needs to be an enhancement to fetch and scan these commits
		// TODO: example PushEvent w/ more than 20 commtis: https://gist.github.com/tleyden/68d972b02b2b9306fa6e2eb26310b751
		commits := v.Commits
		for _, commit := range commits {
			log.Printf("Getting content for commit: %+v url: %v", commit, commit.GetURL())
			content, err := guef.FetchUrlContent(ctx, commit.GetURL())
			if err != nil {
				return nil, err
			}
			buffer.Write(content)
		}

		return buffer.Bytes(), nil

	default:
		// TODO: add case statements for all events that should be scanned
		return []byte(*userEvent.RawPayload), nil

	}

	return nil, nil
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
