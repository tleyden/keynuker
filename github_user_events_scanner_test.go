// Licensed to the Apache Software Foundation (ASF) under one or more contributor license agreements;
// and to You under the Apache License, Version 2.0.  See LICENSE in project root for full license + copyright.

package keynuker

import (
	"encoding/json"
	"log"
	"os"
	"testing"
	"time"

	"context"
	"fmt"

	"io/ioutil"
	"net/url"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/google/go-github/github"
	"github.com/stretchr/testify/assert"
	"github.com/tleyden/keynuker/keynuker-go-common"
	"gopkg.in/h2non/gock.v1"
)

var fakePushEventRawPayload = `{
      "push_id":2031009949,
      "size":1,
      "distinct_size":1,
      "ref":"refs/heads/patch-1",
      "head":"148f8d11f8f50ed4c43026a009176c00dd521ada",
      "before":"64c38a12bfba4aefcb14b60fc9cd196f4f31634e",
      "commits":[
         {
            "sha":"148f8d11f8f50ed4c43026a009176c00dd521ada",
            "author":{
               "email":"testuser@gmail.com",
               "name":"Test User"
            },
            "message":"Fixes #743 Panic when calling github.Event.ParsePayload() with nil RawPayload (artificial mock scenario)",
            "distinct":true,
            "url":"https://api.github.com/repos/testuser/go-github/commits/148f8d11f8f50ed4c43026a009176c00dd521ada"
         }
      ]
   }`

