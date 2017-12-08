// Licensed to the Apache Software Foundation (ASF) under one or more contributor license agreements;
// and to You under the Apache License, Version 2.0.  See LICENSE in project root for full license + copyright.

package keynuker

import (
	"context"
	"net/http"
	"time"

	"log"

	"strings"

	"encoding/base64"

	"github.com/google/go-github/github"
	"github.com/pkg/errors"
	"github.com/tleyden/keynuker/keynuker-go-common"
)

//go:generate goautomock -template=testify -o "github_user_event_fetcher_mock.go" GithubUserEventFetcher

// Abstract the calls to the github API in an interface for dependency injection / mocking purposes
type GithubUserEventFetcher interface {

	// Given a github username and filtering parameters, fetch events from the user event stream
	FetchUserEvents(ctx context.Context, fetchUserEventsInput FetchUserEventsInput) ([]*github.Event, error)

	// Given a specific github event (eg, a commit), scan the actual content for that event for given aws keys
	ScanDownstreamContent(ctx context.Context, userEvent *github.Event, accessKeysToScan []FetchedAwsAccessKey) (leaks []FetchedAwsAccessKey, err error)
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

	curApiResultPage := 1

	eventStack := NewEventStack()

	// Loop over all pages returned by API and accumulate events

	// TODO: #1 It can stop going back if the oldest event from the most recent page of results is older than the checkpoint event.

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
			eventStack.Push(event)
		}

		if curApiResultPage >= response.LastPage {
			// Last page, we're done
			break
		}

		curApiResultPage = response.NextPage

	}

	return eventStack.PopAll(), nil

}

func (guef GoGithubUserEventFetcher) ScanContentForCommits(ctx context.Context, username, repoName string, commits []WrappedCommit, accessKeysToScan []FetchedAwsAccessKey) (leaks []FetchedAwsAccessKey, err error) {

	leaks = []FetchedAwsAccessKey{}

	for _, commit := range commits {

		log.Printf("Getting content for commit: %v url: %v", commit.Sha(), commit.Url())

		repoCommit, _, err := guef.ApiClient.Repositories.GetCommit(
			ctx,
			username,
			repoName,
			commit.Sha(),
		)
		if err != nil {
			if keynuker_go_common.IsTemporaryGithubError(err) {
				// Abort now since this will prevent checkpoint from getting pushed, and will be retried later
				// ErrorInjection: return nil instead of leaks here.  It will discard the leaks collected so far from previous commits.
				return leaks, errors.Wrapf(err, "Temporary error getting content for commit: %v url: %v", commit.Sha(), commit.Url())
			} else {
				log.Printf("Permanent error getting commit for username: %s, reponame: %s, sha: %s.  Skipping commit.  Error: %v", username, repoName, commit.Sha(), err)
				continue
			}
		}

		for _, repoCommitFile := range repoCommit.Files {
			if repoCommitFile.Patch != nil {
				// This commit file has an inline "patch" that has the delta of the content
				patchContent := []byte(repoCommitFile.GetPatch())
				log.Printf("Scanning %d bytes of content for commit: %v url: %v", len(patchContent), commit.Sha(), commit.Url())
				leaksForFile, err := Scan(accessKeysToScan, patchContent)
				if err != nil {
					log.Printf("WARNING: error scanning content for commit: %v.  Err: %v  Skipping this content.", repoCommitFile, err)
					continue
				}
				if len(leaksForFile) > 0 {
					leaks = append(leaks, leaksForFile...)
				}

			} else {

				// This commit file doesn't have an inline "patch", which means its binary data or larger than 1 MB
				// so it's necessary to call a separate API to fetch

				log.Printf("Commit %+v has empty patch data.  Either binary data, or greater than 1MB, fetching using GetBlob API", repoCommitFile)

				if repoCommitFile.SHA == nil {
					// TODO: move this code into a function called GetContentForCommit(commit) which will use GetBlob() if there
					// TODO: is a SHA, otherwise it will fallback to directly getting content in repoCommitFile.RawURL
					// TODO: see logs in "Panic bug scanning commits for CreateEvent" note
					log.Printf("Warning: commit %+v has empty patch data, but cannot fetch blob since repoCommitFile.SHA is nil.  Skipping..", repoCommitFile)
					continue
				}

				leaksForFile, err := guef.ScanBlob(username, repoName, *repoCommitFile.SHA, accessKeysToScan)
				if len(leaksForFile) > 0 {
					leaks = append(leaks, leaksForFile...)
				}
				if err != nil {
					// ScanBlob only returns temporary/recoverable errors, and so at this point might as well abort
					return leaks, err
				}

			}

		}

	}

	return leaks, nil

}

