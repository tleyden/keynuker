// Licensed to the Apache Software Foundation (ASF) under one or more contributor license agreements;
// and to You under the Apache License, Version 2.0.  See LICENSE in project root for full license + copyright.

package keynuker

import (
	"context"
	"github.com/couchbaselabs/go.assert"
	"github.com/tleyden/keynuker/keynuker-go-common"
	"log"
	"os"
	"testing"
)

// Not much of a unit test, just allows running ghUserAggregator.ListMembers() by hand in isolation
func TestGithubUserAggregator(t *testing.T) {

	accessToken, ok := os.LookupEnv(keynuker_go_common.EnvVarKeyNukerTestGithubAccessToken)
	if !ok {
		t.Skip("You must define environment variable keynuker_test_gh_access_token to run this test")
	}

	ghOrgs := []string{
		"couchbase",
		"couchbaselabs",
		"couchbasedeps",
		"couchbase-partners",
	}

	ghUserAggregator := NewGithubUserAggregator(ghOrgs, accessToken)
	users, err := ghUserAggregator.ListMembers(context.Background())

	assert.True(t, err == nil)

	log.Printf("Members of all orgs %v: %v", ghOrgs, users)
	log.Printf("# of Members: %v.", len(users))

	assert.True(t, len(users) > 250)

}