func TestScanGithubUserEventsForAwsKeys(t *testing.T) {

	// Unless integration tests are enabled, use mock fetcher
	useMockFetcher := !IntegrationTestsEnabled()

	accessToken := "github_access_token"

	leakedKey := "FakeAccessKey"

	accessKeyMetadata := []FetchedAwsAccessKey{
		{
			AccessKeyId: aws.String(leakedKey),
			UserName:    aws.String("fakeuser@test.com"),
		},
	}

	githubUser := &github.User{
		Login: aws.String("tleyden"),
	}
	githubUserNoEvents := &github.User{
		Login: aws.String("nobody"),
	}
	githubUserOnlyEventsBeforeCheckpoint := &github.User{
		Login: aws.String("userOnlyEventsBeforeCheckpoint"),
	}
	githubUsers := []*github.User{
		githubUser,
		githubUserNoEvents,
		githubUserOnlyEventsBeforeCheckpoint,
	}

	// Make a fake checkpoint event that is 24 hours ago
	githubCheckpointEventCreatedAt := time.Now().Add(time.Hour * -24)
	githubCheckpointEvent := &github.Event{
		CreatedAt: aws.Time(githubCheckpointEventCreatedAt),
		ID:        aws.String("mockGithubEventBeforeCheckpoint3"),
	}
	githubEventCheckpoints := GithubEventCheckpoints{}
	githubEventCheckpoints[*githubUser.Login] = githubCheckpointEvent
	githubEventCheckpoints[*githubUserNoEvents.Login] = githubCheckpointEvent
	githubEventCheckpoints[*githubUserOnlyEventsBeforeCheckpoint.Login] = githubCheckpointEvent

	var fetcher GithubUserEventFetcher
	var mockFetcher *GithubUserEventFetcherMock
	var liveGithubFetcher *GoGithubUserEventFetcher

	var (
		mockGithubEvent1                 *github.Event
		mockGithubEvent2                 *github.Event
		mockGithubEvent3                 *github.Event
		mockGithubEventBeforeCheckpoint  *github.Event
		mockGithubEventBeforeCheckpoint2 *github.Event
		mockGithubEventBeforeCheckpoint3 *github.Event
	)

	switch useMockFetcher {
	case true:
		// Create mock user event fetcher
		mockFetcher = NewGithubUserEventFetcherMock()

		for _, mockGithubUser := range githubUsers {
			switch mockGithubUser {
			case githubUser:
				// Tee up a response to call to fetcher.FetchUserEvents() that returns a single event
				expectedFetchUserEventsInput := FetchUserEventsInput{
					Username:            *githubUser.Login,
					SinceEventTimestamp: githubCheckpointEvent.CreatedAt,
					CheckpointID:        *githubCheckpointEvent.ID,
				}
				rawPayload := json.RawMessage(fakePushEventRawPayload)
				mockGithubEvent1 = &github.Event{
					Type:       aws.String("PushEvent"),
					ID:         aws.String("mockGithubEvent1"),
					CreatedAt:  aws.Time(time.Now().Add(time.Minute * -5)),
					RawPayload: &rawPayload, // not used except to satisfy go-github parser
				}
				mockGithubEvent2 = &github.Event{
					Type:       aws.String("PushEvent"),
					ID:         aws.String("mockGithubEvent2"),
					CreatedAt:  aws.Time(time.Now().Add(time.Minute * -10)),
					RawPayload: &rawPayload, // not used except to satisfy go-github parser
				}
				mockGithubEvent3 = &github.Event{ // this event will be ignored since it's 24 hours before than checkpoint
					Type:       aws.String("PushEvent"),
					ID:         aws.String("mockGithubEvent3"),
					CreatedAt:  aws.Time(githubCheckpointEventCreatedAt.Add(time.Hour * -24)),
					RawPayload: &rawPayload, // not used except to satisfy go-github parser
				}
				mockFetcher.On("FetchUserEvents", context.Background(), expectedFetchUserEventsInput).Return(
					[]*github.Event{
						mockGithubEvent1,
						mockGithubEvent2,
						mockGithubEvent3,
					},
					nil, // no error
				)

				// Tee up a response to fetcher.ScanDownstreamContent which returns some content
				// 	ScanDownstreamContent(ctx context.Context, userEvent *github.Event) (content []byte, err error)
				mockFetcher.On("ScanDownstreamContent", context.Background(), mockGithubEvent1, accessKeyMetadata).Return(
					accessKeyMetadata,
					nil, // no error
				)
				mockFetcher.On("ScanDownstreamContent", context.Background(), mockGithubEvent2, accessKeyMetadata).Return(
					[]FetchedAwsAccessKey{},
					nil, // no error
				)
			case githubUserNoEvents:
				expectedFetchUserEventsInput := FetchUserEventsInput{
					Username:            *githubUserNoEvents.Login,
					SinceEventTimestamp: githubCheckpointEvent.CreatedAt,
					CheckpointID:        *githubCheckpointEvent.ID,
				}
				mockFetcher.On("FetchUserEvents", context.Background(), expectedFetchUserEventsInput).Return(
					[]*github.Event{},
					nil, // no error
				)
			case githubUserOnlyEventsBeforeCheckpoint:
				expectedFetchUserEventsInput := FetchUserEventsInput{
					Username:            *githubUserOnlyEventsBeforeCheckpoint.Login,
					SinceEventTimestamp: githubCheckpointEvent.CreatedAt,
					CheckpointID:        *githubCheckpointEvent.ID,
				}
				rawPayload := json.RawMessage(fakePushEventRawPayload)
				mockGithubEventBeforeCheckpoint = &github.Event{
					Type:       aws.String("PushEvent"),
					ID:         aws.String("mockGithubEventBeforeCheckpoint"),
					CreatedAt:  aws.Time(githubCheckpointEventCreatedAt.Add(time.Hour * -24)),
					RawPayload: &rawPayload, // not used except to satisfy go-github parser
				}
				mockGithubEventBeforeCheckpoint2 = &github.Event{
					Type:       aws.String("PushEvent"),
					ID:         aws.String("mockGithubEventBeforeCheckpoint2"),
					CreatedAt:  aws.Time(githubCheckpointEventCreatedAt.Add(time.Hour * -12)),
					RawPayload: &rawPayload, // not used except to satisfy go-github parser
				}
				mockGithubEventBeforeCheckpoint3 = &github.Event{ // simulate the case where one of the events is the checkpoint event
					Type:       aws.String("PushEvent"),
					ID:         githubCheckpointEvent.ID, // this github event has same ID and CreatedAt as the checkpoint
					CreatedAt:  aws.Time(githubCheckpointEventCreatedAt),
					RawPayload: &rawPayload, // not used except to satisfy go-github parser
				}
				mockFetcher.On("FetchUserEvents", context.Background(), expectedFetchUserEventsInput).Return(
					[]*github.Event{
						mockGithubEventBeforeCheckpoint,
						mockGithubEventBeforeCheckpoint2,
						mockGithubEventBeforeCheckpoint3,
					},
					nil, // no error
				)

			}
		}

		fetcher = mockFetcher

	default:

		// Don't use a mock fetcher, connect to live github API instead
		var ok bool
		accessToken, ok = os.LookupEnv(keynuker_go_common.EnvVarKeyNukerTestGithubAccessToken)
		if !ok {
			t.Skip("You must define environment variable keynuker_test_gh_access_token to run this test")
		}

		liveGithubFetcher = NewGoGithubUserEventFetcher(accessToken, GetIntegrationGithubApiBaseUrl())

		fetcher = liveGithubFetcher
	}

	params := ParamsScanGithubUserEventsForAwsKeys{
		AccessKeyMetadata: accessKeyMetadata,
		GithubUsers:       githubUsers,
		GithubConnectionParams: GithubConnectionParams{
			GithubAccessToken: accessToken,
		},
		KeyNukerOrg:            "test",
		GithubEventCheckpoints: githubEventCheckpoints,
	}

	// Create events scanner and run
	scanner := NewGithubUserEventsScanner(fetcher)
	docWrapper, err := scanner.ScanAwsKeys(params)

	if useMockFetcher {

		// Make assertions on the result
		assert.True(t, err == nil)
		assert.True(t, len(docWrapper.LeakedKeyEvents) == 1)
		assert.Equal(t, *docWrapper.LeakedKeyEvents[0].GithubEvent.ID, *mockGithubEvent1.ID)
		assert.Equal(t, "testuser@gmail.com", docWrapper.LeakedKeyEvents[0].LeakerEmail)
		assert.True(t, len(docWrapper.GithubEventCheckpoints) == 3)
		assert.Equal(t, *mockGithubEvent1.ID, *docWrapper.GithubEventCheckpoints[*githubUser.Login].ID)

		// The user with actual events should have a non-nil checkpoint
		assert.NotNil(t, docWrapper.GithubEventCheckpoints[*githubUser.Login])

		// The user with no events should get it's original checkpoint back
		assert.NotNil(t, docWrapper.GithubEventCheckpoints[*githubUserNoEvents.Login])
		assert.Equal(t, *docWrapper.GithubEventCheckpoints[*githubUserNoEvents.Login].ID, *githubCheckpointEvent.ID)

		// The user with only events older than the checkpoint should also have a non-nil checkpoint
		// corresponding to the most recent event that was scanned and skipped
		checkpoint := docWrapper.GithubEventCheckpoints[*githubUserOnlyEventsBeforeCheckpoint.Login]
		assert.NotNil(t, checkpoint)

		// Three events were returned for the githubUserOnlyEventsBeforeCheckpoint -- the one that was
		// recorded as the checkpoint should be the most recent from the three events (and happens to be
		// identical to the checkpointed event).  The three events are:
		// - An event that has the same created_at (and ID) as the checkpoint <-- this should be the new checkpoint
		// - A recent event that was 12 hours before checkpoint (36 hours ago)
		// - Older event was 24 hours before checkpoint (48 hours ago)
		assert.True(t, checkpoint.CreatedAt.Equal(githubCheckpointEventCreatedAt))
		assert.True(t, *checkpoint.ID == *githubCheckpointEvent.ID)

		// There should only be two calls to ScanDownstreamContent -- this should catch cases where
		// the event that is older than the checkpoint erroneously triggers a call to ScanDownstreamContent
		mockFetcher.AssertNumberOfCalls(t, "ScanDownstreamContent", 2)

	}

	// Emit the result
	docWrapperBytes, err := json.MarshalIndent(docWrapper, "", "    ")
	assert.True(t, err == nil)
	log.Printf("docWrapperBytes: %v", string(docWrapperBytes))

}

