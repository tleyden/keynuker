// Licensed to the Apache Software Foundation (ASF) under one or more contributor license agreements;
// and to You under the Apache License, Version 2.0.  See LICENSE in project root for full license + copyright.

package keynuker

import (
	"context"
	"fmt"

	"github.com/google/go-github/github"
	"github.com/tleyden/keynuker/keynuker-go-common"
)

// Given a list of github orgs, aggregate all of the users that belong in the orgs
// and emit a json to stdout with those users.
// Intended to be run as an OpenWhisk Action
func AggregateGithubUsers(params ParamsGithubUserAggregator) (DocumentWrapperGithubUserAggregator, error) {

	// Document ID for output parameter, which allows downstream job to stick into a DB
	docId := keynuker_go_common.GenerateDocId(
		keynuker_go_common.DocIdPrefixGithubUsers,
		params.KeyNukerOrg,
	)

	// Create a github user aggregator helper
	ghUserAggregator := NewGithubUserAggregator(
		params.GithubOrgs,
		params.GithubAccessToken,
	)

	// Call out to Github API to get aggregated members of orgs
	users, err := ghUserAggregator.ListMembers(context.Background())
	if err != nil {
		return DocumentWrapperGithubUserAggregator{}, fmt.Errorf("Error listing members for orgs: %v.  Error: %v", params.GithubOrgs, err)
	}

	// Create result doc
	doc := DocumentGithubUserAggregator{
		Id:          docId,
		GithubUsers: users,
	}

	// Create result doc wrapper
	docWrapper := DocumentWrapperGithubUserAggregator{
		Doc:   doc,
		DocId: docId,
	}
	return docWrapper, nil

}

type ParamsGithubUserAggregator struct {

	// This is the name of the KeyNuker "org/tenant".  Defaults to "default", but allows to be extended multi-tenant.
	KeyNukerOrg string

	// The github access token, which needs "read:org" permissions in order to read the concealed "non-public"
	// members of the orgs
	GithubAccessToken string

	// A list of github organizations, eg ["acme", "acme-labs", ...]
	GithubOrgs []string
}

type DocumentGithubUserAggregator struct {
	Id          string `json:"_id"`
	GithubUsers []*github.User
}

type DocumentWrapperGithubUserAggregator struct {
	// Serialize into a form that the cloudant db adapter expects
	Doc   DocumentGithubUserAggregator `json:"doc"`
	DocId string                       `json:"docid"`
}
