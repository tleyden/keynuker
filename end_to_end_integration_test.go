// Licensed to the Apache Software Foundation (ASF) under one or more contributor license agreements;
// and to You under the Apache License, Version 2.0.  See LICENSE in project root for full license + copyright.

package keynuker

import (
	"fmt"
	"testing"

	"log"
	"os"

	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/iam"
	"github.com/tleyden/keynuker/keynuker-go-common"
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
	Cleanup() error
}

type LeakKeyViaCommit struct{}

func (lkvc LeakKeyViaCommit) Leak(accessKey *iam.AccessKey) error {

	// TODO: commit a change to a private github repo in the github org
	// TODO: being monitored.

	return fmt.Errorf("TODO: leak key on github repo")
}

func (lkvc LeakKeyViaCommit) Cleanup() error {
	return nil
}

func GetEndToEndKeyLeakScenarios() []KeyLeakScenario {
	return []KeyLeakScenario{
		LeakKeyViaCommit{},
	}
}

type EndToEndIntegrationTest struct {
	IamUsername string
	IamService  *iam.IAM
	AwsSession  *session.Session
}

func NewEndToEndIntegrationTest() *EndToEndIntegrationTest {
	return &EndToEndIntegrationTest{}
}

func (e *EndToEndIntegrationTest) InitAwsIamSession() error {

	// Todo: figure out a better way than hardcoding this
	e.IamUsername = "KeyNuker"

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

	keyLeakScenarios := GetEndToEndKeyLeakScenarios()
	for _, keyLeakScenario := range keyLeakScenarios {

		awsAccessKey, err := e.CreateKeyToLeak()
		if err != nil {
			return err
		}
		if err := keyLeakScenario.Leak(awsAccessKey); err != nil {
			return fmt.Errorf("Error running testScenario: %v", err)
		}

		if err := e.RunKeyNuker(awsAccessKey); err != nil {
			return fmt.Errorf("Error running keynuker: %v", err)
		}

		nuked, err := e.VerifyKeyNuked(awsAccessKey)
		if err != nil {
			return fmt.Errorf("Error verifying key was nuked: %v", err)
		}

		if !nuked {
			// TODO: ManuallyNukeKey(awsAccessKey) <-- otherwise, leaves residue
			return fmt.Errorf("Key %v should have been nuked, but it wasn't", *awsAccessKey.AccessKeyId)
		}

		if err := keyLeakScenario.Cleanup(); err != nil {
			return fmt.Errorf("Error cleaning up keyleak scenario: %v", err)
		}

	}

	return nil

}

// NOTE: the aws key will need more permissions than usual, will need to be able to create AWS keys.
// Also, the aws key must be owned by a user named "KeyNuker"
func (e EndToEndIntegrationTest) CreateKeyToLeak() (accessKey *iam.AccessKey, err error) {

	// Discover list of IAM users in account
	createAccessKeyInput := &iam.CreateAccessKeyInput{
		UserName: aws.String(e.IamUsername),
	}
	createAccessKeyOutput, err := e.IamService.CreateAccessKey(createAccessKeyInput)
	if err != nil {
		return nil, fmt.Errorf("Error creating access key: %v", err)
	}

	return createAccessKeyOutput.AccessKey, nil

}

func (e EndToEndIntegrationTest) VerifyKeyNuked(nukedAccessKey *iam.AccessKey) (nuked bool, err error) {

	listAccessKeysInput := &iam.ListAccessKeysInput{
		UserName: aws.String(e.IamUsername),
		MaxItems: aws.Int64(1000),
	}

	listAccessKeysOutput, err := e.IamService.ListAccessKeys(listAccessKeysInput)
	if err != nil {
		return false, fmt.Errorf("Error listing access keys for user: %v. Err: %v", e.IamUsername, err)
	}

	// Panic if more than 1K results, which is not handled
	if *listAccessKeysOutput.IsTruncated {
		// TODO: remove panic and put in a paginated loop.  Move to tleyden/awsutils + unit tests against mocks
		return false, fmt.Errorf("Output is truncated and this code does not handle it")
	}

	for _, accessKeyMetadata := range listAccessKeysOutput.AccessKeyMetadata {

		if *accessKeyMetadata.AccessKeyId == *nukedAccessKey.AccessKeyId {
			// Ugh, found the key that was supposed to be nuked.  Something is not working.
			return false, nil
		}
	}

	return true, nil

}

