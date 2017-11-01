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

	"strings"

	"strconv"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/iam"
	"github.com/google/go-github/github"
	"github.com/satori/go.uuid"
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

	keynuker_go_common.LogMemoryUsageLoop()

	SkipIfIntegrationsTestsNotEnabled(t)

	endToEndIntegrationTest := NewEndToEndIntegrationTest()

	// Setup
	if err := endToEndIntegrationTest.InitAwsIamSession(); err != nil {
		t.Fatalf("Error setting up test InitAwsIamSession(): %v", err)
	}
	if err := endToEndIntegrationTest.InitGithubAccess(); err != nil {
		t.Fatalf("Error setting up test InitGithubAccess(): %v", err)
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

		if err := keyLeakScenario.RecordLatestEventFeedID(); err != nil {
			return fmt.Errorf("Error recording latest event id from user event feed: %v", err)
		}

		if err := keyLeakScenario.Leak(awsAccessKey); err != nil {
			return fmt.Errorf("Error running testScenario: %v", err)
		}

		if err := keyLeakScenario.WaitForLeakOnEventFeed(); err != nil {
			return fmt.Errorf("Error waiting for leaked event to appear on user event feed: %v", err)
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

	targetAwsAccounts, err := GetIntegrationTestTargetAwsAccountsFromEnv()
	if err != nil {
		return err
	}
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
		NewLeakKeyViaNewGithubIssue(e.GithubAccessToken, e.GithubRepoLeakTargetRepo),
		NewLeakKeyViaOlderCommit(e.GithubAccessToken, e.GithubRepoLeakTargetRepo),
		NewLeakKeyViaCreateEvent(e.GithubAccessToken, e.GithubRepoLeakTargetRepo),
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

	targetAwsAccounts, err := GetIntegrationTestTargetAwsAccountsFromEnv()
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
		KeyNukerOrg: keyNukerOrg,
		GithubConnectionParams: GithubConnectionParams{
			GithubAccessToken: e.GithubAccessToken,
		},
		GithubOrgs: e.GithubOrgs,
	}

	resultAggregateGithubUsers, err := AggregateGithubUsers(paramsAggregateGithubUsers)
	if err != nil {
		return err
	}

	// ------------------------ Github User Events Scanner -------------------------

	paramsScanGithubUserEventsForAwsKeys := ParamsScanGithubUserEventsForAwsKeys{
		KeyNukerOrg: keyNukerOrg,
		GithubConnectionParams: GithubConnectionParams{
			GithubAccessToken: e.GithubAccessToken,
		}, GithubUsers: resultAggregateGithubUsers.Doc.GithubUsers,
		AccessKeyMetadata: fetchedAwsKeys.Doc.AccessKeyMetadata,
	}

	recentEventTimeWindow := time.Minute * -1 // Last 5 seconds would probably work too, but give it some margin of error

	paramsScanGithubUserEventsForAwsKeys = paramsScanGithubUserEventsForAwsKeys.SetDefaultCheckpointsForMissing(recentEventTimeWindow)

	fetcher := NewGoGithubUserEventFetcher(e.GithubAccessToken, GetIntegrationGithubApiBaseUrl())

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

	// ------------------------ Send Post-Nuke Notification -------------------------

	paramsPostNukeNotifier := ParamsPostNukeNotifier{
		NukedKeyEvents:              resultNukeLeakedAwsKeys.NukedKeyEvents,
		KeynukerAdminEmailCCAddress: os.Getenv(keynuker_go_common.EnvVarKeyNukerAdminEmailCCAddress),
		EmailFromAddress:            os.Getenv(keynuker_go_common.EnvVarKeyNukerEmailFromAddress),
	}

	mockMailgun, err := NewMailgunFromEnvironmentVariables()
	if err != nil {
		return fmt.Errorf("Error creating mailgun: %v", err)
	}

	// Call post nuke notifier
	resultPostNukeNotifier, err := SendPostNukeMockNotifications(mockMailgun, paramsPostNukeNotifier)
	if err != nil {
		return fmt.Errorf("Error sending post-nuke notifications: %v", err)
	}
	log.Printf("resultPostNukeNotifier: %+v", resultPostNukeNotifier)

	return nil

}

