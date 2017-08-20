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

// INSTRUCTIONS to run integration tests are in the Developer Guide (developers.adoc)

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
			log.Printf("Error cleaning up keyleak scenario: %v.  There may be test residue you should cleanup by hand", err)
		}

	}

	return nil

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

func (e EndToEndIntegrationTest) GetEndToEndKeyLeakScenarios() []KeyLeakScenario {
	return []KeyLeakScenario{
		// TEMP COMMENT NewLeakKeyViaNewGithubIssue(e.GithubAccessToken, e.GithubRepoLeakTargetRepo),
		NewLeakKeyViaOlderCommit(e.GithubAccessToken, e.GithubRepoLeakTargetRepo),
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

// ----------------------------------------------- Key Leak Scenarios --------------------------------------------------

type KeyLeakScenario interface {
	Leak(accessKey *iam.AccessKey) error
	Cleanup() error
}

// --------------------------------------------- LeakKeyViaOlderCommit Scenario ----------------------------------------

type LeakKeyViaOlderCommit struct {
	GithubAccessToken        string
	GithubRepoLeakTargetRepo string
	GithubClientWrapper      *GithubClientWrapper
	GithubUser               *github.User
	GitBranch                string
}

func NewLeakKeyViaOlderCommit(githubAccessToken, targetGithubRepo string) *LeakKeyViaOlderCommit {
	leakKeyViaOlderCommit := &LeakKeyViaOlderCommit{
		GithubAccessToken:        githubAccessToken,
		GithubRepoLeakTargetRepo: targetGithubRepo,
		GitBranch:                "refs/heads/master",
	}
	leakKeyViaOlderCommit.GithubClientWrapper = NewGithubClientWrapper(githubAccessToken)
	return leakKeyViaOlderCommit
}

func (lkvoc *LeakKeyViaOlderCommit) Leak(accessKey *iam.AccessKey) error {

	ctx := context.Background()

	var err error

	// Find out the github username (aka user login)
	lkvoc.GithubUser, _, err = lkvoc.GithubClientWrapper.ApiClient.Users.Get(ctx, "")
	if err != nil {
		return err
	}

	// Push harmless commits
	for i := 0; i < 25; i++ { // TODO: bump to 50
		body := fmt.Sprintf("LeakKeyViaOlderCommit commit %d", i)
		path := fmt.Sprintf("KeyNukerEndToEndIntegrationTest harmless commit %d", i)
		if _, err := lkvoc.PushCommit(path, body); err != nil {
			return err
		}
	}

	// Push a single commit with a leaked key
	body := fmt.Sprintf("LeakKeyViaOlderCommit access key id: %v", *accessKey.AccessKeyId)
	if _, err := lkvoc.PushCommit("KeyNukerEndToEndIntegrationTest leaked key commit", body); err != nil {
		return err
	}

	return nil

}

// Push a commit
// Based on:
//   https://gist.github.com/harlantwood/2935203
//   http://www.levibotelho.com/development/commit-a-file-with-the-github-api/
func (lkvoc LeakKeyViaOlderCommit) PushCommit(path, body string) (ref *github.Reference, err error) {

	ctx := context.Background()

	latestCommitSha, err := lkvoc.GetLatestCommitSha()
	if err != nil {
		return nil, err
	}

	// Create a tree
	// https://developer.github.com/v3/git/trees/#create-a-tree
	tree, _, err := lkvoc.GithubClientWrapper.ApiClient.Git.CreateTree(
		ctx,
		*lkvoc.GithubUser.Login,
		lkvoc.GithubRepoLeakTargetRepo,
		latestCommitSha,
		[]github.TreeEntry{
			{
				Path:    aws.String(path),
				Content: aws.String(body),
				Mode:    aws.String("100644"),
			},
		},
	)
	if err != nil {
		return nil, err
	}

	// Create a commit
	commitResult, _, err := lkvoc.GithubClientWrapper.ApiClient.Git.CreateCommit(
		ctx,
		*lkvoc.GithubUser.Login,
		lkvoc.GithubRepoLeakTargetRepo,
		&github.Commit{
			Tree:    tree,
			Message: aws.String(path),
			Parents: []github.Commit{
				{
					SHA: aws.String(latestCommitSha),
				},
			},
		},
	)
	if err != nil {
		return nil, err
	}

	log.Printf("Created commit: %v", *commitResult.SHA)


	// Update a reference
	// https://developer.github.com/v3/git/refs/#update-a-reference
	refResult, _, err := lkvoc.GithubClientWrapper.ApiClient.Git.UpdateRef(
		ctx,
		*lkvoc.GithubUser.Login,
		lkvoc.GithubRepoLeakTargetRepo,
		&github.Reference{
			Ref: aws.String(lkvoc.GitBranch),
			Object: &github.GitObject{
				SHA: commitResult.SHA,
			},
		},
		true,
	)
	if err != nil {
		return nil, err
	}

	log.Printf("Pushed commit: %v", *commitResult.SHA)

	return refResult, nil

}

func (lkvoc LeakKeyViaOlderCommit) GetLatestCommitSha() (commitSha string, err error) {

	ctx := context.Background()

	ref, _, err := lkvoc.GithubClientWrapper.ApiClient.Git.GetRef(
		ctx,
		*lkvoc.GithubUser.Login,
		lkvoc.GithubRepoLeakTargetRepo,
		lkvoc.GitBranch,
	)
	if err != nil {
		return "", err
	}

	return *ref.Object.SHA, nil
}

func (lkvoc LeakKeyViaOlderCommit) Cleanup() error {
	// TODO
	return nil
}

// ------------------------------------------ LeakKeyViaNewGithubIssue Scenario ----------------------------------------

type LeakKeyViaNewGithubIssue struct {
	GithubAccessToken          string
	GithubRepoLeakTargetRepo   string
	GithubClientWrapper        *GithubClientWrapper
	IssueCreatedForLeak        *github.Issue
	IssueCommentCreatedForLeak *github.IssueComment
	GithubUser                 *github.User
}

func NewLeakKeyViaNewGithubIssue(githubAccessToken, targetGithubRepo string) *LeakKeyViaNewGithubIssue {
	leakKeyViaNewGithubIssue := &LeakKeyViaNewGithubIssue{
		GithubAccessToken:        githubAccessToken,
		GithubRepoLeakTargetRepo: targetGithubRepo,
	}
	leakKeyViaNewGithubIssue.GithubClientWrapper = NewGithubClientWrapper(githubAccessToken)
	return leakKeyViaNewGithubIssue
}

func (lkvc *LeakKeyViaNewGithubIssue) Leak(accessKey *iam.AccessKey) error {

	ctx := context.Background()

	var err error

	// Find out the github username (aka user login)
	lkvc.GithubUser, _, err = lkvc.GithubClientWrapper.ApiClient.Users.Get(ctx, "")
	if err != nil {
		return err
	}

	// Make sure the target repo exists and is private, otherwise try to create it
	if err := lkvc.CreateOrVerifyTargetRepo(lkvc.GithubUser); err != nil {
		return err
	}

	// Post an issue
	issueRequest := &github.IssueRequest{
		Title: aws.String("KeyNuker Leaked Key 🔐 End-to-End Test"),
		//Body: aws.String(fmt.Sprintf(
		//	"Nukable 🔐💥 Key: %v.  Keynuker Project url: github.com/tleyden/keynuker",
		//	*accessKey.AccessKeyId,
		//)),
	}
	lkvc.IssueCreatedForLeak, _, err = lkvc.GithubClientWrapper.ApiClient.Issues.Create(
		ctx,
		*lkvc.GithubUser.Login,
		lkvc.GithubRepoLeakTargetRepo,
		issueRequest,
	)
	if err != nil {
		return err
	}

	// Create a comment that has leaked aws key
	issueComment := &github.IssueComment{
		Body: aws.String(fmt.Sprintf(
			"Nukable 🔐💥 Key: %v.  Keynuker Project url: github.com/tleyden/keynuker",
			*accessKey.AccessKeyId,
		)),
	}
	lkvc.IssueCommentCreatedForLeak, _, err = lkvc.GithubClientWrapper.ApiClient.Issues.CreateComment(
		ctx,
		*lkvc.GithubUser.Login,
		lkvc.GithubRepoLeakTargetRepo,
		*lkvc.IssueCreatedForLeak.Number,
		issueComment,
	)
	if err != nil {
		return err
	}

	return nil

}

func (lkvc LeakKeyViaNewGithubIssue) CreateOrVerifyTargetRepo(user *github.User) error {

	ctx := context.Background()

	repo, _, err := lkvc.GithubClientWrapper.ApiClient.Repositories.Get(ctx, *user.Login, lkvc.GithubRepoLeakTargetRepo)
	if err == nil {
		// the repo exists, but make sure it's private
		if !*repo.Private {
			return fmt.Errorf("Repository %v exists, but is not private, and it's not recommended to leak a live key on a public repo", lkvc.GithubRepoLeakTargetRepo)
		}
		// it exists and it's private, nothing to do
		return nil
	}

	// If we got this far, the repo doesn't exist, so create it
	repoToCreate := &github.Repository{
		Name:      aws.String(lkvc.GithubRepoLeakTargetRepo),
		Private:   aws.Bool(true),
		HasIssues: aws.Bool(true),
		AutoInit:  aws.Bool(true),
	}
	_, _, createRepoErr := lkvc.GithubClientWrapper.ApiClient.Repositories.Create(ctx, "", repoToCreate)
	return createRepoErr

}

func (lkvc LeakKeyViaNewGithubIssue) Cleanup() error {

	ctx := context.Background()

	issueComments, _, listCommentsErr := lkvc.GithubClientWrapper.ApiClient.Issues.ListComments(
		ctx,
		*lkvc.GithubUser.Login,
		lkvc.GithubRepoLeakTargetRepo,
		*lkvc.IssueCreatedForLeak.Number,
		nil,
	)
	if listCommentsErr != nil {
		return listCommentsErr
	}

	for _, issueComment := range issueComments {
		_, deleteCommentErr := lkvc.GithubClientWrapper.ApiClient.Issues.DeleteComment(
			ctx,
			*lkvc.GithubUser.Login,
			lkvc.GithubRepoLeakTargetRepo,
			*issueComment.ID,
		)
		if deleteCommentErr != nil {
			return deleteCommentErr
		}
	}

	// Close the issue
	issueRequest := &github.IssueRequest{
		State: aws.String("closed"),
	}
	_, _, closeIssueErr := lkvc.GithubClientWrapper.ApiClient.Issues.Edit(
		ctx,
		*lkvc.GithubUser.Login,
		lkvc.GithubRepoLeakTargetRepo,
		*lkvc.IssueCreatedForLeak.Number,
		issueRequest,
	)
	if closeIssueErr != nil {
		return closeIssueErr
	}

	return nil
}
