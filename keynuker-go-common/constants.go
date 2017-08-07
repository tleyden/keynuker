// Licensed to the Apache Software Foundation (ASF) under one or more contributor license agreements;
// and to You under the Apache License, Version 2.0.  See LICENSE in project root for full license + copyright.

package keynuker_go_common

// Note: this package should not have any dependencies to any other keynuker-go sub-packages,
// since they will probably create circular dependencies

const (
	DefaultKeyNukerOrg = "default"

	DocIdPrefixGithubUsers = "github_users"

	DocIdPrefixAwsKeys = "aws_keys"
)

// Environment Variable Names
const (
	EnvVarKeyNukerTestGithubAccessToken = "GITHUB_ACCESS_TOKEN"

	EnvVarKeyNukerTestDbHost = "KEYNUKER_DB_HOST"

	EnvVarKeyNukerTestDbName = "KEYNUKER_DB_NAME"

	EnvVarKeyNukerTestDbUsername = "KEYNUKER_DB_KEY"

	EnvVarKeyNukerTestDbPassword = "KEYNUKER_DB_SECRET_KEY"

	EnvVarKeyNukerTestTargetAwsAccounts = "TARGET_AWS_ACCOUNTS"

	EnvVarKeyNukerTestTargetAccountAwsAccessKeyId = "TARGET_ACCOUNT_AWS_ACCESS_KEY_ID"

	EnvVarKeyNukerTestIntegrationTestsEnabled = "KEYNUKER_INTEGRATION_TESTS_ENABLED"

	EnvVarKeyNukerTestGithubOrgs = "GITHUB_ORGS"

	EnvVarKeyNukerTestGithubLeakTargetRepo = "KEYNUKER_GITHUB_LEAK_TARGET_REPO"


)
