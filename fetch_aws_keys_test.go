package keynuker

import (
	"encoding/json"
	"log"
	"testing"

	"os"

	"github.com/stretchr/testify/assert"
	"github.com/tleyden/keynuker/keynuker-go-common"
)

func TestFetchAwsKeysViaSTSAssumeRole(t *testing.T) {

	SkipIfIntegrationsTestsNotEnabled(t)

	targetAwsAccounts, err := GetTargetAwsAccountsFromEnv()
	if err != nil {
		t.Skip("Error getting target aws accounts from env: %v", err)
	}

	initiatingAwsAcctCredsRaw, ok := os.LookupEnv(keynuker_go_common.EnvVarKeyNukerInitiatingAwsAccountCreds)
	if !ok {
		t.Fatalf("You must define environment variable %s", keynuker_go_common.EnvVarKeyNukerInitiatingAwsAccountCreds)
	}

	initiatingAwsAcctCreds := AwsCredentials{}
	if err := json.Unmarshal([]byte(initiatingAwsAcctCredsRaw), &initiatingAwsAcctCreds); err != nil {
		t.Fatalf("Error unmarshalling creds.  Err: %v", err)
	}

	params := ParamsFetchAwsKeys{
		KeyNukerOrg:       "default",
		TargetAwsAccounts: targetAwsAccounts,
		InitiatingAwsAccountAssumeRole: initiatingAwsAcctCreds,
	}

	doc, err := FetchAwsKeys(params)

	assert.NoError(t, err, "Unexpected Error")

	marshalled, err := json.MarshalIndent(doc, "", "    ")

	log.Printf("FetchedAwsKeys: %v", string(marshalled))

}
