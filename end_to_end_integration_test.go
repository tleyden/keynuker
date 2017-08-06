// Licensed to the Apache Software Foundation (ASF) under one or more contributor license agreements;
// and to You under the Apache License, Version 2.0.  See LICENSE in project root for full license + copyright.

package keynuker

import (
	"fmt"
	"testing"

	"log"
	"os"

	"time"

	"context"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/iam"
	"github.com/google/go-github/github"
	"github.com/tleyden/keynuker/keynuker-go-common"
)

// - Create a test AWS user w/ minimal permissions
//     - Note: this was simplified and so it just re-uses the KeyNuker user and leaks a key under that account
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

	// Setup
	if err := endToEndIntegrationTest.InitAwsIamSession(); err != nil {
		t.Fatalf("Error setting up test: %v", err)
	}
	if err := endToEndIntegrationTest.InitGithubAccess(); err != nil {
		t.Fatalf("Error setting up test: %v", err)
	}

	// Run the full end-to-end integration test
	if err := endToEndIntegrationTest.Run(); err != nil {
		t.Fatalf("Error running test: %v", err)
	}

}

type EndToEndIntegrationTest struct {
	IamUsername              string
	IamService               *iam.IAM
	AwsSession               *session.Session
	TargetAwsAccount         TargetAwsAccount
	GithubAccessToken        string
	GithubOrgs               []string
	GithubRepoLeakTargetRepo string
}

func NewEndToEndIntegrationTest() *EndToEndIntegrationTest {
	return &EndToEndIntegrationTest{}
}

func (e *EndToEndIntegrationTest) InitGithubAccess() error {

	githubRepoLeakTargetRepo, ok := os.LookupEnv(keynuker_go_common.EnvVarKeyNukerTestGithubLeakTargetRepo)
	if !ok {
		return fmt.Errorf("You must define environment variable %v to run this test", keynuker_go_common.EnvVarKeyNukerTestGithubLeakTargetRepo)
	}
	e.GithubRepoLeakTargetRepo = githubRepoLeakTargetRepo

	githubAccessToken, ok := os.LookupEnv(keynuker_go_common.EnvVarKeyNukerTestGithubAccessToken)
	if !ok {
		return fmt.Errorf("You must define environment variable %v to run this test", keynuker_go_common.EnvVarKeyNukerTestGithubAccessToken)
	}
	e.GithubAccessToken = githubAccessToken

	githubOrgs, err := GetGithubOrgsFromEnv()
	if err != nil {
		return err
	}
	e.GithubOrgs = githubOrgs

	return nil

}

