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

	awsAccessKeyId, ok := os.LookupEnv(keynuker_go_common.EnvVarKeyNukerTestAwsAccessKeyId)
	if !ok {
		t.Skip("You must define environment variable %s to run this test", keynuker_go_common.EnvVarKeyNukerTestAwsAccessKeyId)
	}

	awsSecretAccessKey, ok := os.LookupEnv(keynuker_go_common.EnvVarKeyNukerTestAwsSecretAccessKey)
	if !ok {
		t.Skip("You must define environment variable %s to run this test", keynuker_go_common.EnvVarKeyNukerTestAwsSecretAccessKey)
	}

	targetAwsAccounts := []TargetAwsAccount{
		{
			AwsAccessKeyId:     awsAccessKeyId,
			AwsSecretAccessKey: awsSecretAccessKey,
		},
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
