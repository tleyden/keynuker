// Licensed to the Apache Software Foundation (ASF) under one or more contributor license agreements;
// and to You under the Apache License, Version 2.0.  See LICENSE in project root for full license + copyright.

package keynuker

import (
	"testing"
	"os"
	"github.com/tleyden/keynuker/keynuker-go-common"
	"fmt"
	"strings"
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