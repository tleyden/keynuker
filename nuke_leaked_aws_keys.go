// Licensed to the Apache Software Foundation (ASF) under one or more contributor license agreements;
// and to You under the Apache License, Version 2.0.  See LICENSE in project root for full license + copyright.

package keynuker

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/iam"
)

// . Connect to AWS
// . For each leaked key
// .. Delete the key
// . Return doc with nuked keys and validated github event checkpoints
func NukeLeakedAwsKeys(params ParamsNukeLeakedAwsKeys) (doc DocumentNukeLeakedAwsKeys, err error) {

	for _, leakedKeyEvent := range params.LeakedKeyEvents {

		targetAwsAccount, err := FindAwsAccount(params.TargetAwsAccounts, leakedKeyEvent.AccessKeyMetadata.MonitorAwsAccessKeyId)
		if err != nil {
			return doc, err
		}

		// Create AWS session
		sess, err := session.NewSession(&aws.Config{
			Credentials: credentials.NewCredentials(
				&credentials.StaticProvider{Value: credentials.Value{
					AccessKeyID:     targetAwsAccount.AwsAccessKeyId,
					SecretAccessKey: targetAwsAccount.AwsSecretAccessKey,
				}},
			),
		})
		if err != nil {
			return doc, err
		}

		// Create IAM client with the session.
		svc := iam.New(sess)

		deleteAccessKeyInput := &iam.DeleteAccessKeyInput{
			AccessKeyId: leakedKeyEvent.AccessKeyMetadata.AccessKeyId,
			UserName:    leakedKeyEvent.AccessKeyMetadata.UserName,
		}
		deleteAccessKeyOutput, errDelKey := svc.DeleteAccessKey(deleteAccessKeyInput)
		if errDelKey != nil {
			return doc, errDelKey
		}

		nukedKeyEvent := NukedKeyEvent{
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

	// The list of AWS accounts in scope
	TargetAwsAccounts []TargetAwsAccount

	// This is the name of the KeyNuker "org/tenant".  Defaults to "default", but allows to be extended multi-tenant.
	KeyNukerOrg string

	// The leaked keys to nuke
	LeakedKeyEvents []LeakedKeyEvent

	// The keep per-user checkpoints which will be validated/propagated after keys are successfully nuked
	GithubEventCheckpoints GithubEventCheckpoints
}

type DocumentNukeLeakedAwsKeys struct {
	Id                     string `json:"_id"`
	NukedKeyEvents         []NukedKeyEvent
	GithubEventCheckpoints GithubEventCheckpoints
}

func (p ParamsNukeLeakedAwsKeys) Validate() error {
	return nil

}

func FindAwsAccount(targetAwsAccounts []TargetAwsAccount, monitorAwsAccessKeyId string) (TargetAwsAccount, error) {
	for _, targetAwsAccount := range targetAwsAccounts {
		if targetAwsAccount.AwsSecretAccessKey == monitorAwsAccessKeyId {
			return targetAwsAccount, nil
		}
	}
	return TargetAwsAccount{}, fmt.Errorf("Could not find TargetAwsAccount with monitorAwsAccessKeyId: %v in %v", monitorAwsAccessKeyId, targetAwsAccounts)

}