// ----------------------------------------------- Key Leak Scenarios --------------------------------------------------

type KeyLeakScenario interface {
	Leak(accessKey *iam.AccessKey) error
	Cleanup() error
	RecordLatestEventFeedID() error
	WaitForLeakOnEventFeed() error
}

// --------------------------------------------- LeakKeyViaOlderCommit Scenario ----------------------------------------

type LeakKeyViaOlderCommit struct {
	GithubAccessToken        string
	GithubRepoLeakTargetRepo string
	GithubClientWrapper      *GithubClientWrapper
	GithubUser               *github.User
	GitBranch                string
	PushLargeCommit          bool // Push a commit > 1 MB
}

func NewLeakKeyViaOlderCommit(githubAccessToken, targetGithubRepo string) *LeakKeyViaOlderCommit {
	leakKeyViaOlderCommit := &LeakKeyViaOlderCommit{
		GithubAccessToken:        githubAccessToken,
		GithubRepoLeakTargetRepo: targetGithubRepo,
		GitBranch:                fmt.Sprintf("%v/%v", keynuker_go_common.GithubRefsHeadsPrefix, keynuker_go_common.KeyNukerIntegrationTestBranch),
		PushLargeCommit:          false,
	}

	githubApiBaseUrl := GetIntegrationGithubApiBaseUrl()
	leakKeyViaOlderCommit.GithubClientWrapper = NewGithubClientWrapper(githubAccessToken, githubApiBaseUrl)
	return leakKeyViaOlderCommit

}

func (lkvoc *LeakKeyViaOlderCommit) Leak(accessKey *iam.AccessKey) error {

	log.Printf("LeakKeyViaOlderCommit.Leak() called")

	ctx := context.Background()

	var err error

	// Find out the github username (aka user login)
	lkvoc.GithubUser, _, err = lkvoc.GithubClientWrapper.ApiClient.Users.Get(ctx, "")
	if err != nil {
		return err
	}

	// Create the integration test branch if it doesn't already exist
	lkvoc.CreateBranchIfNotExist(lkvoc.GitBranch)

	// Push harmless commits - needs to be greater than 20 to detect issue https://github.com/tleyden/keynuker/issues/6
	for i := 0; i < 21; i++ {
		body := fmt.Sprintf("LeakKeyViaOlderCommit commit %d", i)
		path := fmt.Sprintf("KeyNukerEndToEndIntegrationHarmlessFile-%d.txt", i)

		if _, err := lkvoc.PushCommitToExistingBranch(path, body); err != nil {
			return err
		}
	}

	// Push a commit.  Either a large (> 1MB commit), or a small commit
	switch lkvoc.PushLargeCommit {
	case true:
		// Push a large commit with a leaked key
		largeFile, err := Asset("testdata/largefile.txt")
		if err != nil {
			return fmt.Errorf("Unable to load large commit asset: %v", err)
		}
		body := fmt.Sprintf("%s access key id: %v", largeFile, *accessKey.AccessKeyId)

		// Make filename unique, otherwise will end up with a small diff from previous test rather than entire file content
		path := fmt.Sprintf("KeyNukerEndToEndIntegrationTestLeakedKeyLargefile-%s.txt", uuid.NewV4())
		log.Printf("Large commit filename: %v", path)
		commit, errPushCommit := lkvoc.PushCommitToExistingBranch(path, body)
		if errPushCommit != nil {
			return errPushCommit
		}
		log.Printf("Pushed large commit with leaked key: %v", *commit.SHA)
	case false:
		// Push a single commit with a leaked key
		body := fmt.Sprintf("LeakKeyViaOlderCommit access key id: %v", *accessKey.AccessKeyId)
		path := "KeyNukerEndToEndIntegrationTestLeakedKeyFile.txt"
		commit, errPushCommit := lkvoc.PushCommitToExistingBranch(path, body)
		if errPushCommit != nil {
			return errPushCommit
		}
		log.Printf("Pushed commit with leaked key: %v", *commit.SHA)

	}

	// Github API: https://developer.github.com/v3/repos/merging/
	mergeCommit, _, err := lkvoc.GithubClientWrapper.ApiClient.Repositories.Merge(
		ctx,
		*lkvoc.GithubUser.Login,
		lkvoc.GithubRepoLeakTargetRepo,
		&github.RepositoryMergeRequest{
			Base:          aws.String("master"),
			Head:          aws.String(keynuker_go_common.KeyNukerIntegrationTestBranch),
			CommitMessage: aws.String("Merge foo -> master"),
		},
	)
	if err != nil {
		return fmt.Errorf("Error merging branch into master: %v", err)
	}
	log.Printf("Merged foo branch into master: %v", *mergeCommit.SHA)

	// Update the commit on the branch to have the merge commit
	// So that future tests based on this branch will succeed
	refResult, _, err := lkvoc.GithubClientWrapper.ApiClient.Git.UpdateRef(
		ctx,
		*lkvoc.GithubUser.Login,
		lkvoc.GithubRepoLeakTargetRepo,
		&github.Reference{
			Ref: aws.String(lkvoc.GitBranch),
			Object: &github.GitObject{
				SHA: mergeCommit.SHA,
			},
		},
		true,
	)
	if err != nil {
		return err
	}
	log.Printf("Update branch to have merge commit: %v", *refResult.Object.SHA)

	lkvoc.WaitForPushEvent(ctx, "master", *mergeCommit.SHA)

	return nil

}