func (e *EndToEndIntegrationTest) InitAwsIamSession() error {


	targetAwsAccounts, err := GetTargetAwsAccountsFromEnv()
	if err != nil {
		return err
	}

	// Just use the first aws account
	e.TargetAwsAccount = targetAwsAccounts[0]

	// Create AWS session
	sess, err := session.NewSession(&aws.Config{
		Credentials: credentials.NewCredentials(
			&credentials.StaticProvider{Value: credentials.Value{
				AccessKeyID:     e.TargetAwsAccount.AwsAccessKeyId,
				SecretAccessKey: e.TargetAwsAccount.AwsSecretAccessKey,
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

	// Discover IAM username based on aws key
	username, err := e.DiscoverIAMUsernameForKey(e.TargetAwsAccount.AwsAccessKeyId)
	if err != nil {
		return err
	}

	e.IamUsername = username

	return nil

}

// List all users
// For each user
// List all keys
// If you find the AwsAccessKeyId param, return the current user
func (e EndToEndIntegrationTest) DiscoverIAMUsernameForKey(AwsAccessKeyId string) (username string, err error) {

	// Fetch list of IAM users
	iamUsers, err := FetchIAMUsers(e.IamService)
	if err != nil {
		return "", err
	}

	for _, user := range iamUsers {

		listAccessKeysInput := &iam.ListAccessKeysInput{
			UserName: user.UserName,
			MaxItems: aws.Int64(1000),
		}

		listAccessKeysOutput, err := e.IamService.ListAccessKeys(listAccessKeysInput)
		if err != nil {
			return "", fmt.Errorf("Error listing access keys for user: %v.  Err: %v", user, err)
		}

		// Panic if more than 1K results, which is not handled
		if *listAccessKeysOutput.IsTruncated {
			// TODO: remove panic and put in a paginated loop.  Move to tleyden/awsutils + unit tests against mocks
			return "", fmt.Errorf("Output is truncated and this code does not handle it")
		}

		for _, accessKeyMetadata := range listAccessKeysOutput.AccessKeyMetadata {

			if *accessKeyMetadata.AccessKeyId == AwsAccessKeyId {
				return *user.UserName, nil
			}

		}

	}

	return "", fmt.Errorf("Unable to lookup username for key")

}

func (e EndToEndIntegrationTest) Run() error {

	// Set this to true to verify that the end-to-end integration test catches a real bug
	SetArtificialErrorInjection(false)

	keyLeakScenarios := e.GetEndToEndKeyLeakScenarios()
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
			e.CleanupUnNuked(awsAccessKey)
			return fmt.Errorf("Key %v should have been nuked, but it wasn't", *awsAccessKey.AccessKeyId)
		}

		if err := keyLeakScenario.Cleanup(); err != nil {
			return fmt.Errorf("Error cleaning up keyleak scenario: %v", err)
		}

	}

	return nil

}

func (e EndToEndIntegrationTest) GetEndToEndKeyLeakScenarios() []KeyLeakScenario {
	return []KeyLeakScenario{
		NewLeakKeyViaCommit(e.GithubAccessToken, e.GithubRepoLeakTargetRepo),
	}
}

// NOTE: the aws key will need more permissions than usual, will need to be able to create AWS keys.
// Also, the aws key must be owned by a user named "KeyNuker"
func (e EndToEndIntegrationTest) CreateKeyToLeak() (accessKey *iam.AccessKey, err error) {

	createAccessKeyInput := &iam.CreateAccessKeyInput{
		UserName: aws.String(e.IamUsername),
	}
	createAccessKeyOutput, err := e.IamService.CreateAccessKey(createAccessKeyInput)
	if err != nil {
		return nil, fmt.Errorf("Error creating access key: %v", err)
	}

	return createAccessKeyOutput.AccessKey, nil

}

// If the initial attempt to nuke the key failed for some reason, maybe a bug, nuke it here
func (e EndToEndIntegrationTest) CleanupUnNuked(accessKeyNukeFailed *iam.AccessKey) (err error) {

	// Nuke the key from AWS
	deleteAccessKeyInput := &iam.DeleteAccessKeyInput{
		AccessKeyId: accessKeyNukeFailed.AccessKeyId,
		UserName:    accessKeyNukeFailed.UserName,
	}
	_, errDelKey := e.IamService.DeleteAccessKey(deleteAccessKeyInput)

	// Only consider it an error if it's not a "KeyNotFound error", which means the key was already nuked
	if errDelKey != nil && !IsKeyNotFoundError(errDelKey) {
		return nil
	}

	return err

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

	paramsAggregateGithubUsers := ParamsGithubUserAggregator{
		KeyNukerOrg:       keyNukerOrg,
		GithubAccessToken: e.GithubAccessToken,
		GithubOrgs:        e.GithubOrgs,
	}

	resultAggregateGithubUsers, err := AggregateGithubUsers(paramsAggregateGithubUsers)
	if err != nil {
		return err
	}

	// ------------------------ Github User Events Scanner -------------------------

	paramsScanGithubUserEventsForAwsKeys := ParamsScanGithubUserEventsForAwsKeys{
		KeyNukerOrg:       keyNukerOrg,
		GithubAccessToken: e.GithubAccessToken,
		GithubUsers:       resultAggregateGithubUsers.Doc.GithubUsers,
		AccessKeyMetadata: fetchedAwsKeys.Doc.AccessKeyMetadata,
	}

	recentEventTimeWindow := time.Minute * -10 // Last 5 seconds would probably work too, but give it some margin of error

	paramsScanGithubUserEventsForAwsKeys = paramsScanGithubUserEventsForAwsKeys.WithDefaultCheckpoints(recentEventTimeWindow)

	fetcher := NewGoGithubUserEventFetcher(e.GithubAccessToken)

	scanner := NewGithubUserEventsScanner(fetcher)

	scanAwsKeysResults, err := scanner.ScanAwsKeys(paramsScanGithubUserEventsForAwsKeys)
	if err != nil {
		return err
	}

	// ------------------------ Nuke Leaked Aws Keys -------------------------

	log.Printf("LeakedKeyEvents: %+v", scanAwsKeysResults.LeakedKeyEvents)

	params := ParamsNukeLeakedAwsKeys{
		KeyNukerOrg:            keyNukerOrg,
		TargetAwsAccounts:      targetAwsAccounts,
		LeakedKeyEvents:        scanAwsKeysResults.LeakedKeyEvents,
		GithubEventCheckpoints: scanAwsKeysResults.GithubEventCheckpoints,
	}

	resultNukeLeakedAwsKeys, err := NukeLeakedAwsKeys(params)
	if err != nil {
		return fmt.Errorf("Error nuking leaked aws keys: %v", err)
	}

	if len(resultNukeLeakedAwsKeys.NukedKeyEvents) <= 0 {
		return fmt.Errorf("Expected a key to be nuked, but none were nuked.  result: %+v", resultNukeLeakedAwsKeys)
	}

	for _, nukedKeyEvent := range resultNukeLeakedAwsKeys.NukedKeyEvents {
		log.Printf("NukedKeyEvent: %+v", nukedKeyEvent)
		if *nukedKeyEvent.LeakedKeyEvent.AccessKeyMetadata.AccessKeyId != *accessKeyToNuke.AccessKeyId {
			return fmt.Errorf(
				"Expected to nuke: %v, but nuked: %v",
				*accessKeyToNuke.AccessKeyId,
				*nukedKeyEvent.LeakedKeyEvent.AccessKeyMetadata.AccessKeyId,
			)
		}
	}

	return nil

}

type KeyLeakScenario interface {
	Leak(accessKey *iam.AccessKey) error
	Cleanup() error
}

type LeakKeyViaNewGithubIssue struct {
	GithubAccessToken        string
	GithubRepoLeakTargetRepo string
}

func NewLeakKeyViaCommit(githubAccessToken, targetGithubRepo string) *LeakKeyViaNewGithubIssue {
	return &LeakKeyViaNewGithubIssue{
		GithubAccessToken:        githubAccessToken,
		GithubRepoLeakTargetRepo: targetGithubRepo,
	}
}

func (lkvc LeakKeyViaNewGithubIssue) Leak(accessKey *iam.AccessKey) error {

	githubApiClient := NewGithubClientWrapper(lkvc.GithubAccessToken)

	ctx := context.Background()

	user, _, err := githubApiClient.ApiClient.Users.Get(ctx, "")
	if err != nil {
		return err
	}
	username := *user.Name

	log.Printf("github login: %v, name: %v", *user.Login, username)

	issueRequest := &github.IssueRequest{
		Title: aws.String("KeyNuker Leaked Key 🔐 End-to-End Test"),
		Body:  aws.String(fmt.Sprintf("Nukable 🔐💥 Key: %v.  Keynuker Project url: github.com/tleyden/keynuker", *accessKey.AccessKeyId)),
	}
	_, _, err = githubApiClient.ApiClient.Issues.Create(ctx, *user.Login, lkvc.GithubRepoLeakTargetRepo, issueRequest)
	if err != nil {
		return err
	}

	return nil

}

func (lkvc LeakKeyViaNewGithubIssue) Cleanup() error {
	return nil
}

