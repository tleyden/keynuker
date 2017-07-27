// Licensed to the Apache Software Foundation (ASF) under one or more contributor license agreements;
// and to You under the Apache License, Version 2.0.  See LICENSE in project root for full license + copyright.

package keynuker

import (
	"encoding/json"
	"log"
	"os"
	"testing"

	"github.com/couchbaselabs/go.assert"
	"github.com/tleyden/keynuker/keynuker-go-common"
)

func TestWriteDoc(t *testing.T) {

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

	docContentStr := `
		{
		   "title":"books",
		   "golang":[
			  {
				 "title":"The Go programming language",
				 "authors":[
					"Alan Donovan",
					"Brian Kernighen"
				 ],
				 "date":"2017-07-11T16:02:01Z",
				 "rating":"⭐️⭐️⭐️⭐️⭐️"
			  },
			  {
				 "title":"The way to Go",
				 "authors":[
					"Ivo B"
				 ],
				 "rating":"⭐️⭐️⭐️⭐"
			  }
		   ]
		}
	`

	doc := map[string]interface{}{}
	err := json.Unmarshal([]byte(docContentStr), &doc)
	if err != nil {
		t.Fatalf("Error unmarshalling doc.  Err: %v", err)
	}

	params := ParamsWriteDoc{
		Host:     dbHost,
		Username: dbUsername,
		Password: dbPassword,
		DbName:   dbName,
		Doc:      doc,
		DocId:    "TestDoc",
	}

	docWrapper, err := WriteDocToDb(params)
	if err != nil {
		log.Printf("Error: %v", err)
	}
	assert.True(t, err == nil)

	docWrapperBytes, err := json.MarshalIndent(docWrapper, "", "    ")
	assert.True(t, err == nil)
	log.Printf("docWrapperBytes: %v", string(docWrapperBytes))

}
