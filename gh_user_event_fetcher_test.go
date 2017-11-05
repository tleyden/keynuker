// Licensed to the Apache Software Foundation (ASF) under one or more contributor license agreements;
// and to You under the Apache License, Version 2.0.  See LICENSE in project root for full license + copyright.

package keynuker

import (
	"context"
	"log"
	"os"
	"testing"

	"encoding/json"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/couchbaselabs/go.assert"
	"github.com/google/go-github/github"
	"github.com/tleyden/keynuker/keynuker-go-common"
	"gopkg.in/h2non/gock.v1"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
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
		Username:            "mock_user",
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
		eventIdStr := fmt.Sprintf("%d", eventId-i)
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
		eventIdStr := fmt.Sprintf("%d", eventId-i)
		eventPage2.ID = &eventIdStr
		eventsPage2 = append(eventsPage2, &eventPage2)
	}

	githubUrl := "https://api.github.com"
	eventsEndpoint := fmt.Sprintf("/users/%s/events", fetchUserEventsInput.Username)
	nextPage := fmt.Sprintf("<%s/%s?page=2&per_page=100>; rel=\"next\", <%s/%s?page=2&per_page=100>; rel=\"last\"",
		githubUrl, eventsEndpoint, githubUrl, eventsEndpoint)

	gock.New(githubUrl).
		Get(eventsEndpoint).
		MatchParam("per_page", "100").
		Reply(200).
		AddHeader("Link", nextPage).
		JSON(eventsPage1)

	gock.New(githubUrl).
		Get(eventsEndpoint).
		MatchParam("per_page", "100").
		MatchParam("page", "2").
		Reply(200).
		JSON(eventsPage2)

	userEvents, err := ghUserEventFetcher.FetchUserEvents(ctx, fetchUserEventsInput)
	if err != nil {
		t.Fatalf("Failed to fetch events: %v", err)
	}

	log.Printf("returned %d user events", len(userEvents))
	assert.Equals(t, len(userEvents), 200)

	for i, userEvent := range userEvents {
		log.Printf("userEvent #%d ID: %v createdAt: %v", i, *userEvent.ID, userEvent.CreatedAt)
	}

	// These events should be returned in *reverse chronological order*.  Eg, userEvents[0] should have an older
	// CreatedAt value than userEvents[-1] (last event).  That's because we want get the oldest event that's after
	// the checkpoint and start scanning towards the most recent event.
	firstEvent := userEvents[0]
	lastEvent := userEvents[len(userEvents)-1]

	assert.True(t, (*firstEvent).CreatedAt.Before(*lastEvent.CreatedAt))

	// log.Printf("userEvents: %+v, err", userEvents, err )

}

func TestScanContentForCommits(t *testing.T) {

	defer gock.Off() // Flush pending mocks after test execution

	ctx := context.Background()

	leakedKey := "FakeAccessKey"
	accessKeysToScan := []FetchedAwsAccessKey{
		{
			AccessKeyId: aws.String(leakedKey),
			UserName:    aws.String("fakeuser@test.com"),
		},
	}

	githubUrl := "https://api.github.com"

	ghUserEventFetcher := NewGoGithubUserEventFetcher("mock_access_token", GetIntegrationGithubApiBaseUrl())

	// Grab the fields needed from github.ErrorResponse.  Can't use whole struct since
	// it ends up nil'ing out the Response field, which doesn't have an `omitempty` specifier
	githubAbuseError := struct {
		DocumentationURL string `json:"documentation_url,omitempty"`
		Message          string `json:"message"` // error message
	}{
		DocumentationURL: "https://developer.github.com/v3#abuse-rate-limits",
		Message:          "abuse detection",
	}

	commit1 := &github.RepositoryCommit{
		SHA: aws.String("sha1"),
		URL: aws.String("url1"),
		Files: []github.CommitFile{
			{
				Filename: aws.String("file1a"),
				Patch:    aws.String(fmt.Sprintf("leaked key: %v", leakedKey)),
			},
			{
				Filename: aws.String("file1b"),
				Patch:    aws.String(fmt.Sprintf("nothingburger")),
			},
		},
	}

	// This commit simulates a github temporary error
	commit2TempError := &github.RepositoryCommit{
		SHA: aws.String("sha2"),
		URL: aws.String("url2"),
		Files: []github.CommitFile{
			{
				Filename: aws.String("file2a"),
				Patch:    aws.String(fmt.Sprintf("nothingburger")),
			},
			{
				Filename: aws.String("file2b"),
				Patch:    aws.String(fmt.Sprintf("nothingburger")),
			},
		},
	}

	commits := []*github.RepositoryCommit{
		commit1,
		commit2TempError,
	}

	gock.New(githubUrl).
		Get("/repos/username/reponame/commits/sha1").
		Reply(200).
		JSON(commit1)

	gock.New(githubUrl).
		Get("/repos/username/reponame/commits/sha2").
		Reply(http.StatusForbidden).
		JSON(githubAbuseError)

	leaks, err := ghUserEventFetcher.ScanContentForCommits(
		ctx,
		"username",
		"reponame",
		ConvertRepositoryCommits(commits),
		accessKeysToScan,
	)

	assert.True(t, len(leaks) == 1)

	// Expect a github temporary error
	assert.True(t, err != nil)
	assert.True(t, keynuker_go_common.IsTemporaryGithubError(err))

}

// Run ghUserEventFetcher.ScanDownstreamContent() by hand in isolation
func TestRunGithubUserEventDownstreamContentFetcher(t *testing.T) {

	SkipIfIntegrationsTestsNotEnabled(t)

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

	for i, userEvent := range userEvents {
		log.Printf("userEvent #%d ID: %v createdAt: %v", i, *userEvent.ID, userEvent.CreatedAt)
	}

	leakedKeys, err := ghUserEventFetcher.ScanDownstreamContent(ctx, userEvent, []FetchedAwsAccessKey{})
	if err != nil {
		log.Printf("error: %v", err)
	}

	log.Printf("leakedKeys: %+v", leakedKeys)

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