func (guef GoGithubUserEventFetcher) ScanDownstreamContent(ctx context.Context, userEvent *github.Event, accessKeysToScan []FetchedAwsAccessKey) (leaks []FetchedAwsAccessKey, err error) {

	payload, err := userEvent.ParsePayload()
	if err != nil {
		return nil, err
	}

	switch v := payload.(type) {
	case *github.PullRequestEvent:

		// TODO: what if this Pull Request has more commits than will be returned in content from patch url?

		content, err := guef.FetchUrlContent(ctx, v.PullRequest.GetPatchURL())
		if err != nil {
			return nil, err
		}
		log.Printf("Scanning %d bytes of content for event: %v url: %v", len(content), *userEvent.ID, v.PullRequest.GetPatchURL())

		return Scan(accessKeysToScan, content)

	case *github.CreateEvent:

		log.Printf("CreateEvent id: %v", *userEvent.ID)

		switch *v.RefType {
		case "tag":
			log.Printf("CreateEvent.  New tag: %v", *v.Ref)
			fallthrough
		case "branch":
			repoNameComponents := strings.Split(*userEvent.Repo.Name, "/")
			username := repoNameComponents[0]
			repoName := repoNameComponents[1]

			log.Printf("CreateEvent.  New branch/tag: %v in repo %v.  Scanning recent commits.", *v.Ref, *userEvent.Repo.Name)

			// If it's not running in the context of an integration test, then ignore test branches
			if !IntegrationTestsEnabled() && strings.Contains(*v.Ref, keynuker_go_common.KeyNukerIntegrationTestBranch) {
				// skip this since as an experiment
				log.Printf("Skipping CreateEvent %v since it's on %v testing branch", *userEvent.ID, keynuker_go_common.KeyNukerIntegrationTestBranch)
				return []FetchedAwsAccessKey{}, nil
			}

			// This will list the last 20 commits on the branch or tag and scan them
			// TODO: detect if there are more than 20 commits that haven't been scanned yet (currently no way to do that)
			// TODO: and if there are, trigger a deep scan on this repo, which will git clone the repo scan local content
			commitListOptions := &github.CommitsListOptions{
				SHA: *v.Ref,
				ListOptions: github.ListOptions{
					PerPage: 20,
					Page:    0,
				},
			}
			commits, _, err := guef.ApiClient.Repositories.ListCommits(
				ctx,
				username,
				repoName,
				commitListOptions,
			)
			if err != nil {
				// Ignore 404 not found errors, since the branch may no longer exist
				if strings.Contains(err.Error(), "404 Not Found") {
					log.Printf("Warning: Skipping branch/tag since it apparently no longer exists.  Err: %v Err Type: %T", err, err)
					return []FetchedAwsAccessKey{}, nil
				}
				return []FetchedAwsAccessKey{}, errors.Wrapf(err, "Error calling ListCommits on CreateEvent id: %v", *userEvent.ID)
			}

			return guef.ScanContentForCommits(
				ctx,
				username,
				repoName,
				ConvertRepositoryCommits(commits),
				accessKeysToScan,
			)

		case "repo":
			log.Printf("CreateEvent.  New repo.  Not scanning any commits.")
		default:
			log.Printf("Unknown CreateEvent reftype: %v", *v.RefType)
		}

	case *github.PushEvent:

		leaks = []FetchedAwsAccessKey{}

		// The github API only returns a maximum of 20 commits per push event
		maxCommitsPerPushEvent := 20

		// Ignore test branches.  These can always be ignored, even in integration tests, because in
		// the current integration tests it ends up merging to master and scanning the commits on master.
		if strings.Contains(*v.Ref, keynuker_go_common.KeyNukerIntegrationTestBranch) {
			// skip this since as an experiment
			log.Printf("Skipping push event %v since it's on %v testing branch", *v.PushID, keynuker_go_common.KeyNukerIntegrationTestBranch)
			return []FetchedAwsAccessKey{}, nil
		}

		// split "org/reponame" into separate strings (["org", "reponame"])
		repoNameComponents := strings.Split(*userEvent.Repo.Name, "/")
		username := repoNameComponents[0]
		repoName := repoNameComponents[1]

		leaksForCommits, err := guef.ScanContentForCommits(
			ctx,
			username,
			repoName,
			ConvertPushEventCommits(v.Commits),
			accessKeysToScan,
		)
		if err != nil {
			return leaksForCommits, errors.Wrapf(err, "Error calling ScanContentForCommits on PushEvent: %v", *v.PushID)
		}
		if len(leaksForCommits) > 0 {
			leaks = append(leaks, leaksForCommits...)
		}

		// If more than 20 commits for this PushEvent, fetch content for the remaining commits.
		// Example PushEvent w/ more than 20 commits: https://gist.github.com/tleyden/68d972b02b2b9306fa6e2eb26310b751
		if *v.Size > maxCommitsPerPushEvent {

			log.Printf("PushEvent %v has > 20 commits but this API call only returns 20.  Making separate API call.", *v.PushID)

			// Fetch the rest of the commits for this push event and append downstream content to buffer
			// TODO: this should be changed to ScanCommitsForPushEvent
			leaksForPushEvent, err := guef.ScanCommitsForPushEvent(ctx, userEvent, v, accessKeysToScan)
			if err != nil {
				return leaksForPushEvent, errors.Wrapf(err, "Error fetching additional commits for push event: %v", *v.PushID)
			}
			if len(leaksForPushEvent) > 0 {
				leaks = append(leaks, leaksForPushEvent...)
			}

		}
		return leaks, nil
		// return Scan(accessKeysToScan, buffer.Bytes())

	default:
		// TODO: add case statements for all events that should be scanned
		return Scan(accessKeysToScan, []byte(*userEvent.RawPayload))
	}

	return nil, nil
}