func (lkvoc LeakKeyViaOlderCommit) WaitForPushEvent(ctx context.Context, branch, headSHA string) error {

	log.Printf("WaitForPushEvent called with branch: %s headSHA: %s", branch, headSHA)

	// Wait until we see a PushEvent on events feed that is on the master branch and has a head commit
	// of mergeCommit.SHA
	for {

		listOptions := &github.ListOptions{
			PerPage: MaxPerPage,
			Page:    0,
		}

		events, _, err := lkvoc.GithubClientWrapper.ApiClient.Activity.ListEventsPerformedByUser(
			ctx,
			*lkvoc.GithubUser.Login,
			false,
			listOptions,
		)
		if err != nil {
			return err
		}

		for _, event := range events {

			eventPayload, err := event.ParsePayload()
			if err != nil {
				return err
			}

			pushEvent, ok := eventPayload.(*github.PushEvent)
			if !ok {
				continue
			}

			if strings.Contains(*pushEvent.Ref, branch) && *pushEvent.Head == headSHA {
				return nil
			}

		}

		log.Printf("WaitForPushEvent didn't see branch: %s headSHA: %s, sleeping and retrying", branch, headSHA)

		time.Sleep(time.Second * 5)
	}

	return nil

}

func (lkvoc LeakKeyViaOlderCommit) RecordLatestEventFeedID() error {
	// No-op due to being implemented via WaitForPushEvent()
	return nil
}

func (lkvoc LeakKeyViaOlderCommit) WaitForLeakOnEventFeed() error {
	// No-op due to being implemented via WaitForPushEvent()\
	return nil
}

func (lkvoc LeakKeyViaOlderCommit) CreateBranchIfNotExist(branch string) error {

	ctx := context.Background()

	_, err := lkvoc.GetLatestCommitSha(lkvoc.GitBranch)
	if err == nil {
		// Looks like the branch already exists, nothing to do
		return nil
	}

	// There was an error getting latest commit sha, assume it's because the branch doesn't exist

	log.Printf("Branch %v does not exist.  Creating.", lkvoc.GitBranch)

	// Get latest commit on master
	masterBranchName := fmt.Sprintf("%v/%v", keynuker_go_common.GithubRefsHeadsPrefix, keynuker_go_common.GithubMasterBranch)
	latestMasterCommitSha, err := lkvoc.GetLatestCommitSha(masterBranchName)
	if err != nil {
		return err
	}

	log.Printf("Creating branch %v from master commit: %v", lkvoc.GitBranch, latestMasterCommitSha)

	// Update a reference
	// https://developer.github.com/v3/git/refs/#update-a-reference
	refResult, _, err := lkvoc.GithubClientWrapper.ApiClient.Git.CreateRef(
		ctx,
		*lkvoc.GithubUser.Login,
		lkvoc.GithubRepoLeakTargetRepo,
		&github.Reference{
			Ref: aws.String(lkvoc.GitBranch),
			Object: &github.GitObject{
				SHA: &latestMasterCommitSha,
			},
		},
	)
	if err != nil {
		return err
	}

	log.Printf("Created branch %v, result.SHA: %v", lkvoc.GitBranch, *refResult.Object.SHA)

	return nil

}

