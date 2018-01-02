// Licensed to the Apache Software Foundation (ASF) under one or more contributor license agreements;
// and to You under the Apache License, Version 2.0.  See LICENSE in project root for full license + copyright.

package keynuker

import (
	"context"
	"encoding/json"
	"fmt"

	_ "github.com/go-kivik/couchdb" // The CouchDB driver
	"github.com/go-kivik/kivik"
	"github.com/pkg/errors"
	"github.com/tleyden/keynuker/keynuker-go-common"
)

// For a given KeyNuker org (or default org), look up:
//   - All known github users produced by github-user-aggregator
//   - All known aws keys produced by fetch-aws-keys
// And combine them in a single document and emit the document

func LookupGithubUsersAwsKeys(params ParamsLookupGithubUsersAwsKeys) (docWrapper DocumentWrapperLookupGithubUsersAwsKeys, err error) {

	ctx := context.Background()

	if err := params.Validate(); err != nil {
		return docWrapper, err
	}

	dataSourceName := fmt.Sprintf(
		"https://%s:%s@%s",
		params.Username,
		params.Password,
		params.Host,
	)

	client, err := kivik.New(ctx, "couch", dataSourceName)
	if err != nil {
		return docWrapper, err
	}

	db, err := client.DB(ctx, params.DbName)
	if err != nil {
		return docWrapper, err
	}

	// Get doc id for github users doc in a keynuker org
	docIdGithubUsers := keynuker_go_common.GenerateDocId(
		keynuker_go_common.DocIdPrefixGithubUsers,
		params.KeyNukerOrg,
	)

	options := kivik.Options{}
	rowGithubUsers := db.Get(ctx, docIdGithubUsers, options)

	docGithubUsers := DocumentWithGithubUsers{}
	if err := rowGithubUsers.ScanDoc(&docGithubUsers); err != nil {
		return docWrapper, err
	}
	docWrapper.GithubUsers = docGithubUsers.GithubUsers

	// Get doc with aws keys
	docIdAwsKeys := keynuker_go_common.GenerateDocId(
		keynuker_go_common.DocIdPrefixAwsKeys,
		params.KeyNukerOrg,
	)
	rowAwsKeys := db.Get(ctx, docIdAwsKeys, options)

	docAwsKeys := DocumentWithAwsKeys{}
	if err := rowAwsKeys.ScanDoc(&docAwsKeys); err != nil {
		return docWrapper, err
	}
	docWrapper.AccessKeyMetadata = docAwsKeys.AccessKeyMetadata

	// Lookup github checkpoints doc
	docIdGithubEventCheckpoints := keynuker_go_common.GenerateDocId(
		keynuker_go_common.DocIdPrefixGithubEventCheckpoints,
		params.KeyNukerOrg,
	)
	rowGithubEventCheckpoints := db.Get(ctx, docIdGithubEventCheckpoints, options)

	if rowGithubEventCheckpoints != nil {
		docGithubEventCheckpoints := DocumentWithGithubEventCheckpoints{}
		if err := rowGithubEventCheckpoints.ScanDoc(&docGithubEventCheckpoints); err != nil {
			return docWrapper, err
		}
		docWrapper.GithubEventCheckpoints = docGithubEventCheckpoints.GithubEventCheckpoints
	}

	return docWrapper, nil

}

type ParamsLookupGithubUsersAwsKeys struct {

	// This is the name of the KeyNuker "org/tenant".  Defaults to "default", but allows to be extended multi-tenant.
	KeyNukerOrg string

	// DB connection params
	Username string
	Password string
	Host     string
	DbName   string
}

func (p ParamsLookupGithubUsersAwsKeys) Validate() error {
	if p.Host == "" {
		return errors.Errorf("No DB Host specified in params")
	}
	if p.DbName == "" {
		return errors.Errorf("No DB name specified in params")
	}
	return nil

}

type DocumentWrapperLookupGithubUsersAwsKeys struct {

	// A list of github users
	GithubUsers *json.RawMessage

	// AWS access keys to scan for
	AccessKeyMetadata *json.RawMessage

	// Github event checkpoint which represent the last scanned github event for each known github user
	GithubEventCheckpoints *json.RawMessage
}

type DocumentWithGithubUsers struct {
	GithubUsers *json.RawMessage
}

type DocumentWithAwsKeys struct {
	AccessKeyMetadata *json.RawMessage
}

type DocumentWithGithubEventCheckpoints struct {
	GithubEventCheckpoints *json.RawMessage
}
