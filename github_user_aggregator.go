// Licensed to the Apache Software Foundation (ASF) under one or more contributor license agreements;
// and to You under the Apache License, Version 2.0.  See LICENSE in project root for full license + copyright.

package keynuker

import (
	"github.com/google/go-github/github"
)


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
