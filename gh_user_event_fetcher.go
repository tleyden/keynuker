// Licensed to the Apache Software Foundation (ASF) under one or more contributor license agreements;
// and to You under the Apache License, Version 2.0.  See LICENSE in project root for full license + copyright.

package keynuker

import (
	"context"
	"io/ioutil"
	"net/http"
	"time"

	"bytes"

	"log"

	"io"
	"strings"

	"encoding/base64"
	"fmt"
	"github.com/google/go-github/github"
	"github.com/tleyden/keynuker/keynuker-go-common"
)

//go:generate goautomock -template=testify -o "github_user_event_fetcher_mock.go" GithubUserEventFetcher

// Abstract the calls to the github API in an interface for dependency injection / mocking purposes
type GithubUserEventFetcher interface {

	// Given a github username and filtering parameters, fetch events from the user event stream
	FetchUserEvents(ctx context.Context, fetchUserEventsInput FetchUserEventsInput) ([]*github.Event, error)

	// Given a specific github event (eg, a commit), get the actual content for that event to be scanned for aws keys
	FetchDownstreamContent(ctx context.Context, userEvent *github.Event) (content []byte, err error)
}

type GoGithubUserEventFetcher struct {
	*GithubClientWrapper
}

// Input parameters for the github user event fetcher, which include filtering params such as checkpoint filtering
type FetchUserEventsInput struct {

	// The github username
	Username string

	// For checkpointing purposes, only consider events that are _after_ this timestamp
	SinceEventTimestamp *time.Time

	// For checkpointing purposes.  Ignore events with same ID as checkpoint.  (note: this could
	// eventually replace the time based checkpointing)
	CheckpointID string
}

func (f FetchUserEventsInput) MatchesCheckpointID(event *github.Event) bool {
	if event == nil {
		return false
	}

	return *event.ID == f.CheckpointID
}

// If you want to use the default github API (as opposed to github enterprise), pass
// in an empty string for the githubApiBaseUrl
func NewGoGithubUserEventFetcher(accessToken, githubApiBaseUrl string) *GoGithubUserEventFetcher {

	return &GoGithubUserEventFetcher{
		GithubClientWrapper: NewGithubClientWrapper(accessToken, githubApiBaseUrl),
	}
}

func eventCreatedAtBefore(event *github.Event, sinceEventTimestamp *time.Time) bool {
	return (*event.CreatedAt).Before(*sinceEventTimestamp)
}

func (guef GoGithubUserEventFetcher) FetchUserEvents(ctx context.Context, fetchUserEventsInput FetchUserEventsInput) ([]*github.Event, error) {

	publicOnly := false
	curApiResultPage := 0
	events := []*github.Event{}

	// Loop over all pages returned by API and accumulate events
	// TODO: #1 needs to also collect github gists
	// TODO: #2 filter out any events that aren't in EventTypesToInclude
	// TODO: #3 the code to loop over all pages could be extracted out into a re-usable wrappr function
	// TODO: #4 what should publicOnly be set to?  Seems like it depends on the permissions of the accesstoken
	// TODO:    and it would be good to grab private events if the access token gives enough permissions

	for {

		listOptions := &github.ListOptions{
			PerPage: MaxPerPage,
			Page:    curApiResultPage,
		}

		eventsPerPage, response, err := guef.ApiClient.Activity.ListEventsPerformedByUser(
			ctx,
			fetchUserEventsInput.Username,
			publicOnly,
			listOptions,
		)

		if err != nil {
			return nil, err
		}

		// Loop over events and append to result
		for _, event := range eventsPerPage {
			events = append(events, event)
		}

		if response.NextPage <= curApiResultPage {
			// Last page, we're done
			break
		}

		curApiResultPage += 1

	}

	return events, nil

}

