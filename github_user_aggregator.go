// Licensed to the Apache Software Foundation (ASF) under one or more contributor license agreements;
// and to You under the Apache License, Version 2.0.  See LICENSE in project root for full license + copyright.

package keynuker

import (
	"context"

	"github.com/google/go-github/github"
	"github.com/pkg/errors"
	"github.com/tleyden/keynuker/keynuker-go-common"
	"golang.org/x/oauth2"
)

// Looks up and aggregates all users in the given Github Organizations and returns the de-deduplicated list
// of users

// TODO: --------------------- Multiple Access tokens / Rate limit handling
// TODO: To lower chance that requests are rejected by github for exceeding rate limit, this should
// TODO: take a slice of access tokens, and it should choose the access token with the most requests available.

const (
	MaxPerPage = 100
)

type GithubUserAggregator struct {
	AccessToken string
	GithubOrgs  []string
	ApiClient   *github.Client
}

func NewGithubUserAggregator(orgs []string, accessToken string) *GithubUserAggregator {

	ctx := context.Background()

	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: accessToken},
	)
	tc := oauth2.NewClient(ctx, ts)

	client := github.NewClient(tc)

	return &GithubUserAggregator{
		AccessToken: accessToken,
		GithubOrgs:  orgs,
		ApiClient:   client,
	}

}

func (gua GithubUserAggregator) ListMembers(ctx context.Context) ([]*github.User, error) {


	// The resulting list of users aggregated across all orgs
	users := []*github.User{}

	for _, org := range gua.GithubOrgs {

		usersPerOrg, err := gua.ListMembersForOrg(ctx, org)
		if err != nil {
			return nil, err
		}

		users = appendUsersDeDupe(users, usersPerOrg)

	}

	return users, nil

}




// Get a compacted list of the users that only contains the user data that concerns the keynuker application
// TODO: this should convert to a keynuker.GithubUser (doesn't exist yet) to increase compile time checking
func (gua GithubUserAggregator) CompactedUsers(users []*github.User) []*github.User {
	resultUsers := []*github.User{}
	for _, user := range users {
		resultUser := &github.User{}
		resultUser.Login = user.Login
		resultUsers = append(resultUsers, resultUser)
	}
	return resultUsers
}

func (gua GithubUserAggregator) ListMembersForOrg(ctx context.Context, org string) ([]*github.User, error) {

	// Keep track of which page we are on when iterating over API results
	curApiResultPage := 1

	// The resulting list of users for this org
	users := []*github.User{}

	for {

		opts := github.ListMembersOptions{
			ListOptions: github.ListOptions{
				PerPage: MaxPerPage,
				Page:    curApiResultPage,
			},
		}

		// TODO: ----------------- Handle non-existent orgs better
		// TODO: When I tested with non-existent orgs, it panicked with:
		// TODO: Error listing members for orgs: [foo bar].  Error: GET https://api.github.com/orgs/foo/members?per_page=100: 404 Not Found []
		// TODO: Instead of panicking, should gracefully handle the 404 error, log a warning to stderr, and skip that org.
		// TODO: ----------------- Verify access token can collect non-public concealed members
		// TODO: Make sure the Github Client is using an access token that is able to read concealed (non-public)
		// TODO:   members of the github org, otherwise it will be an incomplete list, which is a huge bug.
		// TODO:   If the access token does not have sufficient permissions, then emit warnings to the logs
		// TODO:   which will require structured logging of some sort.  (See the logging library used in
		// TODO:   github/tleyden/cecil).  One way to do it would be to make a call to
		// TODO:   https://api.github.com/user/memberships/orgs using this access token, and make sure the org is there.
		usersPerOrg, response, err := gua.ApiClient.Organizations.ListMembers(
			ctx,
			org,
			&opts,
		)
		if err != nil {
			return nil, err
		}

		users = appendUsersDeDupe(users, usersPerOrg)

		if curApiResultPage >= response.LastPage {
			// Lost page, we're done
			break
		}

		curApiResultPage = response.NextPage

	}

	return users, nil

}

// Append incoming to existing as long as the user is not already in existing, and return the result
func appendUsersDeDupe(existing, incoming []*github.User) []*github.User {

	// The resulting list of users
	users := []*github.User{}

	// Start out w/ the set of existing users
	users = append(users, existing...)

	// Map of user logins for de-dupe purposes
	existingUserLogins := map[string]struct{}{}

	// Build a map of user id's
	for _, existingUser := range existing {
		existingUserLogins[*existingUser.Login] = struct{}{}
	}

	for _, incomingUser := range incoming {

		// Make sure we don't already have this user in existing
		_, ok := existingUserLogins[*incomingUser.Login]
		if ok {
			// Already have this user in existing, skip
			continue
		}

		// Append incoming user
		users = append(users, incomingUser)

	}

	return users

}

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
		return DocumentWrapperGithubUserAggregator{}, errors.Wrapf(err, "Error listing members for orgs: %v.", params.GithubOrgs)
	}

	// Add all the individual github users from the GithubUsers param
	individualGithubUsers := params.GetGithubUsers()
	users = appendUsersDeDupe(users, individualGithubUsers)

	// Strip out unneeded fields from users
	compactedUsers := ghUserAggregator.CompactedUsers(users)

	// Create result doc
	doc := DocumentGithubUserAggregator{
		Id:          docId,
		GithubUsers: compactedUsers,
	}

	// Create result doc wrapper
	docWrapper := DocumentWrapperGithubUserAggregator{
		Doc:   doc,
		DocId: docId,
	}
	return docWrapper, nil

}

type ParamsGithubUserAggregator struct {

	// Github API URL and access token
	GithubConnectionParams

	// This is the name of the KeyNuker "org/tenant".  Defaults to "default", but allows to be extended multi-tenant.
	KeyNukerOrg string

	// A list of github organizations, eg ["acme", "acme-labs", ...]
	GithubOrgs []string

	// A list of individual github user logins you would like to monitor, which is appended to the users found from looking up the users in GithubOrgs.  Eg, ["defunkt", "torvalds"]
	GithubUsers []string
}

func (p ParamsGithubUserAggregator) GetGithubUsers() []*github.User {
	resultUsers := []*github.User{}
	for _, githubUserLogin := range p.GithubUsers {
		resultUser := &github.User{}
		resultUser.Login = &githubUserLogin
		resultUsers = append(resultUsers, resultUser)
	}
	return resultUsers
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
