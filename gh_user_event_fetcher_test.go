// Licensed to the Apache Software Foundation (ASF) under one or more contributor license agreements;
// and to You under the Apache License, Version 2.0.  See LICENSE in project root for full license + copyright.

package keynuker

import (
	"context"
	"log"
	"os"
	"testing"

	"github.com/couchbaselabs/go.assert"
	"github.com/google/go-github/github"
	"github.com/tleyden/keynuker/keynuker-go-common"
	"gopkg.in/h2non/gock.v1"
	"fmt"
	"io/ioutil"
	"encoding/json"
	"time"
	"strconv"
)

// Not much of a unit test, just makes it easy to run ghUserEventFetcher.FetchUserEvents() by hand in isolation
func TestGithubUserEventFetcher(t *testing.T) {

	SkipIfIntegrationsTestsNotEnabled(t)

	accessToken, ok := os.LookupEnv(keynuker_go_common.EnvVarKeyNukerTestGithubAccessToken)
	if !ok {
		t.Skip("You must define environment variable keynuker_test_gh_access_token to run this test")
	}

	ctx := context.Background()

	ghUserEventFetcher := NewGoGithubUserEventFetcher(accessToken, GetIntegrationGithubApiBaseUrl())

	fetchUserEventsInput := FetchUserEventsInput{
		Username: "tahmmee",
	}
	userEvents, err := ghUserEventFetcher.FetchUserEvents(ctx, fetchUserEventsInput)
	assert.True(t, err == nil)

	assert.True(t, len(userEvents) > 0)
	log.Printf("User events: %v", userEvents)
	log.Printf("# events: %v", len(userEvents))

}

func TestGithubUserEventDownstreamContentFetcher(t *testing.T) {

	// Return 200 mock events
	// The checkpoint will be older than the oldest event, so that all events should be scanned
	// The first attempt, it will return a github temporary error on the 150th event.
	// At this point, it should return the checkpoint of the last event scanned before the error (149th event)
	// Then rerun it with the same mock results, but with no github temp error
	// Should scan all the content and return most recent event as the checkpoint

	ctx := context.Background()

	ghUserEventFetcher := NewGoGithubUserEventFetcher("mock_access_token", GetIntegrationGithubApiBaseUrl())

	checkpointTime := time.Now().Add(time.Hour * -24)
	fetchUserEventsInput := FetchUserEventsInput{
		Username: "mock_user",
		SinceEventTimestamp: &checkpointTime,
	}

	// ------------------------------------ Setup Gock HTTP mock -------------------------------------------------------

	defer gock.Off() // Flush pending mocks after test execution

	filename := "testdata/large_push_event.json"
	largePushEventData, err := ioutil.ReadFile(filename)
	if err != nil {
		t.Fatalf("Unable to read file: %v.  Err: %v", filename, err)
	}

	largePushEvent := &github.Event{}
	errUnmarshal := json.Unmarshal(largePushEventData, largePushEvent)
	if errUnmarshal != nil {
		t.Fatalf("Unable to unmarshal data: %v", errUnmarshal)
	}

	eventsPage1 := []*github.Event{}
	for i := 0; i < 100; i++ {
		eventPage1 := *largePushEvent
		createdAt := time.Now().Add(time.Minute * -1 * time.Duration(i))
		eventPage1.CreatedAt = &createdAt
		eventId, err := strconv.Atoi(*eventPage1.ID)
		if err != nil {
			t.Fatalf("Failed to generate event id: %v", err)
		}
		eventIdStr := fmt.Sprintf("%d", eventId - i)
		eventPage1.ID = &eventIdStr
		eventsPage1 = append(eventsPage1, &eventPage1)
	}

	eventsPage2 := []*github.Event{}

	for i := 100; i < 200; i++ {
		eventPage2 := *largePushEvent
		createdAt := time.Now().Add(time.Minute * -1 * time.Duration(i))
		eventPage2.CreatedAt = &createdAt
		eventId, err := strconv.Atoi(*eventPage2.ID)
		if err != nil {
			t.Fatalf("Failed to generate event id: %v", err)
		}
		eventIdStr := fmt.Sprintf("%d", eventId - i)
		eventPage2.ID = &eventIdStr
		eventsPage2 = append(eventsPage2, &eventPage2)
	}

	githubUrl := "https://api.github.com"
	eventsEndpoint := fmt.Sprintf("/users/%s/events", fetchUserEventsInput.Username)
	nextPage := fmt.Sprintf("<%s/%s?page=2&per_page=100>; rel=\"next\"", githubUrl, eventsEndpoint)

	gock.New(githubUrl).
		Get(eventsEndpoint).
		MatchParam("per_page", "100").
		Reply(200).
		AddHeader("Link", nextPage).
		JSON(eventsPage1)

	gock.New(githubUrl).
		Get(eventsEndpoint).
		MatchParam("per_page", "100").
		MatchParam("page", "1").
		Reply(200).
		JSON(eventsPage2)

	userEvents, err := ghUserEventFetcher.FetchUserEvents(ctx, fetchUserEventsInput)
	if err != nil {
		t.Fatalf("Failed to fetch events: %v", err)
	}

	log.Printf("returned %d user events", len(userEvents))

	for i, userEvent := range userEvents {
		log.Printf("userEvent #%d ID: %v createdAt: %v", i, *userEvent.ID, userEvent.CreatedAt)
	}

	// log.Printf("userEvents: %+v, err", userEvents, err )


}