func (guef GoGithubUserEventFetcher) FetchDownstreamContent(ctx context.Context, userEvent *github.Event) (content []byte, err error) {

	payload, err := userEvent.ParsePayload()
	if err != nil {
		return nil, err
	}

	switch v := payload.(type) {
	case *github.PullRequestEvent:

		content, err := guef.FetchUrlContent(ctx, v.PullRequest.GetPatchURL())
		if err != nil {
			return nil, err
		}
		return content, nil

	case *github.PushEvent:

		buffer := bytes.Buffer{}

		maxCommitsPerPushEvent := 20

		if strings.Contains(*v.Ref, keynuker_go_common.KeyNukerIntegrationTestBranch) {
			// skip this since as an experiment
			log.Printf("Skipping push event %v since it's on %v testing branch", *v.PushID, keynuker_go_common.KeyNukerIntegrationTestBranch)
			return []byte(""), nil
		}

		commits := v.Commits
		for _, commit := range commits {
			log.Printf("Getting content for commit: %v url: %v", *commit.SHA, commit.GetURL())

			// split "org/reponame" into separate strings (["org", "reponame"])
			repoNameComponents := strings.Split(*userEvent.Repo.Name, "/")
			username := repoNameComponents[0]
			repoName := repoNameComponents[1]

			repoCommit, _, err := guef.ApiClient.Repositories.GetCommit(
				ctx,
				username,
				repoName,
				*commit.SHA,
			)
			if err != nil {
				return nil, fmt.Errorf("Error getting content for commit: %v url: %v.  Error: %v", *commit.SHA, commit.GetURL(), err)
			}

			for _, repoCommitFile := range repoCommit.Files {
				buffer.WriteString(repoCommitFile.GetPatch())
				if repoCommitFile.Patch == nil {

					// This means its binary data or larger than 1 MB, call separate API to fetch
					log.Printf("Warning: commit %+v has empty patch data.  Either binary data, or greater than 1MB", repoCommitFile)

					blob, _, err := guef.ApiClient.Git.GetBlob(
						ctx,
						username,
						repoName,
						*repoCommitFile.SHA,
					)
					if err != nil {
						return nil, fmt.Errorf("Error getting content for commit file: %+v via blob api.  Error: %v", repoCommitFile, err)
					}
					if *blob.Encoding != "base64" {
						return nil, fmt.Errorf("Unexpected encoding for blob commit file: %+v via blob api.  Encoding: %v", repoCommitFile, *blob.Encoding)
					}
					if *blob.Size > keynuker_go_common.MaxSizeBytesBlobContent {
						log.Printf("Warning: skipping blob from commit file %+v, since size > max size (%d)", repoCommitFile, keynuker_go_common.MaxSizeBytesBlobContent)
						continue
					}

					decodedBlobContent, err := base64.StdEncoding.DecodeString(blob.GetContent())
					if err != nil {
						return nil, fmt.Errorf("Unexpected decoding base64 for blob commit file: %+v via blob api.  Err: %v", repoCommitFile, err)
					}
					buffer.Write(decodedBlobContent)

				}

			}

			buffer.Write(content)
		}

		// If more than 20 commits for this PushEvent, fetch content for the remaining commits.
		// Example PushEvent w/ more than 20 commits: https://gist.github.com/tleyden/68d972b02b2b9306fa6e2eb26310b751
		if *v.Size > maxCommitsPerPushEvent {

			log.Printf("PushEvent %v has > 20 commits but this API call only returns 20.  Making separate API call.", *v.PushID)

			// Fetch the rest of the commits for this push event and append downstream content to buffer
			_, err := guef.FetchCommitsForPushEvent(ctx, userEvent, v, &buffer)
			if err != nil {
				log.Printf("Warning: Error fetching additional commits for push event: %v.  Error: %v", *v.PushID, err)
				return nil, err
			}

		}

		return buffer.Bytes(), nil

	default:
		// TODO: add case statements for all events that should be scanned
		return []byte(*userEvent.RawPayload), nil

	}

	return nil, nil
}

