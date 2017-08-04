// Licensed to the Apache Software Foundation (ASF) under one or more contributor license agreements;
// and to You under the Apache License, Version 2.0.  See LICENSE in project root for full license + copyright.

package keynuker

import (
	"fmt"
	"strings"

	"log"

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

		// Don't nuke the key that keynuker itself is using!
		if leakedKeyEvent.LeakedKeyIsMonitorKey() {
			return doc, fmt.Errorf("Cannot nuke the key being used to monitor.  LeakedKeyEvent: %+v", leakedKeyEvent)
		}

		// Find which target AWS account this leaked key is associated with
		targetAwsAccount, err := FindAwsAccount(params.TargetAwsAccounts, leakedKeyEvent.AccessKeyMetadata.MonitorAwsAccessKeyId)
		if err != nil {
			return doc, err
		}

		// Create AWS session on target AWS account
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

		// Nuke the key from AWS
		deleteAccessKeyInput := &iam.DeleteAccessKeyInput{
			AccessKeyId: leakedKeyEvent.AccessKeyMetadata.AccessKeyId,
			UserName:    leakedKeyEvent.AccessKeyMetadata.UserName,
		}
		log.Printf("Nuking key: %v", *leakedKeyEvent.AccessKeyMetadata.AccessKeyId)
		deleteAccessKeyOutput, errDelKey := svc.DeleteAccessKey(deleteAccessKeyInput)

		log.Printf("errDelKey: %v, %+v, type: %T", errDelKey, errDelKey, errDelKey)

		// Only consider it an error if it's not a "KeyNotFound error", which means the key was already nuked
		if errDelKey != nil && !IsKeyNotFoundError(errDelKey) {
			return doc, errDelKey
		}

		//if errDelKey != nil {
		//	// TODO: if the error is "Err: NoSuchEntity: The Access Key with id ****** cannot be found.", no need to return error and cause a panic
		//	// TODO: can be fixed by first querying API and making sure Access Key actually exists.  (may have already been nuked)
		//	return doc, errDelKey
		//}

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

func IsKeyNotFoundError(err error) bool {
	// awsErr, ok := err.(awserr.BatchedErrors)
	//if awsErr, ok := err.(awserr.Error); ok {
	//
	//
	//}

	if err == nil {
		return false
	}

	// Otherwise resort to a string search
	if strings.Contains(err.Error(), "NoSuchEntity") && strings.Contains(err.Error(), "cannot be found") {
		return true
	}

	return false
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

	if ArtificialErrorInjection() {
		return FindAwsAccountArtificialError(targetAwsAccounts, monitorAwsAccessKeyId)
	}

	for _, targetAwsAccount := range targetAwsAccounts {
		if strings.TrimSpace(targetAwsAccount.AwsAccessKeyId) == strings.TrimSpace(monitorAwsAccessKeyId) {
			return targetAwsAccount, nil
		}
	}
	return TargetAwsAccount{}, fmt.Errorf("Could not find TargetAwsAccount with monitorAwsAccessKeyId: |%v| in %+v", monitorAwsAccessKeyId, targetAwsAccounts)

}