// Regression test against mock for reprducing https://github.com/tleyden/keynuker/issues/6
// TODO: This test isn't finished.  The mock needs to be extended to handle the scanning of additional commits
// TODO: See the end-to-end-integration test for examples.
func TestScanGithubLargePushEvents(t *testing.T) {

	// ------------------------------------ Create Event Fetcher -------------------------------------------------------

	var ok bool
	accessToken, ok := os.LookupEnv(keynuker_go_common.EnvVarKeyNukerTestGithubAccessToken)
	if !ok {
		t.Skip("You must define environment variable keynuker_test_gh_access_token to run this test")
	}

	// Create user event fetcher
	fetcher := NewGoGithubUserEventFetcher(accessToken, "")

	githubUser := &github.User{
		Login: aws.String("tleyden"),
	}

	// Make a fake checkpoint event that has the current timestamp
	githubCheckpointEvent := &github.Event{
		CreatedAt: aws.Time(time.Now().Add(time.Hour * -24)),
		ID:        aws.String("FakeEventID"),
	}
	githubEventCheckpoints := GithubEventCheckpoints{}
	githubEventCheckpoints[*githubUser.Login] = githubCheckpointEvent

	leakedKey := "FakeAccessKey"

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

	payload, err := largePushEvent.ParsePayload()
	if err != nil {
		t.Fatalf("Unable to parse payload.  Err: %v", err)
	}
	pushEvent := payload.(*github.PushEvent)

	events := []*github.Event{
		largePushEvent,
	}

	gock.New("https://api.github.com").
		Get(fmt.Sprintf("/users/%s/events", *githubUser.Login)).
		MatchParam("per_page", "100").
		Reply(200).
		JSON(events)

	for i, commit := range pushEvent.Commits {
		commitUrl, _ := url.Parse(commit.GetURL())
		schemeAndHost := fmt.Sprintf("%s://%s", commitUrl.Scheme, commitUrl.Host)
		gock.New(schemeAndHost).
			Get(commitUrl.Path).
			Reply(200).
			JSON(map[string]string{
				"content": fmt.Sprintf("commit %d", i),
			},
			)
	}

	// ------------------------------------ Invoke Scanner -------------------------------------------------------------

	params := ParamsScanGithubUserEventsForAwsKeys{
		AccessKeyMetadata: []FetchedAwsAccessKey{
			{
				AccessKeyId: aws.String(leakedKey),
				UserName:    aws.String("fakeuser@test.com"),
			},
		},
		GithubUsers: []*github.User{
			githubUser,
		},
		GithubConnectionParams: GithubConnectionParams{
			GithubAccessToken: "github_access_token",
		},
		KeyNukerOrg:            "test",
		GithubEventCheckpoints: githubEventCheckpoints,
	}

	// Create events scanner and run
	scanner := NewGithubUserEventsScanner(fetcher)
	docWrapper, err := scanner.ScanAwsKeys(params)

	if err != nil {
		t.Fatalf("Error calling ScanAwsKeys(): %v", err)
	}

	log.Printf("doc result: %v err: %v", docWrapper, err)

}

