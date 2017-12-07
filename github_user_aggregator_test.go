package keynuker

import (
	"context"
	"log"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tleyden/keynuker/keynuker-go-common"
)

// Not much of a unit test, just allows running ghUserAggregator.ListMembers() by hand in isolation
func TestGithubUserAggregator(t *testing.T) {

	SkipIfIntegrationsTestsNotEnabled(t)

	accessToken, ok := os.LookupEnv(keynuker_go_common.EnvVarKeyNukerTestGithubAccessToken)
	if !ok {
		t.Skipf("You must define environment variable keynuker_test_gh_access_token to run this test")
	}


	ghOrgs := []string{
		"couchbase",
		"couchbaselabs",
		"couchbasedeps",
		"couchbase-partners",
	}

	largestOrg := 0
	numUsersInEachOrgIndividually := 0
	for _, ghOrg := range ghOrgs {
		ghUserAggregator := NewGithubUserAggregator([]string{ghOrg}, accessToken)
		users, err := ghUserAggregator.ListMembers(context.Background())
		assert.True(t, err == nil)
		numUsersInEachOrgIndividually += len(users)
		if len(users) > largestOrg {
			largestOrg = len(users)
		}
	}


	ghUserAggregator := NewGithubUserAggregator(ghOrgs, accessToken)
	users, err := ghUserAggregator.ListMembers(context.Background())

	assert.True(t, err == nil)

	log.Printf("Members of all orgs %v: %v", ghOrgs, users)
	log.Printf("# of Members: %v.  numUsersInEachOrgIndividually: %v", len(users), numUsersInEachOrgIndividually)

	assert.True(t, numUsersInEachOrgIndividually >= len(users))  // Can be lots of overlap
	assert.True(t, len(users) > largestOrg)

}
