// Licensed to the Apache Software Foundation (ASF) under one or more contributor license agreements;
// and to You under the Apache License, Version 2.0.  See LICENSE in project root for full license + copyright.

package keynuker

import (
	"encoding/json"
	"log"
	"os"
	"testing"

	"github.com/google/go-github/github"
	"github.com/stretchr/testify/assert"
	"github.com/tleyden/keynuker/keynuker-go-common"
)

func TestLookupGithubUsersAwsKeys(t *testing.T) {

	SkipIfIntegrationsTestsNotEnabled(t)

	dbHost, ok := os.LookupEnv(keynuker_go_common.EnvVarKeyNukerTestDbHost)
	if !ok {
		t.Skip("You must define environment variable %s to run this test", keynuker_go_common.EnvVarKeyNukerTestDbHost)
	}

	dbName, ok := os.LookupEnv(keynuker_go_common.EnvVarKeyNukerTestDbName)
	if !ok {
		t.Skip("You must define environment variable %s to run this test", keynuker_go_common.EnvVarKeyNukerTestDbName)
	}

	dbUsername, ok := os.LookupEnv(keynuker_go_common.EnvVarKeyNukerTestDbUsername)
	if !ok {
		t.Skip("You must define environment variable %s to run this test", keynuker_go_common.EnvVarKeyNukerTestDbUsername)
	}

	dbPassword, ok := os.LookupEnv(keynuker_go_common.EnvVarKeyNukerTestDbPassword)
	if !ok {
		t.Skip("You must define environment variable %s to run this test", keynuker_go_common.EnvVarKeyNukerTestDbPassword)
	}

	params := ParamsLookupGithubUsersAwsKeys{
		Host:        dbHost,
		Username:    dbUsername,
		Password:    dbPassword,
		DbName:      dbName,
		KeyNukerOrg: "default",
	}

	docWrapper, err := LookupGithubUsersAwsKeys(params)
	if err != nil {
		log.Printf("Error: %v", err)
	}
	assert.True(t, err == nil)

	docWrapperBytes, err := json.MarshalIndent(docWrapper, "", "    ")
	assert.True(t, err == nil)
	log.Printf("docWrapperBytes: %v", string(docWrapperBytes))

}

type TestReproKivikIssueDoc struct {
	GithubUsers []*github.User
}

type Params struct {
	Doc map[string]interface{}
}

type GithubUsersDoc struct {
	Doc struct {
		GithubUsers []*github.User
	}
}