func (e EndToEndIntegrationTest) RunKeyNuker(accessKeyToNuke *iam.AccessKey) (err error) {

	keyNukerOrg := keynuker_go_common.DefaultKeyNukerOrg

	// ------------------------ Fetch Aws Keys -------------------------

	targetAwsAccounts, err := GetTargetAwsAccountsFromEnv()
	if err != nil {
		return err
	}

	paramsFetchAwsKeys := ParamsFetchAwsKeys{
		KeyNukerOrg:       keyNukerOrg,
		TargetAwsAccounts: targetAwsAccounts,
	}

	fetchedAwsKeys, err := FetchAwsKeys(paramsFetchAwsKeys)
	if err != nil {
		return err
	}
	log.Printf("fetchedAwsKeys: %v", fetchedAwsKeys)

	// ------------------------ Github User Aggregator -------------------------

	githubAccessToken, ok := os.LookupEnv(keynuker_go_common.EnvVarKeyNukerTestGithubAccessToken)
	if !ok {
		return fmt.Errorf("You must define environment variable keynuker_test_gh_access_token to run this test")
	}

	githubOrgs, err := GetGithubOrgsFromEnv()
	if err != nil {
		return err
	}

	paramsAggregateGithubUsers := ParamsGithubUserAggregator{
		KeyNukerOrg:       keyNukerOrg,
		GithubAccessToken: githubAccessToken,
		GithubOrgs:        githubOrgs,
	}

	resultAggregateGithubUsers, err := AggregateGithubUsers(paramsAggregateGithubUsers)
	if err != nil {
		return err
	}

	// ------------------------ Github User Events Scanner -------------------------

	paramsScanGithubUserEventsForAwsKeys := ParamsScanGithubUserEventsForAwsKeys{
		KeyNukerOrg:       keyNukerOrg,
		GithubAccessToken: githubAccessToken,
		GithubUsers:       resultAggregateGithubUsers.Doc.GithubUsers,
		AccessKeyMetadata: fetchedAwsKeys.Doc.AccessKeyMetadata,
	}

	paramsScanGithubUserEventsForAwsKeys = paramsScanGithubUserEventsForAwsKeys.WithDefaultCheckpoints(time.Hour * -12)

	fetcher := NewGoGithubUserEventFetcher(githubAccessToken)

	scanner := NewGithubUserEventsScanner(fetcher)

	scanAwsKeysResults, err := scanner.ScanAwsKeys(paramsScanGithubUserEventsForAwsKeys)
	if err != nil {
		return err
	}

	// ------------------------ Nuke Leaked Aws Keys -------------------------

	params := ParamsNukeLeakedAwsKeys{
		KeyNukerOrg:            keyNukerOrg,
		TargetAwsAccounts:      targetAwsAccounts,
		LeakedKeyEvents:        scanAwsKeysResults.LeakedKeyEvents,
		GithubEventCheckpoints: scanAwsKeysResults.GithubEventCheckpoints,
	}

	resultNukeLeakedAwsKeys, err := NukeLeakedAwsKeys(params)
	if err != nil {
		return err
	}

	if len(resultNukeLeakedAwsKeys.NukedKeyEvents) <= 0 {
		return fmt.Errorf("Expected a key to be nuked, but none were nuked.  result: %+v", resultNukeLeakedAwsKeys)
	}

	nukedKeyEvent := resultNukeLeakedAwsKeys.NukedKeyEvents[0]
	if *nukedKeyEvent.LeakedKeyEvent.AccessKeyMetadata.AccessKeyId != *accessKeyToNuke.AccessKeyId {
		return fmt.Errorf(
			"Expected to nuke: %v, but nuked: %v",
			*accessKeyToNuke.AccessKeyId,
			*nukedKeyEvent.LeakedKeyEvent.AccessKeyMetadata.AccessKeyId,
		)
	}

	return nil

}
