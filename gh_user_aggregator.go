// Licensed to the Apache Software Foundation (ASF) under one or more contributor license agreements;
// and to You under the Apache License, Version 2.0.  See LICENSE in project root for full license + copyright.

package keynuker

import (
	"context"

	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
)

// Looks up and aggregates all users in the given Github Organizations and returns the de-deduplicated list
// of users

// WARN: don't write anything to stdout.  Since this is invoked by an OpenWhisk action, that will break the JSON output.

// TODO: --------------------- Individual Usernames
// TODO: This should also be able to take a list of individual github usernames ["tleyden", ...], and those
// TODO: would be added to the resulting list of users.  The list of orgs could be empty and the only input
// TODO: could be a list of usernames (or even a single user).  The OpenWhisk action wrapper will also need to be
// TODO: updated accordingly.
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

func (gua GithubUserAggregator) ListMembersForOrg(ctx context.Context, org string) ([]*github.User, error) {

	// Keep track of which page we are on when iterating over API results
	curApiResultPage := 0

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

		users = append(users, usersPerOrg...)

		if response.NextPage <= curApiResultPage {
			// Lost page, we're done
			break
		}

		curApiResultPage += 1

	}

	return users, nil

}

// Append incoming to existing as long as the user is not already in existing, and return the result
func appendUsersDeDupe(existing, incoming []*github.User) []*github.User {

	// The resulting list of users
	users := []*github.User{}

	// Start out w/ the set of existing users
	users = append(users, existing...)

	// Map of user id's for de-dupe purposes
	existingUserIds := map[int]struct{}{}

	// Build a map of user id's
	for _, existingUser := range existing {
		id := *existingUser.ID
		existingUserIds[id] = struct{}{}
	}

	for _, incomingUser := range incoming {

		// Make sure we don't already have this user in existing
		id := *incomingUser.ID
		_, ok := existingUserIds[id]
		if ok {
			// Already have this user in existing, skip
			continue
		}

		// Append incoming user
		users = append(users, incomingUser)

	}

	return users

}
