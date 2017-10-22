// Licensed to the Apache Software Foundation (ASF) under one or more contributor license agreements;
// and to You under the Apache License, Version 2.0.  See LICENSE in project root for full license + copyright.

package keynuker

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"testing"

	"github.com/tleyden/keynuker/keynuker-go-common"
)

var (
	ArtificialErrorInjectionEnabled = false
)

func SkipIfIntegrationsTestsNotEnabled(t *testing.T) {

	errMsg := fmt.Sprintf("You must define environment variable %s and set to true to enable integration tests", keynuker_go_common.EnvVarKeyNukerTestIntegrationTestsEnabled)

	enabled := IntegrationTestsEnabled()

	if !enabled {
		t.Skip(errMsg)
	}

}

func IntegrationTestsEnabled() bool {

	testIntegrationTestsEnabled, ok := os.LookupEnv(keynuker_go_common.EnvVarKeyNukerTestIntegrationTestsEnabled)
	if !ok {
		return false
	}

	if strings.ToLower(testIntegrationTestsEnabled) != "true" {
		return false
	}

	return true
}

func ArtificialErrorInjection() bool {
	return ArtificialErrorInjectionEnabled
}

func SetArtificialErrorInjection(val bool) {
	ArtificialErrorInjectionEnabled = val
}

func FindAwsAccountArtificialError(targetAwsAccounts []TargetAwsAccount, monitorAwsAccessKeyId string) (TargetAwsAccount, error) {

	// Simulate catastrophic error fixed in commit a6ccbfcfc5d3fec24909bd3ce14b2c000a183d9b
	return TargetAwsAccount{}, fmt.Errorf("Could not find TargetAwsAccount with monitorAwsAccessKeyId: |%v| in %+v", monitorAwsAccessKeyId, targetAwsAccounts)

}

func GetTargetAwsAccountsFromEnv() (targetAwsAccounts []TargetAwsAccount, err error) {

	targetAwsAccountsRaw, ok := os.LookupEnv(keynuker_go_common.EnvVarKeyNukerTestTargetAwsAccounts)
	if !ok {
		return targetAwsAccounts, fmt.Errorf("You must define environment variable %s to run this test", keynuker_go_common.EnvVarKeyNukerTestTargetAwsAccounts)
	}

	targetAwsAccounts = []TargetAwsAccount{}

	err = json.Unmarshal([]byte(targetAwsAccountsRaw), &targetAwsAccounts)
	if err != nil {
		return targetAwsAccounts, fmt.Errorf("Unexpected Error: %v", err)
	}

	if len(targetAwsAccounts) == 0 {
		return targetAwsAccounts, fmt.Errorf("No target aws accounts")
	}

	return targetAwsAccounts, nil

}

func GetIntegrationTestAwsCredentials() (accessKey, secretAccessKey string, err error) {

	accessKey, ok := os.LookupEnv(keynuker_go_common.EnvVarKeyNukerTestIntegrationAccessKey)
	if !ok {
		return "", "", fmt.Errorf("You must define environment variable %s to run this test", keynuker_go_common.EnvVarKeyNukerTestIntegrationAccessKey)
	}

	secretAccessKey, okSecretKey := os.LookupEnv(keynuker_go_common.EnvVarKeyNukerTestIntegrationSecretAccessKey)
	if !okSecretKey {
		return "", "", fmt.Errorf("You must define environment variable %s to run this test", keynuker_go_common.EnvVarKeyNukerTestIntegrationSecretAccessKey)
	}

	return accessKey, secretAccessKey, nil

}

func GetIntegrationGithubApiBaseUrl() string {
	githubApiBaseUrl, ok := os.LookupEnv(keynuker_go_common.EnvVarKeyNukerTestIntegrationGithubApiBaseUrl)
	if !ok {
		return ""
	}
	return githubApiBaseUrl
}

func GetIntegrationTestTargetAwsAccountsFromEnv() (targetAwsAccounts []TargetAwsAccount, err error) {

	accessKey, secretAccessKey, err := GetIntegrationTestAwsCredentials()

	targetAwsAccounts = []TargetAwsAccount{
		TargetAwsAccount{
			AwsCredentials{
				AwsAccessKeyId:     accessKey,
				AwsSecretAccessKey: secretAccessKey,
			},
			TargetAwsAccountAssumeRole{},
		},
	}

	return targetAwsAccounts, nil

}

func GetGithubOrgsFromEnv() (githubOrgs []string, err error) {

	githubOrgsRaw, ok := os.LookupEnv(keynuker_go_common.EnvVarKeyNukerTestGithubOrgs)
	if !ok {
		return []string{}, fmt.Errorf("You must define environment variable %s to run this test", keynuker_go_common.EnvVarKeyNukerTestGithubOrgs)
	}

	githubOrgs = []string{}

	err = json.Unmarshal([]byte(githubOrgsRaw), &githubOrgs)
	if err != nil {
		return githubOrgs, fmt.Errorf("Unexpected Error: %v", err)
	}

	if len(githubOrgs) == 0 {
		return githubOrgs, fmt.Errorf("No github orgs")
	}

	return githubOrgs, nil

}