func TestCreateFetchUserEventsInputNilCheckpoint(t *testing.T) {

	testUsername := "testuser"
	p := ParamsScanGithubUserEventsForAwsKeys{
		GithubEventCheckpoints: GithubEventCheckpoints{
			testUsername: nil,
		},
	}

	user := &github.User{}
	user.Login = aws.String(testUsername)
	fetchUserEventsInput := p.CreateFetchUserEventsInput(user)
	assert.EqualValues(t, fetchUserEventsInput.Username, testUsername)
	assert.True(t, fetchUserEventsInput.SinceEventTimestamp.Before(time.Now()))

}

func TestCreateFetchUserEventsInputNormalCheckpoint(t *testing.T) {

	testUsername := "testuser"
	now := time.Now()
	p := ParamsScanGithubUserEventsForAwsKeys{
		GithubEventCheckpoints: GithubEventCheckpoints{
			testUsername: &github.Event{
				CreatedAt: aws.Time(now),
				ID:        aws.String("FakeEventID"),
			},
		},
	}

	user := &github.User{}
	user.Login = aws.String(testUsername)
	fetchUserEventsInput := p.CreateFetchUserEventsInput(user)
	assert.EqualValues(t, fetchUserEventsInput.Username, testUsername)
	assert.True(t, fetchUserEventsInput.SinceEventTimestamp.Equal(now))

}

func TestSetDefaultCheckpointsForMissing(t *testing.T) {

	testUsername := "testuser"
	params := ParamsScanGithubUserEventsForAwsKeys{
		GithubUsers: []*github.User{
			{
				Login: aws.String(testUsername),
			},
		},
		GithubEventCheckpoints: GithubEventCheckpoints{
			testUsername: nil,
		},
	}
	paramsWithDefaultCheckpoints := params.SetDefaultCheckpointsForMissing(keynuker_go_common.DefaultCheckpointEventTimeWindow)
	checkpointForTestUser := paramsWithDefaultCheckpoints.GithubEventCheckpoints[testUsername]
	assert.True(t, checkpointForTestUser.CreatedAt.Before(time.Now()))

}
