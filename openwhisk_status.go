package keynuker

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"

	"github.com/apache/incubator-openwhisk-client-go/whisk"
)

// Connect to OpenWhisk API and scan the list of recent activations and look for any failures.
// If any failures found, return {"status": "failure"}.  Otherwise return {"status": "success"}.
// The idea is that this would be served up by a web action that a monitoring tool could poll
// and send alerts if any failures occurred.
func OpenWhiskRecentActivationsStatus() (keynukerStatus map[string]interface{}) {

	keynukerStatus = map[string]interface{}{}
	keynukerStatus["status"] = "failure"

	whiskConfig, err := WhiskConfigFromEnvironment()
	if err != nil {
		msg := fmt.Sprintf("Error getting whisk config from environment: %v", err)
		log.Printf(msg)
		keynukerStatus["error"] = msg
		return keynukerStatus
	}

	// whiskConfig.Debug = true

	failedActivations, err := ScanActivationsForFailures(whiskConfig)

	if err != nil {
		msg := fmt.Sprintf("Error scanning activations for failures: %v", err)
		log.Printf(msg)
		keynukerStatus["error"] = msg
		// Don't return an actual error since this is a monitoring tool and it should always return a result
		// so the upstream web action returns the JSON response to the caller
		return keynukerStatus
	}

	if len(failedActivations) == 0 {
		keynukerStatus["status"] = "success"
	}

	return keynukerStatus

}

// Loop over all activations and return the ones that have a whisk.Result with Success == false
// TODO: To improve this, it should take a list of activation types to ignore (like monitor-activation where
// TODO: it sees it's own echoes of monitoring!) and then keeps pulling activations until it sees X non-ignored
// TODO: activations.  
func ScanActivationsForFailures(whiskConfig *whisk.Config) (failedActivations []whisk.Activation, err error) {

	client, err := whisk.NewClient(http.DefaultClient, whiskConfig)
	if err != nil {
		return failedActivations, err
	}

	listActivationsOptions := &whisk.ActivationListOptions{
		Docs:  true, // Need to include this to get the activation doc body, which ends up using lots of memory
		Limit: 20,   // This must limited to a small number, otherwise it will exceed memory limits and get killed abruptly
	}

	activations, _, err := client.Activations.List(listActivationsOptions)
	if err != nil {
		return failedActivations, err
	}

	for _, activation := range activations {
		if activation.Response.Success == false {
			log.Printf("Detected failed activation: %v", activation.ActivationID)
			failedActivations = append(failedActivations, activation)
		}
	}

	return failedActivations, nil
}

func WhiskConfigFromEnvironment() (config *whisk.Config, err error) {

	// First try to load from env variables and return that (eg, __OW_API_HOST).  This is what will run when running
	// on the BlueMix cloud
	config, err = WhiskConfigFromOwEnvVars()
	if err != nil {
		return nil, err
	}
	if config != nil {
		return config, nil
	}

	// Otherwise try to load config based on the contents of the WSK_CONFIG_FILE

	config = &whisk.Config{}

	whiskPropsMap, err := WhiskPropsMapFromWskConfigFile()
	if err != nil {
		return nil, err
	}

	for key, val := range whiskPropsMap {
		switch strings.ToUpper(key) {
		case "AUTH":
			config.AuthToken = val
		case "APIHOST":

			// Add "api" to workaround https://github.com/apache/incubator-openwhisk-client-go/issues/25
			apiUrl := fmt.Sprintf("http://%v/api", val)
			apiHost, err := url.Parse(apiUrl)
			if err != nil {
				return nil, fmt.Errorf("Unable to parse url (%v).  Error: %v", val, err)
			}
			apiHost.Scheme = "http" // TODO: what should this be?
			config.BaseURL = apiHost
		}

	}

	return config, nil

}

// Given a base hostname like "openwhisk.ng.bluemix.net" or "https://openwhisk.ng.bluemix.net:443", return a URL
// that includes a trailing "/api" in the path.
// The trailing /api is needed due to https://github.com/apache/incubator-openwhisk-client-go/issues/25
func CreateApiHostBaseUrl(hostname string) (baseUrl *url.URL, err error) {

	hostnameWithPath := fmt.Sprintf("%v/api", hostname)
	baseUrl, err = url.Parse(hostnameWithPath)
	if err != nil {
		return nil, fmt.Errorf("Unable to parse url (%v).  Error: %v", hostnameWithPath, err)
	}

	return baseUrl, nil
}

func WhiskConfigFromOwEnvVars() (config *whisk.Config, err error) {

	config = &whisk.Config{}

	owApiHost := os.Getenv("__OW_API_HOST")
	owApiKey := os.Getenv("__OW_API_KEY")

	// None of the env vars are set, return nil
	if owApiHost == "" || owApiKey == "" {
		return nil, nil
	}

	baseUrl, err := CreateApiHostBaseUrl(owApiHost)
	if err != nil {
		return nil, err
	}
	config.BaseURL = baseUrl

	config.AuthToken = owApiKey

	return config, nil

}

func WhiskPropsMapFromWskConfigFile() (map[string]string, error) {

	whiskPropsMap := map[string]string{}

	wskConfigFilePath := os.Getenv("WSK_CONFIG_FILE")
	if wskConfigFilePath == "" {
		return nil, fmt.Errorf("You need to set WSK_CONFIG_FILE to specify where to find .wskprops")
	}

	wskConfigFile, err := os.Open(wskConfigFilePath)
	if err != nil {
		return nil, fmt.Errorf("Error opening file: %v.  Err: %v", wskConfigFilePath, err)
	}
	defer wskConfigFile.Close()

	scanner := bufio.NewScanner(wskConfigFile)
	for scanner.Scan() {
		log.Println(scanner.Text())
		fields := strings.Split(scanner.Text(), "=")
		key := fields[0]
		val := fields[1]
		whiskPropsMap[key] = val

	}

	// check for errors
	if err = scanner.Err(); err != nil {
		return nil, fmt.Errorf("Error reading lines from file: %v. Err: %v", wskConfigFilePath, err)
	}

	return whiskPropsMap, nil

}