// Run ghUserEventFetcher.FetchDownstreamContent() by hand in isolation
func RunGithubUserEventDownstreamContentFetcher() {

	accessToken, ok := os.LookupEnv(keynuker_go_common.EnvVarKeyNukerTestGithubAccessToken)
	if !ok {
		log.Printf("You must define environment variable keynuker_test_gh_access_token to run this test")
		return
	}

	ctx := context.Background()

	ghUserEventFetcher := NewGoGithubUserEventFetcher(accessToken, GetIntegrationGithubApiBaseUrl())

	fetchUserEventsInput := FetchUserEventsInput{
		Username: "tleyden",
	}
	userEvents, err := ghUserEventFetcher.FetchUserEvents(ctx, fetchUserEventsInput)

	userEvent := userEvents[0]
	log.Printf("userEvent: %+v", userEvent)

	downstreamEventContent, err := ghUserEventFetcher.FetchDownstreamContent(ctx, userEvent)
	if err != nil {
		log.Printf("error: %v", err)
	}

	log.Printf("downstreamEventContent: %+v", string(downstreamEventContent))

}

// Not much of a unit test, just makes it easy to run RepositoriesService.GetContents() by hand in isolation
func TestGithubGetContents(t *testing.T) {

	SkipIfIntegrationsTestsNotEnabled(t)

	accessToken, ok := os.LookupEnv(keynuker_go_common.EnvVarKeyNukerTestGithubAccessToken)
	if !ok {
		t.Skip("You must define environment variable keynuker_test_gh_access_token to run this test")
	}

	githubClientWrapper := NewGithubClientWrapper(accessToken, GetIntegrationGithubApiBaseUrl())
	opt := &github.RepositoryContentGetOptions{
		Ref: "master",
	}

	fileContent, dirContent, response, err := githubClientWrapper.ApiClient.Repositories.GetContents(
		context.Background(),
		"tleyden",
		"keynuker-playground",
		"KeyNukerEndToEndIntegrationTestLeakedKeyLargefile-7486bd90-bd3f-4a97-9066-94f7eb7329ca.txt",
		opt,
	)

	log.Printf("fileContent: %v", fileContent)
	log.Printf("dirContent: %v", dirContent)
	log.Printf("response: %v", response)
	log.Printf("err: %v", err)

}

// Not much of a unit test, just makes it easy to run GitService.GetBlob() by hand in isolation
func TestGithubGetBlob(t *testing.T) {

	SkipIfIntegrationsTestsNotEnabled(t)

	accessToken, ok := os.LookupEnv(keynuker_go_common.EnvVarKeyNukerTestGithubAccessToken)
	if !ok {
		t.Skip("You must define environment variable keynuker_test_gh_access_token to run this test")
	}

	sha := "cc9d0ebcaff7e33b0d08535b0393483bf70ea804"
	githubClientWrapper := NewGithubClientWrapper(accessToken, GetIntegrationGithubApiBaseUrl())
	blob, response, err := githubClientWrapper.ApiClient.Git.GetBlob(
		context.Background(),
		"tleyden",
		"keynuker-playground",
		sha,
	)

	log.Printf("blob: %v", blob)
	log.Printf("response: %v", response)
	log.Printf("err: %v", err)

}
