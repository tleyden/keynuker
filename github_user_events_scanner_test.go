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

	"github.com/aws/aws-sdk-go/aws"
	"github.com/google/go-github/github"
	"github.com/stretchr/testify/assert"
	"github.com/tleyden/keynuker/keynuker-go-common"
	"gopkg.in/h2non/gock.v1"
)

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
	githubUsers := []*github.User{
		githubUser,
		githubUserNoEvents,
	}

	// Make a fake checkpoint event that has the current timestamp
	githubCheckpointEvent := &github.Event{
		CreatedAt: aws.Time(time.Now().Add(time.Hour * -24)),
	}
	githubEventCheckpoints := GithubEventCheckpoints{}
	githubEventCheckpoints[*githubUser.Login] = githubCheckpointEvent
	githubEventCheckpoints[*githubUserNoEvents.Login] = githubCheckpointEvent

	var fetcher GithubUserEventFetcher
	var mockFetcher *GithubUserEventFetcherMock
	var liveGithubFetcher *GoGithubUserEventFetcher

	var mockGithubEvent1 *github.Event
	var mockGithubEvent2 *github.Event

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
				}
				mockGithubEvent1 = &github.Event{
					Type:      aws.String("PushEvent"),
					ID:        aws.String("mockGithubEvent1"),
					CreatedAt: aws.Time(time.Now()),
				}
				mockGithubEvent2 = &github.Event{
					Type:      aws.String("PushEvent"),
					ID:        aws.String("mockGithubEvent2"),
					CreatedAt: aws.Time(time.Now()),
				}
				mockFetcher.On("FetchUserEvents", context.Background(), expectedFetchUserEventsInput).Return(
					[]*github.Event{
						mockGithubEvent1,
						mockGithubEvent2,
					},
					nil, // no error
				)

				// Tee up a response to fetcher.FetchDownstreamContent which returns some content
				// 	FetchDownstreamContent(ctx context.Context, userEvent *github.Event) (content []byte, err error)
				mockFetcher.On("FetchDownstreamContent", context.Background(), mockGithubEvent1).Return(
					[]byte(fmt.Sprintf("Fake content intermixed w/ %s, which is a leaked key", leakedKey)),
					nil, // no error
				)
				mockFetcher.On("FetchDownstreamContent", context.Background(), mockGithubEvent2).Return(
					[]byte(fmt.Sprintf("Fake content sans leaked key")),
					nil, // no error
				)
			case githubUserNoEvents:
				expectedFetchUserEventsInput := FetchUserEventsInput{
					Username:            *githubUserNoEvents.Login,
					SinceEventTimestamp: githubCheckpointEvent.CreatedAt,
				}
				mockFetcher.On("FetchUserEvents", context.Background(), expectedFetchUserEventsInput).Return(
					[]*github.Event{},
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

		liveGithubFetcher = NewGoGithubUserEventFetcher(accessToken)

		fetcher = liveGithubFetcher
	}

	params := ParamsScanGithubUserEventsForAwsKeys{
		AccessKeyMetadata:      accessKeyMetadata,
		GithubUsers:            githubUsers,
		GithubAccessToken:      accessToken,
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
		assert.True(t, len(docWrapper.GithubEventCheckpoints) == 2)
		assert.Equal(t, *docWrapper.GithubEventCheckpoints[*githubUser.Login].ID, *mockGithubEvent2.ID)

	}

	// Emit the result
	docWrapperBytes, err := json.MarshalIndent(docWrapper, "", "    ")
	assert.True(t, err == nil)
	log.Printf("docWrapperBytes: %v", string(docWrapperBytes))

}

// Regression test against mock for reprducing https://github.com/tleyden/keynuker/issues/6
func TestScanGithubLargePushEvents(t *testing.T) {

	defer gock.Off() // Flush pending mocks after test execution

	payloadStr := `
	{
      "push_id": 1932650429,
      "size": 1,
      "distinct_size": 1,
      "ref": "refs/heads/KeyNukerIntegrationTestBranch-5a2f42dd3058f53ac9c5f22153257db7b594c663",
      "head": "5ce788260531d0830c67eaa9dbad0279d85373ed",
      "before": "21f1643d6b32cf984eea4118943146a349ffa0bd",
      "commits": [
        {
          "sha": "5ce788260531d0830c67eaa9dbad0279d85373ed",
          "author": {
            "email": "traun.leyden@gmail.com",
            "name": "Traun Leyden"
          },
          "message": "KeyNukerEndToEndIntegrationTest harmless commit 19",
          "distinct": true,
          "url": "https://api.github.com/repos/tleyden/keynuker-playground/commits/5ce788260531d0830c67eaa9dbad0279d85373ed"
        }
      ]
	}
	`
	jsonRaw := json.RawMessage(payloadStr)

	events := []*github.Event{
		{
			ID: aws.String("FakeEvent1"),
			CreatedAt: aws.Time(time.Now()),
			Type: aws.String("PushEvent"),
			RawPayload: &jsonRaw,
		},
	}

	gock.New("https://api.github.com").
		Get("/users/tleyden/events").
		MatchParam("per_page", "100").
		Reply(200).
		JSON(events)

	// res, err := http.Get("https://api.github.com/bar")

	// log.Printf("res: %v, err: %v", res, err)

	var ok bool
	accessToken, ok := os.LookupEnv(keynuker_go_common.EnvVarKeyNukerTestGithubAccessToken)
	if !ok {
		t.Skip("You must define environment variable keynuker_test_gh_access_token to run this test")
	}

	// Create mock user event fetcher
	fetcher := NewGoGithubUserEventFetcher(accessToken)

	githubUser := &github.User{
		Login: aws.String("tleyden"),
	}

	// Make a fake checkpoint event that has the current timestamp
	githubCheckpointEvent := &github.Event{
		CreatedAt: aws.Time(time.Now().Add(time.Hour * -24)),
	}
	githubEventCheckpoints := GithubEventCheckpoints{}
	githubEventCheckpoints[*githubUser.Login] = githubCheckpointEvent

	leakedKey := "FakeAccessKey"

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
		GithubAccessToken:      "github_access_token",
		KeyNukerOrg:            "test",
		GithubEventCheckpoints: githubEventCheckpoints,
	}

	// Create events scanner and run
	scanner := NewGithubUserEventsScanner(fetcher)
	docWrapper, err := scanner.ScanAwsKeys(params)

	log.Printf("doc result: %v err: %v", docWrapper, err)

}