func (lkvoc LeakKeyViaOlderCommit) PushCommitToNewBranch(path, body string) (commit *github.Commit, err error) {

	ctx := context.Background()

	latestCommitSha, err := lkvoc.GetLatestMasterCommitSha()
	if err != nil {
		return nil, fmt.Errorf("Error GetLatestCommitSha: %v", err)
	}

	commitResult, err := lkvoc.CreateTreeAndCommit(latestCommitSha, path, body)
	if err != nil {
		return nil, fmt.Errorf("Error calling CreateCommit: %v", err)
	}

	// Create a unique branch name (or this could be changed to first delete the branch)
	branchName := fmt.Sprintf("%v-newbranch", lkvoc.GitBranch)  // TODO: change "4" to uuid

	// Try to delete the branch.  Ignore errors -- if branch already exists it will give an error
	_, _ = lkvoc.GithubClientWrapper.ApiClient.Git.DeleteRef(
		ctx,
		*lkvoc.GithubUser.Login,
		lkvoc.GithubRepoLeakTargetRepo,
		branchName,
	)

	// Create a reference
	// https://developer.github.com/v3/git/refs/#create-a-reference

	refResult, _, err := lkvoc.GithubClientWrapper.ApiClient.Git.CreateRef(
		ctx,
		*lkvoc.GithubUser.Login,
		lkvoc.GithubRepoLeakTargetRepo,
		&github.Reference{
			Ref: aws.String(branchName),
			Object: &github.GitObject{
				SHA: commitResult.SHA,
			},
		},
	)
	if err != nil {
		return commitResult, fmt.Errorf("Error doing CreateRef: %v", err)
	}
	log.Printf("Created branch to point to commit: %v", *refResult.Object.SHA)

	return commitResult, nil

}

func (lkvoc LeakKeyViaOlderCommit) CreateTreeAndCommit(previousCommitSha, path, body string) (commit *github.Commit, err error) {

	ctx := context.Background()

	// Create a tree
	// Github API docs: https://developer.github.com/v3/git/trees/#create-a-tree
	tree, _, err := lkvoc.GithubClientWrapper.ApiClient.Git.CreateTree(
		ctx,
		*lkvoc.GithubUser.Login,
		lkvoc.GithubRepoLeakTargetRepo,
		previousCommitSha,
		[]github.TreeEntry{
			{
				Path:    aws.String(path),
				Content: aws.String(body),
				Mode:    aws.String("100644"),
				// TODO: should this be Type: blob?
			},
		},
	)
	if err != nil {
		return nil, fmt.Errorf("Error doing CreateTree: %v", err)
	}

	// Create a commit
	// Github API docs: https://developer.github.com/v3/git/commits/#create-a-commit
	commitResult, _, err := lkvoc.GithubClientWrapper.ApiClient.Git.CreateCommit(
		ctx,
		*lkvoc.GithubUser.Login,
		lkvoc.GithubRepoLeakTargetRepo,
		&github.Commit{
			Tree:    tree,
			Message: aws.String(path),
			Parents: []github.Commit{
				{
					SHA: aws.String(previousCommitSha),
				},
			},
		},
	)
	if err != nil {
		return nil, fmt.Errorf("Error doing CreateCommit: %v", err)
	}

	return commitResult, nil


}


