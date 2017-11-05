// Licensed to the Apache Software Foundation (ASF) under one or more contributor license agreements;
// and to You under the Apache License, Version 2.0.  See LICENSE in project root for full license + copyright.

package keynuker_go_common

import "time"

// Note: this package should not have any dependencies to any other keynuker-go sub-packages,
// since they will probably create circular dependencies

const (
	DefaultKeyNukerOrg = "default"

	DocIdPrefixGithubUsers = "github_users"

	DocIdPrefixAwsKeys = "aws_keys"

	DocIdPrefixGithubEventCheckpoints = "github_event_checkpoints"
)

// Environment Variable Names
const (
	EnvVarKeyNukerTestGithubAccessToken = "KEYNUKER_INTEGRATION_TEST_GITHUB_ACCESS_TOKEN"

	EnvVarKeyNukerTestDbHost = "KEYNUKER_DB_HOST"

	EnvVarKeyNukerTestDbName = "KEYNUKER_DB_NAME"

	EnvVarKeyNukerTestDbUsername = "KEYNUKER_DB_KEY"

	EnvVarKeyNukerTestDbPassword = "KEYNUKER_DB_SECRET_KEY"

	EnvVarKeyNukerTestTargetAwsAccounts = "KEYNUKER_TARGET_AWS_ACCOUNTS"

	EnvVarKeyNukerTestIntegrationTestsEnabled = "KEYNUKER_INTEGRATION_TESTS_ENABLED"

	EnvVarKeyNukerTestIntegrationAccessKey       = "KEYNUKER_INTEGRATION_TEST_ACCESS_KEY"
	EnvVarKeyNukerTestIntegrationSecretAccessKey = "KEYNUKER_INTEGRATION_TEST_SECRET_ACCESS_KEY"

	EnvVarKeyNukerInitiatingAwsAccountCreds = "KEYNUKER_INITIATING_AWS_ACCOUNT"

	EnvVarKeyNukerTestIntegrationGithubApiBaseUrl = "KEYNUKER_INTEGRATION_TEST_GITHUB_BASE_API_URL"

	EnvVarKeyNukerTestGithubOrgs = "KEYNUKER_INTEGRATION_TEST_GITHUB_ORGS"

	EnvVarKeyNukerTestGithubLeakTargetRepo = "KEYNUKER_GITHUB_LEAK_TARGET_REPO"

	EnvVarKeyNukerEmailFromAddress = "KEYNUKER_EMAIL_FROM_ADDRESS"

	EnvVarKeyNukerAdminEmailCCAddress = "KEYNUKER_ADMIN_EMAIL_CC_ADDRESS"

	EnvVarKeyNukerMailerDomain = "KEYNUKER_MAILER_DOMAIN"

	EnvVarKeyNukerMailerApiKey = "KEYNUKER_MAILER_API_KEY"

	EnvVarKeyNukerMailerPublicApiKey = "KEYNUKER_MAILER_PUBLIC_API_KEY"

)

// Misc
const (

	// The integration test creates branches on the github repo.
	// Add improbable string at end to make it as unlikely as possible to collide with a real branch name
	KeyNukerIntegrationTestBranch = "KeyNukerIntegrationTestBranch-5a2f42dd3058f53ac9c5f22153257db7b594c663"

	GithubRefsHeadsPrefix = "refs/heads"

	GithubMasterBranch = "master"

	// The max size in bytes of blob content that will be scanned.
	// This should be raised to 100 MB once the stream based scanning is implemented.
	MaxSizeBytesBlobContent = 10000000 // 10 MB

	// The max execution time for an action in seconds
	MaxActionExecutionSeconds = time.Second * 300

	// The high watermark at which point an action should be aborted since approaching max
	HighWatermarkExecutionSeconds = MaxActionExecutionSeconds - (time.Second * 30)

	// The approximate max heap size before the action will be killed by the platform.
	MaxHeapSizeBytes = 512000000

	// The heap size at which it tries to do a clean exit and abort (but push checkpoints forward)
	HighWatermarkHeapSizeBytes = MaxHeapSizeBytes / 3

)

var (

	// If there is no recorded checkpoint for a user, how far back should the scanning go in the github user event history?
	DefaultCheckpointEventTimeWindow time.Duration
)

func init() {
	DefaultCheckpointEventTimeWindow = time.Hour * -12


}