// Since PushEvents only contain 20 commits max, this fetches the remaining commits and writes the content to the
// writer passed in.  For example, pushEvent.Size might indicate that there were 100 commits in the push events,
// and so the remaining 80 commits will need to be scanned.
//
// Github API: https://developer.github.com/v3/repos/commits/
func (guef GoGithubUserEventFetcher) ScanCommitsForPushEvent(
	ctx context.Context, userEvent *github.Event, pushEvent *github.PushEvent, accessKeysToScan []FetchedAwsAccessKey) (leaks []FetchedAwsAccessKey, err error) {

	log.Printf("ScanCommitsForPushEvent: %v.  Number of total commits in push event: %d",
		*pushEvent.PushID, *pushEvent.Size)

	leaks = []FetchedAwsAccessKey{}

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
			return leaks, nil
		}

		if resultPage > maxPages {
			notScanned := numCommitsToScan - numCommitsScanned
			log.Printf("WARNING: not scanning %d commits, due to current limitations.  See https://github.com/tleyden/keynuker/issues/24", notScanned)
			return leaks, nil
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
			return leaks, errors.Wrapf(err, "Error calling ApiClient.Repositories.ListCommits on PushEvent: %v", *pushEvent.PushID)
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
				return leaks, errors.Wrapf(err, "Error getting additional content for commit: %v url: %v", *additionalCommit.SHA, additionalCommit.GetURL())
			}

			// Loop over the files in the commit and append the content to the writer
			for _, repoCommitFile := range repoCommit.Files {

				leaksForCommit, err := Scan(accessKeysToScan, []byte(repoCommitFile.GetPatch()))
				if err != nil {
					log.Printf("Warning: Error scanning commit: %v.  Error: %v.  Skipping commit", repoCommitFile, err)
					continue
				}
				if len(leaksForCommit) > 0 {
					leaks = append(leaks, leaksForCommit...)
				}

				if repoCommitFile.Patch == nil {

					// This means its binary data or larger than 1 MB, call separate API to fetch
					log.Printf("Warning: additional commit %+v has empty patch data.  Either binary data, or greater than 1MB", repoCommitFile)

					if repoCommitFile.SHA == nil {
						// TODO: move this code into a function called GetContentForCommit(commit) which will use GetBlob() if there
						// TODO: is a SHA, otherwise it will fallback to directly getting content in repoCommitFile.RawURL
						// TODO: see logs in "Panic bug scanning commits for CreateEvent" note
						log.Printf("Warning: additional commit %+v has empty patch data, but repoCommitFile.SHA is nil.  Skipping", repoCommitFile)
						continue
					}

					leaksForFile, err := guef.ScanBlob(owner, repoName, *repoCommitFile.SHA, accessKeysToScan)
					if len(leaksForFile) > 0 {
						leaks = append(leaks, leaksForFile...)
					}
					if err != nil {
						// ScanBlob only returns temporary/recoverable errors, and so at this point might as well abort
						return leaks, err
					}

				}

			}

			numCommitsScanned += 1

			if numCommitsScanned >= numCommitsToScan {
				// done
				return leaks, nil
			}

		}

		resultPage = resp.NextPage

	}

	return leaks, nil
}