// Push a commit
// Based on:
//   https://gist.github.com/harlantwood/2935203
//   http://www.levibotelho.com/development/commit-a-file-with-the-github-api/
func (lkvoc LeakKeyViaOlderCommit) PushCommitToExistingBranch(path, body string) (commit *github.Commit, err error) {

	ctx := context.Background()

	latestCommitSha, err := lkvoc.GetLatestCommitSha(lkvoc.GitBranch)
	if err != nil {
		return nil, fmt.Errorf("Error GetLatestCommitSha: %v", err)
	}

	commitResult, err := lkvoc.CreateTreeAndCommit(latestCommitSha, path, body)
	if err != nil {
		return nil, fmt.Errorf("Error calling CreateCommit: %v", err)
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
		return commitResult, fmt.Errorf("Error doing UpdateRef: %v", err)
	}
	log.Printf("Updated branch to point to commit: %v", *refResult.Object.SHA)

	return commitResult, nil

}

func (lkvoc LeakKeyViaOlderCommit) GetLatestCommitSha(branch string) (commitSha string, err error) {

	ctx := context.Background()

	ref, _, err := lkvoc.GithubClientWrapper.ApiClient.Git.GetRef(
		ctx,
		*lkvoc.GithubUser.Login,
		lkvoc.GithubRepoLeakTargetRepo,
		branch,
	)
	if err != nil {
		return "", err
	}

	return *ref.Object.SHA, nil
}