// Since PushEvents only contain 20 commits max, this fetches the remaining commits and writes the content to the
// writer passed in.  For example, pushEvent.Size might indicate that there were 100 commits in the push events,
// and so the remaining 80 commits will need to be scanned.
//
// Github API: https://developer.github.com/v3/repos/commits/
func (guef GoGithubUserEventFetcher) FetchCommitsForPushEvent(
	ctx context.Context, userEvent *github.Event, pushEvent *github.PushEvent, w io.Writer) (completed bool, err error) {

	log.Printf("FetchCommitsForPushEvent: %v.  Number of total commits in push event: %d",
		*pushEvent.PushID, *pushEvent.Size)

	numCommitsToScan := *pushEvent.Size
	numCommitsScanned := 0
	resultPage := 0

	// One large PushEvent with > 900 commits on https://api.github.com/repos/nolanlawson/mastodon was killing KeyNuker in two ways:
	// 1. Using up all the GithubApi requests from the allotted 5K / hour
	// 2. Blowing up the memory usage since it is grossly inefficent and rolls up all content from each event into a buffer
	// To limit the damage, limit the number of commits scanned in a single push event to approximately 220 (assuming MaxPerPage is 100)
	maxPages := 2

	// The inline commits in the push event don't need to be re-scanned, so build a map of their sha's
	inlineCommits := map[string]struct{}{}
	for _, commit := range pushEvent.Commits {
		inlineCommits[*commit.SHA] = struct{}{}
	}

	// Keep listing commits on that branch until we go back far enough to reach the last commit in pushEvent.size

	for {

		if numCommitsScanned >= numCommitsToScan {
			// done
			return true, nil
		}

		if resultPage > maxPages {
			notScanned := numCommitsToScan - numCommitsScanned
			log.Printf("WARNING: not scanning %d commits, due to current limitations.  See https://github.com/tleyden/keynuker/issues/24", notScanned)
			return false, nil
		}

		commitListOptions := &github.CommitsListOptions{
			SHA: *pushEvent.Head,
			ListOptions: github.ListOptions{
				PerPage: MaxPerPage,
				Page:    resultPage,
			},
		}
		log.Printf("Listing additional commits: %+v for repo: %s", commitListOptions, *userEvent.Repo.Name)

		// "tleyden/keynuker" -> ["tleyden", "keynuker"] -> "keynuker"
		repoNameAndUsername := *userEvent.Repo.Name
		repoNameAndUsernameComponents := strings.Split(repoNameAndUsername, "/")
		owner := repoNameAndUsernameComponents[0]
		repoName := repoNameAndUsernameComponents[1]

		additionalCommits, resp, err := guef.ApiClient.Repositories.ListCommits(
			ctx,
			owner,
			repoName,
			commitListOptions,
		)
		if err != nil {
			return false, fmt.Errorf("Error calling ApiClient.Repositories.ListCommits: %v", err)
		}

		for _, additionalCommit := range additionalCommits {

			_, foundInlineCommit := inlineCommits[*additionalCommit.SHA]
			if foundInlineCommit {
				numCommitsScanned += 1
				continue
			}

			// The commit struct returned from ListCommits() will just be a stub without content.
			// Call GetCommit() to get the patch content of the commit, as long as it's < 1 MB.
			repoCommit, _, err := guef.ApiClient.Repositories.GetCommit(
				ctx,
				owner,
				repoName,
				*additionalCommit.SHA,
			)
			if err != nil {
				return false, fmt.Errorf("Error getting additional content for commit: %v url: %v.  Error: %v", *additionalCommit.SHA, additionalCommit.GetURL(), err)
			}

			// Loop over the files in the commit and append the content to the writer
			for _, repoCommitFile := range repoCommit.Files {
				w.Write([]byte(repoCommitFile.GetPatch()))
				if repoCommitFile.Patch == nil {

					// This means its binary data or larger than 1 MB, call separate API to fetch
					log.Printf("Warning: additional commit %+v has empty patch data.  Either binary data, or greater than 1MB", repoCommitFile)

					blob, _, err := guef.ApiClient.Git.GetBlob(
						ctx,
						owner,
						repoName,
						*repoCommitFile.SHA,
					)
					if err != nil {
						return false, fmt.Errorf("Error getting additional content for commit file: %+v via blob api.  Error: %v", repoCommitFile, err)
					}
					if *blob.Encoding != "base64" {
						return false, fmt.Errorf("Unexpected encoding for additional blob commit file: %+v via blob api.  Encoding: %v", repoCommitFile, *blob.Encoding)
					}
					if *blob.Size > keynuker_go_common.MaxSizeBytesBlobContent {
						log.Printf("Warning: skipping additional blob from commit file %+v, since size > max size (%d)", repoCommitFile, keynuker_go_common.MaxSizeBytesBlobContent)
						continue
					}

					decodedBlobContent, err := base64.StdEncoding.DecodeString(blob.GetContent())
					if err != nil {
						return false, fmt.Errorf("Unexpected decoding base64 for additional blob commit file: %+v via blob api.  Err: %v", repoCommitFile, err)
					}
					w.Write(decodedBlobContent)

				}

			}

			numCommitsScanned += 1

			if numCommitsScanned >= numCommitsToScan {
				// done
				return true, nil
			}

		}

		log.Printf("resp %+v.  NextPage: %v.  LastPage: %v", resp, resp.NextPage, resp.LastPage)

		resultPage = resp.NextPage

	}

	return true, nil
}

func (guef GoGithubUserEventFetcher) FetchUrlContent(ctx context.Context, url string) (content []byte, err error) {

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	q := req.URL.Query()
	q.Add("access_token", guef.AccessToken)
	req.URL.RawQuery = q.Encode()

	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	resp_body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return resp_body, nil

}
