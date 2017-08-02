// Licensed to the Apache Software Foundation (ASF) under one or more contributor license agreements;
// and to You under the Apache License, Version 2.0.  See LICENSE in project root for full license + copyright.

package keynuker

import (
	"github.com/stretchr/testify/assert"
	"log"
	"os"
	"testing"
	"time"

	"encoding/json"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/google/go-github/github"
	"github.com/tleyden/keynuker/keynuker-go-common"
)

func TestNukeLeakedAwsKeys(t *testing.T) {

	targetAwsAccountsRaw, ok := os.LookupEnv(keynuker_go_common.EnvVarKeyNukerTestTargetAwsAccounts)
	if !ok {
		t.Skip("You must define environment variable %s to run this test", keynuker_go_common.EnvVarKeyNukerTestTargetAwsAccounts)
	}

	targetAwsAccounts := []TargetAwsAccount{}

	err := json.Unmarshal([]byte(targetAwsAccountsRaw), &targetAwsAccounts)
	assert.NoError(t, err, "Unexpected Error")

	leakedKeyEvent := LeakedKeyEvent{
		AccessKeyMetadata: FetchedAwsAccessKey{
			AccessKeyId: aws.String("******"),
			UserName:    aws.String("******"),
		},
	}

	// Make a fake checkpoint event that has the current timestamp
	githubCheckpointEvent := &github.Event{
		CreatedAt: aws.Time(time.Now().Add(time.Hour * -24)),
	}
	githubEventCheckpoints := GithubEventCheckpoints{}
	githubEventCheckpoints["tleyden"] = githubCheckpointEvent

	params := ParamsNukeLeakedAwsKeys{
		KeyNukerOrg:            "default",
		TargetAwsAccounts:      targetAwsAccounts,
		LeakedKeyEvents:        []LeakedKeyEvent{leakedKeyEvent},
		GithubEventCheckpoints: githubEventCheckpoints,
	}

	doc, err := NukeLeakedAwsKeys(params)
	log.Printf("NukeLeakedAwsKeys() returned doc: %v, err: %v", doc, err)

	if err != nil {
		log.Printf("NukeLeakedAwsKeys() returned err: %v", err)
	}
	assert.True(t, err != nil)

	assert.True(t, len(doc.NukedKeyEvents) == 1)

}