func (lkvoc LeakKeyViaOlderCommit) GetLatestMasterCommitSha() (commitSha string, err error) {

	ctx := context.Background()

	masterBranch := fmt.Sprintf("%v/master", keynuker_go_common.GithubRefsHeadsPrefix)

	ref, _, err := lkvoc.GithubClientWrapper.ApiClient.Git.GetRef(
		ctx,
		*lkvoc.GithubUser.Login,
		lkvoc.GithubRepoLeakTargetRepo,
		masterBranch,
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
	LatestEventPreLeak         *github.Event
}

func NewLeakKeyViaNewGithubIssue(githubAccessToken, targetGithubRepo string) *LeakKeyViaNewGithubIssue {
	leakKeyViaNewGithubIssue := &LeakKeyViaNewGithubIssue{
		GithubAccessToken:        githubAccessToken,
		GithubRepoLeakTargetRepo: targetGithubRepo,
	}
	leakKeyViaNewGithubIssue.GithubClientWrapper = NewGithubClientWrapper(githubAccessToken, GetIntegrationGithubApiBaseUrl())
	return leakKeyViaNewGithubIssue
}

func (lkvc *LeakKeyViaNewGithubIssue) Leak(accessKey *iam.AccessKey) error {

	log.Printf("LeakKeyViaNewGithubIssue.Leak() called")

	ctx := context.Background()

	var err error

	// Find out the github username (aka user login)
	if err := lkvc.discoverGithubUser(); err != nil {
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

func (lkvc *LeakKeyViaNewGithubIssue) discoverGithubUser() (err error) {

	ctx := context.Background()

	// Find out the github username (aka user login)
	lkvc.GithubUser, _, err = lkvc.GithubClientWrapper.ApiClient.Users.Get(ctx, "")
	return err

}

func (lkvc *LeakKeyViaNewGithubIssue) RecordLatestEventFeedID() error {

	ctx := context.Background()

	// Find out the github username (aka user login)
	if err := lkvc.discoverGithubUser(); err != nil {
		return err
	}

	// Get latest event from feed and store it
	listOptions := &github.ListOptions{
		PerPage: MaxPerPage,
		Page:    0,
	}

	events, _, err := lkvc.GithubClientWrapper.ApiClient.Activity.ListEventsPerformedByUser(
		ctx,
		*lkvc.GithubUser.Login,
		false,
		listOptions,
	)

	if err != nil {
		return err
	}

	for _, event := range events {

		// If we don't yet have an event stored for the LatestEventPreLeak, then use this this one no
		// matter what it's value
		if lkvc.LatestEventPreLeak == nil {
			lkvc.LatestEventPreLeak = event
			continue
		}

		existingEventId, err := lkvc.LatestEventPreLeakID()
		if err != nil {
			return err
		}

		eventId, err := strconv.Atoi(*event.ID)
		if err != nil {
			return err
		}

		if eventId > existingEventId {
			lkvc.LatestEventPreLeak = event
		}

	}

	return nil
}

func (lkvc LeakKeyViaNewGithubIssue) LatestEventPreLeakID() (int, error) {
	if lkvc.LatestEventPreLeak == nil {
		return 0, fmt.Errorf("No latest event stored")
	}
	rawEventIdString := *lkvc.LatestEventPreLeak.ID
	return strconv.Atoi(rawEventIdString)
}

func (lkvc LeakKeyViaNewGithubIssue) WaitForLeakOnEventFeed() error {

	// Polling loop to wait for an event that is _after_ the recorded latest event id, and
	// has an event type of PushEvent

	ctx := context.Background()

	for {

		// Get latest event from feed and store it

		listOptions := &github.ListOptions{
			PerPage: MaxPerPage,
			Page:    0,
		}

		events, _, err := lkvc.GithubClientWrapper.ApiClient.Activity.ListEventsPerformedByUser(
			ctx,
			*lkvc.GithubUser.Login,
			false,
			listOptions,
		)

		if err != nil {
			return err
		}

		if lkvc.LatestEventPreLeak == nil {
			return fmt.Errorf("No LatestEventPreLeak")
		}

		for _, event := range events {

			latestPreLeakEventId, err := lkvc.LatestEventPreLeakID()
			if err != nil {
				return err
			}

			eventId, err := strconv.Atoi(*event.ID)
			if err != nil {
				return err
			}

			// Ignore any events that are earlier or equal to the stored latest preleak event id
			if eventId <= latestPreLeakEventId {
				continue
			}

			// is this an IssueEvent?  Then it's the one we were waiting for
			payload, err := event.ParsePayload()
			if err != nil {
				return err
			}

			switch payload.(type) {
			case *github.IssueCommentEvent:
				// I've been WAITING for an event like you ...
				return nil
			}

		}

		log.Printf("Waiting for IssueCommentEvent to show up on event feed.  Sleep/retry")
		time.Sleep(time.Second)

	}

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


// --------------------------------------------- LeakKeyViaCreateEvent Scenario ----------------------------------------

type LeakKeyViaCreateEvent struct {
	LeakKeyViaOlderCommit
}

func NewLeakKeyViaCreateEvent(githubAccessToken, targetGithubRepo string) *LeakKeyViaCreateEvent {

	leakKeyViaOlderCommit := NewLeakKeyViaOlderCommit(githubAccessToken, targetGithubRepo)
	leakKeyViaOlderCommit.GitBranch = fmt.Sprintf("%v/%v", keynuker_go_common.GithubRefsHeadsPrefix, keynuker_go_common.KeyNukerIntegrationTestBranch)
	return &LeakKeyViaCreateEvent{
		LeakKeyViaOlderCommit: *leakKeyViaOlderCommit,
	}

}

func (lkvce *LeakKeyViaCreateEvent) Leak(accessKey *iam.AccessKey) error {

	log.Printf("LeakKeyViaCreateEvent.Leak() called")

	ctx := context.Background()

	var err error

	// Find out the github username (aka user login)
	lkvce.GithubUser, _, err = lkvce.GithubClientWrapper.ApiClient.Users.Get(ctx, "")
	if err != nil {
		return err
	}

	// Create the integration test branch if it doesn't already exist
	// lkvce.CreateBranchIfNotExist(lkvce.GitBranch)

	body := fmt.Sprintf("LeakKeyViaCreateEvent with access key: %v", *accessKey.AccessKeyId)
	path := fmt.Sprintf("KeyNukerEndToEndIntegrationHarmlessFile.txt")
	if _, err := lkvce.PushCommitToNewBranch(path, body); err != nil {
		return err
	}

	return nil

}