func (guef GoGithubUserEventFetcher) ScanBlob(owner string, repo string, sha string, accessKeysToScan []FetchedAwsAccessKey) (leaks []FetchedAwsAccessKey, err error) {

	ctx := context.Background()

	leaks = []FetchedAwsAccessKey{}

	blob, _, err := guef.ApiClient.Git.GetBlob(
		ctx,
		owner,
		repo,
		sha,
	)
	if err != nil {
		if keynuker_go_common.IsTemporaryGithubError(err) {
			// Abort now since this will prevent checkpoint from getting pushed, and will be retried later
			return leaks, errors.Wrapf(err, "Temporary error getting content for commit file via blob api.  Owner: %v Repo: %v Sha: %v", owner, repo, sha)
		} else {
			log.Printf("Permanent error getting content for commit file via blob api. Skipping commit.  Owner: %v Repo: %v Sha: %v", owner, repo, sha)
			return leaks, nil
		}
	}

	if *blob.Encoding != "base64" {
		log.Printf("Warning: skipping blob from commit with owner: %v Repo: %v Sha: %v, since unexpected encoding (%v)", owner, repo, sha, *blob.Encoding)
		return leaks, nil
	}
	if *blob.Size > keynuker_go_common.MaxSizeBytesBlobContent {
		log.Printf("Warning: skipping blob from commit with owner: %v Repo: %v Sha: %v, since size > max size (%d).", owner, repo, sha, keynuker_go_common.MaxSizeBytesBlobContent)
		return leaks, nil
	}

	decodedBlobContent, err := base64.StdEncoding.DecodeString(blob.GetContent())
	if err != nil {
		log.Printf("Warning: error decoding base64 for blob commit with owner: %v Repo: %v Sha: %v via blob api.  Err: %v.  Skipping commit.", owner, repo, sha, err)
		return leaks, nil
	}

	log.Printf("Scanning %d bytes of content for blob commit with owner: %v Repo: %v Sha: %v", len(decodedBlobContent), owner, repo, sha)

	leaksForFile, err := Scan(accessKeysToScan, decodedBlobContent)
	if err != nil {
		log.Printf("WARNING: error scanning content for blob commit with owner: %v Repo: %v Sha: %v Err: %v  Skipping commit.", owner, repo, sha, err)
		return leaks, nil
	}
	if len(leaksForFile) > 0 {
		leaks = append(leaks, leaksForFile...)
	}

	return leaks, nil

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

	// Read from the response, but limit the number of bytes read to 10MB to avoid blowing up the memory
	// for extra large commits
	resp_body, err := keynuker_go_common.ReadLimited(resp.Body, keynuker_go_common.MaxSizeBytesBlobContent)
	if err != nil {
		return nil, err
	}

	if len(resp_body) == keynuker_go_common.MaxSizeBytesBlobContent {
		// TODO: This gives a false warning if the content was _exactly_ keynuker_go_common.MaxSizeBytesBlobContent bytes.
		// TODO: The ReadLimited() function should return a boolean indicating if there was unread content
		log.Printf("WARNING: only %d bytes of content were scanned for url: %v.  Some content was not scanned.", keynuker_go_common.MaxSizeBytesBlobContent, url)
	}

	return resp_body, nil

}
