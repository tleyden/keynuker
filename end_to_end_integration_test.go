// Licensed to the Apache Software Foundation (ASF) under one or more contributor license agreements;
// and to You under the Apache License, Version 2.0.  See LICENSE in project root for full license + copyright.

package keynuker

import (
	"testing"
	"github.com/couchbaselabs/go.assert"
)

// - Create a test AWS user w/ minimal permissions
// - Loop over leaked key scenarios
//    - Create AWS key #1
//    - Scenario 1: Commit and push text file to github private repo w/ leaked key
//    - Invoke keynuker
//    - Verify that the key is nuked
//    - Create AWS key #2
//    - Scenario 1: Create a secret gist w/ leaked key
// - Verify that all leaked keys were nuked
// - Cleanup test user
// - Cleanup other residue (gists, etc)
func TestEndToEndIntegration(t *testing.T) {

	SkipIfIntegrationsTestsNotEnabled(t)

	SetArtificialErrorInjection(true)

	testIAMUsername := "KeynukerTestEndToEndIntegration"

	CreateTestIAMUserMinimalPermissions(testIAMUsername)

	keyLeakScenarios := GetEndToEndKeyLeakScenarios()
	for _, keyLeakScenario := range keyLeakScenarios {

		awsKey := CreateKeyToLeak(testIAMUsername)
		if err := keyLeakScenario.Leak(awsKey); err != nil {
			t.Fatalf("Error running testScenario: %v", err)
		}

		RunKeyNuker()

		nuked, err := VerifyKeyNuked(awsKey)
		if err != nil {
			t.Fatalf("Error verifying key was nuked: %v", err)
		}

		assert.True(t, nuked)

		keyLeakScenario.Cleanup()


	}

	DeleteIAMUser(testIAMUsername)


}
