// Licensed to the Apache Software Foundation (ASF) under one or more contributor license agreements;
// and to You under the Apache License, Version 2.0.  See LICENSE in project root for full license + copyright.

package keynuker

import (
	"log"
	"os"
	"testing"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/iam"
	"github.com/couchbaselabs/go.assert"
	"github.com/google/go-github/github"
	"github.com/tleyden/keynuker/keynuker-go-common"
)

func TestNukeLeakedAwsKeys(t *testing.T) {

	awsAccessKeyId, ok := os.LookupEnv(keynuker_go_common.EnvVarKeyNukerTestAwsAccessKeyId)
	if !ok {
		t.Skip("You must define environment variable %s to run this test", keynuker_go_common.EnvVarKeyNukerTestAwsAccessKeyId)
	}

	awsSecretAccessKey, ok := os.LookupEnv(keynuker_go_common.EnvVarKeyNukerTestAwsSecretAccessKey)
	if !ok {
		t.Skip("You must define environment variable %s to run this test", keynuker_go_common.EnvVarKeyNukerTestAwsSecretAccessKey)
	}

	leakedKeyEvent := LeakedKeyEvent{
		AccessKeyMetadata: iam.AccessKeyMetadata{
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
		AwsAccessKeyId:         awsAccessKeyId,
		AwsSecretAccessKey:     awsSecretAccessKey,
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
