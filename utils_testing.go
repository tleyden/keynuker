package keynuker

import (
	"testing"
	"os"
	"github.com/tleyden/keynuker/keynuker-go-common"
	"fmt"
	"strings"
)

func SkipIfIntegrationsTestsNotEnabled(t *testing.T) {

	errMsg := fmt.Sprintf("You must define environment variable %s and set to true to enable integration tests", keynuker_go_common.EnvVarKeyNukerTestIntegrationTestsEnabled)

	enabled := IntegrationTestsEnabled()

	if !enabled {
		t.Skip(errMsg)
	}

}

func IntegrationTestsEnabled() bool {

	testIntegrationTestsEnabled, ok := os.LookupEnv(keynuker_go_common.EnvVarKeyNukerTestIntegrationTestsEnabled)
	if !ok {
		return false
	}

	if strings.ToLower(testIntegrationTestsEnabled) != "true" {
		return false
	}

	return true
}
