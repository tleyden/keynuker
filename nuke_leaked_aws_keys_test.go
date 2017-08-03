// Licensed to the Apache Software Foundation (ASF) under one or more contributor license agreements;
// and to You under the Apache License, Version 2.0.  See LICENSE in project root for full license + copyright.

package keynuker

import (
	"log"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"encoding/json"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/google/go-github/github"
	"github.com/tleyden/keynuker/keynuker-go-common"
)

// This is an integration test that can only be run with some manual setup:
//
// 1. Choose one of your target AWS accounts, and get the AwsAccessKeyId associated with that account and
//    use it to set EnvVarKeyNukerTestLeakedAwsAccessKeyId env var
// 2. In that account, create an IAM user called 'KeyLeaker' with no permissions and create an AWS access
//    key / secret for that user, and use it to set EnvVarKeyNukerTestTargetAccountAwsAccessKeyId env var
// 3. Run the test, it should pass
// 4. Verify the key was nuked from the actual AWS account
func TestNukeLeakedAwsKeys(t *testing.T) {

	targetAwsAccountsRaw, ok := os.LookupEnv(keynuker_go_common.EnvVarKeyNukerTestTargetAwsAccounts)
	if !ok {
		t.Skip("You must define environment variable %s to run this test", keynuker_go_common.EnvVarKeyNukerTestTargetAwsAccounts)
	}

	leakedAwsAccessKeyId, ok := os.LookupEnv(keynuker_go_common.EnvVarKeyNukerTestLeakedAwsAccessKeyId)
	if !ok {
		t.Skip("You must define environment variable %s to run this test", keynuker_go_common.EnvVarKeyNukerTestLeakedAwsAccessKeyId)
	}

	targetAccountAwsAccessKeyId, ok := os.LookupEnv(keynuker_go_common.EnvVarKeyNukerTestTargetAccountAwsAccessKeyId)
	if !ok {
		t.Skip("You must define environment variable %s to run this test", keynuker_go_common.EnvVarKeyNukerTestTargetAccountAwsAccessKeyId)
	}

	targetAwsAccounts := []TargetAwsAccount{}

	err := json.Unmarshal([]byte(targetAwsAccountsRaw), &targetAwsAccounts)
	assert.NoError(t, err, "Unexpected Error")

	leakedKeyEvent := LeakedKeyEvent{
		AccessKeyMetadata: FetchedAwsAccessKey{
			AccessKeyId:           aws.String(leakedAwsAccessKeyId),
			UserName:              aws.String("KeyLeaker"),
			MonitorAwsAccessKeyId: targetAccountAwsAccessKeyId,
		},
	}

	// Make a fake checkpoint event that has the current timestamp
	githubCheckpointEvent := &github.Event{
		CreatedAt: aws.Time(time.Now().Add(time.Hour * -24)),
	}
	githubEventCheckpoints := GithubEventCheckpoints{}
	githubEventCheckpoints["githubuser"] = githubCheckpointEvent

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
	assert.True(t, err == nil)
	assert.True(t, len(doc.NukedKeyEvents) == 1)

}

func TestFindAwsAccount(t *testing.T) {

	monitorAwsAccessKeyId := "TestAwsAccessKeyId"

	targetAwsAccounts := []TargetAwsAccount{
		{
			AwsAccessKeyId:     monitorAwsAccessKeyId,
			AwsSecretAccessKey: "TestAwsSecretAccessKey",
		},
		{
			AwsAccessKeyId:     "TestAwsAccessKeyId2",
			AwsSecretAccessKey: "TestAwsSecretAccessKey2",
		},
	}

	targetAwsAccount, err := FindAwsAccount(targetAwsAccounts, monitorAwsAccessKeyId)
	assert.NoError(t, err, "Unexpected error")
	assert.EqualValues(t, targetAwsAccount.AwsAccessKeyId, monitorAwsAccessKeyId)

}
