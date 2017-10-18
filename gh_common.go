// Licensed to the Apache Software Foundation (ASF) under one or more contributor license agreements;
// and to You under the Apache License, Version 2.0.  See LICENSE in project root for full license + copyright.

package keynuker

import (
	"context"
	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
	"net/url"
	"fmt"
)

type GithubClientWrapper struct {
	AccessToken string
	ApiClient   *github.Client
}

// If you want to use the default github API (as opposed to github enterprise), pass
// in an empty string for the githubApiBaseUrl
func NewGithubClientWrapper(accessToken, githubApiBaseUrl string) *GithubClientWrapper {

	ctx := context.Background()

	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: accessToken},
	)
	tc := oauth2.NewClient(ctx, ts)
	
	client := github.NewClient(tc)

	// If an alternative github API base url was given, use that
	if githubApiBaseUrl != "" {
		baseUrl, err := url.Parse(githubApiBaseUrl)
		if err != nil {
			panic(fmt.Sprintf("Invalid Github API url given: %v", githubApiBaseUrl))
		}
		client.BaseURL = baseUrl
	}

	return &GithubClientWrapper{
		AccessToken: accessToken,
		ApiClient:   client,
	}
}
