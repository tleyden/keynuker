// Licensed to the Apache Software Foundation (ASF) under one or more contributor license agreements;
// and to You under the Apache License, Version 2.0.  See LICENSE in project root for full license + copyright.

package keynuker

import "github.com/google/go-github/github"

// Githhub User Login -> Last Processed Github Event
type GithubEventCheckpoints map[string]*github.Event

func (gec GithubEventCheckpoints) CheckpointForUser(user *github.User) (checkpoint *github.Event, found bool) {
	checkpoint, found = gec[*user.Login]
	return checkpoint, found
}
