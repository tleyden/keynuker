// Licensed to the Apache Software Foundation (ASF) under one or more contributor license agreements;
// and to You under the Apache License, Version 2.0.  See LICENSE in project root for full license + copyright.

package keynuker

import (
	"encoding/json"
	"log"
	"os"
	"testing"

	"context"
	"fmt"
	"net/http"
	"regexp"

	"bytes"
	"github.com/flimzy/kivik"
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

// See https://github.com/flimzy/kivik/issues/156
func TestReproKivikIssue(t *testing.T) {

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

	ghUsername := "torvalds"

	ctx := context.Background()

	// Connect to DB
	dataSourceName := fmt.Sprintf(
		"http://%s:%s@%s",
		dbUsername,
		dbPassword,
		dbHost,
	)

	client, _ := kivik.New(ctx, "couch", dataSourceName)
	db, _ := client.DB(ctx, dbName)

	// Get a github user json
	githubClient := github.NewClient(http.DefaultClient)
	user, _, _ := githubClient.Users.Get(ctx, ghUsername)

	users := []*github.User{}
	users = append(users, user)
	doc := TestReproKivikIssueDoc{
		GithubUsers: users,
	}

	// Save to db
	db.Put(ctx, ghUsername, doc)

	// Read from db
	var usersFromDb TestReproKivikIssueDoc
	row, _ := db.Get(ctx, ghUsername, kivik.Options{})

	// Marshal into github.user
	row.ScanDoc(&usersFromDb)

	docWrapperBytes, err := json.MarshalIndent(usersFromDb, "", "    ")
	assert.True(t, err == nil)
	log.Printf("docWrapperBytes: %v", string(docWrapperBytes))

}

type Params struct {
	Doc map[string]interface{}
}

type GithubUsersDoc struct {
	Doc struct {
		GithubUsers []*github.User
	}
}

// See https://github.com/flimzy/kivik/issues/156
// https://play.golang.org/p/pfNMH4UzK_  <-- works!?  is that b/c they are using a newer version of Go?
func TestReproUnmarshalVsDecodeIssue(t *testing.T) {

	var params Params

	docJsonRawMessage := []byte(`{
   "doc":{
      "GithubUsers":[
         {
            "id":3411441,
            "type":"User"
         }
      ],
      "_id":"github_users_default"
   }
}`)

	// Using json.Unmarshal -- doesn't work

	json.Unmarshal(docJsonRawMessage, &params)

	marshalled, _ := json.MarshalIndent(params, "", "    ")
	log.Printf("marshalled: %v", string(marshalled))

	r, _ := regexp.Compile("3411441,")

	if !r.Match(marshalled) {
		log.Printf("Was not marshalled as expected")
	}

	var githubUsersDoc GithubUsersDoc
	err := json.Unmarshal(marshalled, &githubUsersDoc)
	if err != nil {
		log.Printf("Couldn't unmarshal into github user object: %v", err)
	}
	log.Printf("user: %+v", githubUsersDoc)

	// Using json.NewDecoder() + UseNumber()

	var params2 Params
	buffer := bytes.NewBuffer(docJsonRawMessage)
	d := json.NewDecoder(buffer)
	d.UseNumber()
	d.Decode(&params2)

	marshalled2, _ := json.MarshalIndent(params2, "", "    ")
	if !r.Match(marshalled2) {
		log.Printf("Was not marshalled as expected")
	}
	log.Printf("marshalled2: %v", string(marshalled2))

	err = json.Unmarshal(marshalled2, &githubUsersDoc)
	if err != nil {
		log.Printf("Couldn't unmarshal into github user object: %v", err)
	}
	log.Printf("user: %+v", githubUsersDoc)

}
