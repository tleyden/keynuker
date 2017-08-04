// Licensed to the Apache Software Foundation (ASF) under one or more contributor license agreements;
// and to You under the Apache License, Version 2.0.  See LICENSE in project root for full license + copyright.

package keynuker

import (
	"testing"
	"os"
	"github.com/tleyden/keynuker/keynuker-go-common"
	"fmt"
	"strings"
	"encoding/json"
)

var (
	ArtificialErrorInjectionEnabled = true
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