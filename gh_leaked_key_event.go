// Licensed to the Apache Software Foundation (ASF) under one or more contributor license agreements;
// and to You under the Apache License, Version 2.0.  See LICENSE in project root for full license + copyright.

package keynuker

import (
	"github.com/aws/aws-sdk-go/service/iam"
	"github.com/google/go-github/github"
	"time"
)

type LeakedKeyEvent struct {
	AccessKeyMetadata iam.AccessKeyMetadata

	GithubUser *github.User

	GithubEvent *github.Event

	NearbyContent []byte
}

type NukedKeyEvent struct {
	LeakedKeyEvent LeakedKeyEvent

	DeleteAccessKeyOutput *iam.DeleteAccessKeyOutput

	NukedOn time.Time
}
