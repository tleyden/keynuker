// Licensed to the Apache Software Foundation (ASF) under one or more contributor license agreements;
// and to You under the Apache License, Version 2.0.  See LICENSE in project root for full license + copyright.

package keynuker

import (
	"testing"
	"github.com/couchbaselabs/go.assert"
	"fmt"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/service/iam"
	"os/user"
	"go/doc"
)

// - Create a test AWS user w/ minimal permissions
// - Loop over leaked key scenarios
//    - Create AWS key #1
//    - Scenario 1: Commit and push text file to github private repo w/ leaked key
//    - Invoke keynuker
//    - Verify that the AWS key #1 was nuked
//    - Create AWS key #2
//    - Scenario 1: Create a secret gist w/ leaked key
//    - Verify that the AWS key #2 was nuked
// - Cleanup test user
// - Cleanup other residue (gists, etc)
func TestEndToEndIntegration(t *testing.T) {

	SkipIfIntegrationsTestsNotEnabled(t)

	endToEndIntegrationTest := NewEndToEndIntegrationTest()

	if err := endToEndIntegrationTest.InitAwsIamSession(); err != nil {
		t.Fatalf("Error setting up test: %v", err)
	}

	if err := endToEndIntegrationTest.Run(); err != nil {
		t.Fatalf("Error running test: %v", err)
	}

}

type KeyLeakScenario interface {
	Leak(accessKey *iam.AccessKey) error
}

type LeakKeyViaCommit struct {}

func (lkvc LeakKeyViaCommit) Leak(accessKey *iam.AccessKey) error {

	// TODO: commit a change to a private github repo in the github org
	// TODO: being monitored.

	return nil
}

func GetEndToEndKeyLeakScenarios() []KeyLeakScenario {
	return []KeyLeakScenario{
		LeakKeyViaCommit{},
	}
}

type EndToEndIntegrationTest struct {
	IamService *iam.IAM
	AwsSession *session.Session
}

func NewEndToEndIntegrationTest() *EndToEndIntegrationTest {
	return &EndToEndIntegrationTest{}
}

func (e *EndToEndIntegrationTest) InitAwsIamSession() error {

	targetAwsAccounts, err := GetTargetAwsAccountsFromEnv()
	if err != nil {
		return err
	}

	// Just use the first aws account
	targetAwsAccount := targetAwsAccounts[0]

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
		return fmt.Errorf("Error creating aws session: %v", err)
	}

	// Create IAM client with the session.
	svc := iam.New(sess)

	e.AwsSession = sess
	e.IamService = svc

	return nil

}

func (e EndToEndIntegrationTest) Run() error {

	SetArtificialErrorInjection(true)

	// Todo: figure out a better way than hardcoding this
	testIAMUsername := "KeyNuker"

	keyLeakScenarios := GetEndToEndKeyLeakScenarios()
	for _, keyLeakScenario := range keyLeakScenarios {

		awsAccessKey, err := CreateKeyToLeak(testIAMUsername)
		if err != nil {
			return err
		}
		if err := keyLeakScenario.Leak(awsAccessKey); err != nil {
			return fmt.Errorf("Error running testScenario: %v", err)
		}

		RunKeyNuker()

		nuked, err := VerifyKeyNuked(awsAccessKey)
		if err != nil {
			return fmt.Errorf("Error verifying key was nuked: %v", err)
		}

		if !nuked {
			// TODO: ManuallyNukeKey(awsAccessKey) <-- otherwise, leaves residue
			return fmt.Errorf("Key %v should have been nuked, but it wasn't", *awsAccessKey.AccessKeyId)
		}

		keyLeakScenario.Cleanup()


	}


}


func VerifyKeyNuked() {

	listAccessKeysInput := &iam.ListAccessKeysInput{
		UserName: user.UserName,
		MaxItems: aws.Int64(1000),
	}

	listAccessKeysOutput, err := svc.ListAccessKeys(listAccessKeysInput)
	if err != nil {
		return DocumentWrapperFetchAwsKeys{}, fmt.Errorf("Error listing access keys for user: %v.  Err: %v", user, err)
	}

	// Panic if more than 1K results, which is not handled
	if *listAccessKeysOutput.IsTruncated {
		// TODO: remove panic and put in a paginated loop.  Move to tleyden/awsutils + unit tests against mocks
		return DocumentWrapperFetchAwsKeys{}, fmt.Errorf("Output is truncated and this code does not handle it")
	}

	for _, accessKeyMetadata := range listAccessKeysOutput.AccessKeyMetadata {

		fetchedAwsAccessKey := NewFetchedAwsAccessKey(
			accessKeyMetadata,
			targetAwsAccount.AwsAccessKeyId,
		)

		doc.AccessKeyMetadata = append(doc.AccessKeyMetadata, *fetchedAwsAccessKey)
	}

}

// NOTE: the aws key will need more permissions than usual, will need to be able to create AWS keys.
// Also, the aws key must be owned by a user named "KeyNuker"
func CreateKeyToLeak(awsUserName string) (accessKey *iam.AccessKey, err error) {

	targetAwsAccounts, err := GetTargetAwsAccountsFromEnv()
	if err != nil {
		return nil, err
	}

	// Just use the first aws account
	targetAwsAccount := targetAwsAccounts[0]

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
		return nil, fmt.Errorf("Error creating aws session: %v", err)
	}

	// Create IAM client with the session.
	svc := iam.New(sess)

	// Discover list of IAM users in account
	createAccessKeyInput := &iam.CreateAccessKeyInput{
		UserName: aws.String(awsUserName),
	}
	createAccessKeyOutput, err := svc.CreateAccessKey(createAccessKeyInput)
	if err != nil {
		return nil, fmt.Errorf("Error creating access key: %v", err)
	}

	return createAccessKeyOutput.AccessKey, nil


}
