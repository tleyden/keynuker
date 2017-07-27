// Licensed to the Apache Software Foundation (ASF) under one or more contributor license agreements;
// and to You under the Apache License, Version 2.0.  See LICENSE in project root for full license + copyright.

package keynuker

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/iam"
	"github.com/tleyden/keynuker/keynuker-github"
)

// . Connect to AWS
// . For each leaked key
// .. Delete the key
// . Return doc with nuked keys and validated github event checkpoints
func NukeLeakedAwsKeys(params ParamsNukeLeakedAwsKeys) (doc DocumentNukeLeakedAwsKeys, err error) {

	// Create AWS session
	sess, err := session.NewSession(&aws.Config{
		Credentials: credentials.NewCredentials(
			&credentials.StaticProvider{Value: credentials.Value{
				AccessKeyID:     params.AwsAccessKeyId,
				SecretAccessKey: params.AwsSecretAccessKey,
			}},
		),
	})
	if err != nil {
		return doc, err
	}

	// Create IAM client with the session.
	svc := iam.New(sess)

	for _, leakedKeyEvent := range params.LeakedKeyEvents {

		deleteAccessKeyInput := &iam.DeleteAccessKeyInput{
			AccessKeyId: leakedKeyEvent.AccessKeyMetadata.AccessKeyId,
			UserName:    leakedKeyEvent.AccessKeyMetadata.UserName,
		}
		deleteAccessKeyOutput, errDelKey := svc.DeleteAccessKey(deleteAccessKeyInput)
		if errDelKey != nil {
			return doc, errDelKey
		}

		nukedKeyEvent := keynuker_github.NukedKeyEvent{
			LeakedKeyEvent:        leakedKeyEvent,
			DeleteAccessKeyOutput: deleteAccessKeyOutput,
		}

		doc.NukedKeyEvents = append(doc.NukedKeyEvents, nukedKeyEvent)

	}

	// Validate and propagate github checkpoints
	doc.GithubEventCheckpoints = params.GithubEventCheckpoints

	return doc, nil

}

type ParamsNukeLeakedAwsKeys struct {

	// The aws access key to connect as.  This only needs permissions to list IAM users and access keys,
	// and delete access keys (in the case they are nuked)
	AwsAccessKeyId string

	// The secret access key corresponding to AwsAccessKeyId
	AwsSecretAccessKey string

	// This is the name of the KeyNuker "org/tenant".  Defaults to "default", but allows to be extended multi-tenant.
	KeyNukerOrg string

	// The leaked keys to nuke
	LeakedKeyEvents []keynuker_github.LeakedKeyEvent

	// The keep per-user checkpoints which will be validated/propagated after keys are successfully nuked
	GithubEventCheckpoints keynuker_github.GithubEventCheckpoints
}

type DocumentNukeLeakedAwsKeys struct {
	Id                     string `json:"_id"`
	NukedKeyEvents         []keynuker_github.NukedKeyEvent
	GithubEventCheckpoints keynuker_github.GithubEventCheckpoints
}

func (p ParamsNukeLeakedAwsKeys) Validate() error {
	return nil

}
