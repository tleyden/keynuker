// Licensed to the Apache Software Foundation (ASF) under one or more contributor license agreements;
// and to You under the Apache License, Version 2.0.  See LICENSE in project root for full license + copyright.

package keynuker_github

import (
	"context"
	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
)

type GithubClientWrapper struct {
	AccessToken string
	ApiClient   *github.Client
}

func NewGithubClientWrapper(accessToken string) *GithubClientWrapper {

	ctx := context.Background()

	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: accessToken},
	)
	tc := oauth2.NewClient(ctx, ts)

	client := github.NewClient(tc)

	return &GithubClientWrapper{
		AccessToken: accessToken,
		ApiClient:   client,
	}
}
