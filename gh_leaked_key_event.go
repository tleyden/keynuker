// Licensed to the Apache Software Foundation (ASF) under one or more contributor license agreements;
// and to You under the Apache License, Version 2.0.  See LICENSE in project root for full license + copyright.

package keynuker

import (
	"github.com/aws/aws-sdk-go/service/iam"
	"github.com/google/go-github/github"
	"time"
)

type LeakedKeyEvent struct {
	AccessKeyMetadata FetchedAwsAccessKey

	GithubUser *github.User

	GithubEvent *github.Event

	NearbyContent []byte
}

// Is the key that was leaked the same (limited permission) key that keynuker is using to monitor aws?
// Don't nuke it, since this key has very limited permissions and is needed to nuke other keys.
// Should raise serious alarms if this ever happens.
func(l LeakedKeyEvent) LeakedKeyIsMonitorKey() bool {
	return *l.AccessKeyMetadata.AccessKeyId == l.AccessKeyMetadata.MonitorAwsAccessKeyId
}

type NukedKeyEvent struct {
	LeakedKeyEvent LeakedKeyEvent

	DeleteAccessKeyOutput *iam.DeleteAccessKeyOutput

	NukedOn time.Time
}
