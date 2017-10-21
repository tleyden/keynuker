package keynuker

import (
	"encoding/json"
	"log"
	"testing"

	"os"

	"github.com/stretchr/testify/assert"
	"github.com/tleyden/keynuker/keynuker-go-common"
)

func TestFetchAwsKeys(t *testing.T) {

	SkipIfIntegrationsTestsNotEnabled(t)

	targetAwsAccounts, err := GetTargetAwsAccountsFromEnv()
	if err != nil {
		t.Skip("Error getting target aws accounts from env: %v", err)
	}

	accessKeyId, ok := os.LookupEnv(keynuker_go_common.EnvVarKeyNukerAwsAccessKeyId)
	if !ok {
		t.Fatalf("You must define environment variable %s", keynuker_go_common.EnvVarKeyNukerAwsAccessKeyId)
	}

	secretAccessKey, ok := os.LookupEnv(keynuker_go_common.EnvVarKeyNukerAwsSecretAccessKey)
	if !ok {
		t.Fatalf("You must define environment variable %s", keynuker_go_common.EnvVarKeyNukerAwsAccessKeyId)
	}

	params := ParamsFetchAwsKeys{
		KeyNukerOrg:       "default",
		TargetAwsAccounts: targetAwsAccounts,
		InitiatingAwsAccountAssumeRole: AwsCredentials{
			AwsAccessKeyId:     accessKeyId,
			AwsSecretAccessKey: secretAccessKey,
		},
	}

	doc, err := FetchAwsKeys(params)

	assert.NoError(t, err, "Unexpected Error")

	marshalled, err := json.MarshalIndent(doc, "", "    ")

	log.Printf("FetchedAwsKeys: %v", string(marshalled))

}
