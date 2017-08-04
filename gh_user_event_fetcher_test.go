// Licensed to the Apache Software Foundation (ASF) under one or more contributor license agreements;
// and to You under the Apache License, Version 2.0.  See LICENSE in project root for full license + copyright.

package keynuker

import (
	"context"
	"log"
	"os"
	"testing"

	"github.com/couchbaselabs/go.assert"
	"github.com/tleyden/keynuker/keynuker-go-common"
)

// Not much of a unit test, just makes it easy to run ghUserEventFetcher.FetchUserEvents() by hand in isolation
func TestGithubUserEventFetcher(t *testing.T) {

	SkipIfIntegrationsTestsNotEnabled(t)

	accessToken, ok := os.LookupEnv(keynuker_go_common.EnvVarKeyNukerTestGithubAccessToken)
	if !ok {
		t.Skip("You must define environment variable keynuker_test_gh_access_token to run this test")
	}

	ctx := context.Background()

	ghUserEventFetcher := NewGoGithubUserEventFetcher(accessToken)

	fetchUserEventsInput := FetchUserEventsInput{
		Username: "tahmmee",
	}
	userEvents, err := ghUserEventFetcher.FetchUserEvents(ctx, fetchUserEventsInput)
	assert.True(t, err == nil)

	assert.True(t, len(userEvents) > 0)
	log.Printf("User events: %v", userEvents)
	log.Printf("# events: %v", len(userEvents))

}

// Not much of aunit test, just makes it easy to run ghUserEventFetcher.FetchDownstreamContent() by hand in isolation
func TestGithubUserEventDownstreamContentFetcher(t *testing.T) {

	SkipIfIntegrationsTestsNotEnabled(t)

	accessToken, ok := os.LookupEnv(keynuker_go_common.EnvVarKeyNukerTestGithubAccessToken)
	if !ok {
		t.Skip("You must define environment variable keynuker_test_gh_access_token to run this test")
	}

	ctx := context.Background()

	ghUserEventFetcher := NewGoGithubUserEventFetcher(accessToken)

	fetchUserEventsInput := FetchUserEventsInput{
		Username: "tleyden",
	}
	userEvents, err := ghUserEventFetcher.FetchUserEvents(ctx, fetchUserEventsInput)
	assert.True(t, err == nil)
	assert.True(t, len(userEvents) > 0)

	userEvent := userEvents[0]
	log.Printf("userEvent: %+v", userEvent)

	downstreamEventContent, err := ghUserEventFetcher.FetchDownstreamContent(ctx, userEvent)
	if err != nil {
		log.Printf("error: %v", err)
	}
	assert.True(t, err == nil)

	log.Printf("downstreamEventContent: %+v", string(downstreamEventContent))

}
