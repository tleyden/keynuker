package keynuker

import (
	"encoding/json"
	"log"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tleyden/keynuker/keynuker-go-common"
)

func TestFetchAwsKeys(t *testing.T) {

	SkipIfIntegrationsTestsNotEnabled(t)

	targetAwsAccounts, err := GetTargetAwsAccountsFromEnv()
	if err != nil {
		t.Skip("Error getting target aws accounts from env: %v", err)
	}

	params := ParamsFetchAwsKeys{
		KeyNukerOrg:       "default",
		TargetAwsAccounts: targetAwsAccounts,
	}

	doc, err := FetchAwsKeys(params)

	assert.NoError(t, err, "Unexpected Error")

	marshalled, err := json.MarshalIndent(doc, "", "    ")

	log.Printf("FetchedAwsKeys: %v", string(marshalled))

}
