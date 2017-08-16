package main

import (
	"encoding/json"
	"log"

	"bufio"
	"fmt"
	"net/http"
	"os"

	"github.com/apache/incubator-openwhisk-client-go/whisk"
	"github.com/tleyden/keynuker/keynuker-go-common"
	"strings"
	"net/url"
)

func main() {

	keynuker_go_common.InvokeActionStdIo(OpenWhiskCallback)
}

func OpenWhiskCallback(value json.RawMessage) (interface{}, error) {

	log.Printf("monitor-activations called with %v", string(value))

	keyNukerStatus := KeyNukerStatus()

	marshalled, _ := json.MarshalIndent(keyNukerStatus, "", "    ")
	log.Printf("%v\n", marshalled)

	return keyNukerStatus, nil

}

// Connect to OpenWhisk API and scan the list of recent activations and look for any failures.
// If any failures found, return {"status": "failure"}.  Otherwise return {"status": "success"}.
// The idea is that this would be served up by a web action that a monitoring tool could poll
// and send alerts if any failures occurred.
func KeyNukerStatus() (keynukerStatus map[string]interface{}) {

	keynukerStatus = map[string]interface{}{}
	keynukerStatus["status"] = "failure"

	whiskConfig, err := WhiskConfigFromEnvironment()
	if err != nil {
		msg := fmt.Sprintf("Error getting whisk config from environment: %v", err)
		log.Printf(msg)
		keynukerStatus["error"] = msg
		return keynukerStatus
	}

	whiskConfig.Debug = true

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
func ScanActivationsForFailures(whiskConfig *whisk.Config) (failedActivations []whisk.Activation, err error) {

	/*

	The json returned for a failed activation will look like this:

	"response": {
        "status": "action developer error",
        "statusCode": 0,
        "success": false,
        "result": {
            "error": "The action did not return a dictionary."
        }
    },

    Whereas the json for a success activation will look more like this:

    "response": {
        "status": "success",
        "statusCode": 0,
        "success": true,
        "result": {
            "payload": ""
        }
    },

	 */

	client, _ := whisk.NewClient(http.DefaultClient, whiskConfig)

	/*

	type ActivationListOptions struct {
    Name  string `url:"name,omitempty"`
    Limit int    `url:"limit"`
    Skip  int    `url:"skip"`
    Since int64  `url:"since,omitempty"`
    Upto  int64  `url:"upto,omitempty"`
    Docs  bool   `url:"docs,omitempty"`
}
	 */

	listActivationsOptions := &whisk.ActivationListOptions{
		Docs: true,  // Need to include this to get the activation doc body
	}

	activations, _, err := client.Activations.List(listActivationsOptions)
	if err != nil {
		return failedActivations, err
	}
	for _, activation := range activations {
		if activation.Response.Success == false {
			failedActivations = append(failedActivations, activation)
		}
	}
	return failedActivations, nil
}


func WhiskPropsMapFromEnvironment() (map[string]string, error) {

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

func WhiskConfigFromEnvironment() (*whisk.Config, error) {

	config := &whisk.Config{}

	whiskPropsMap, err := WhiskPropsMapFromEnvironment()
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
			apiHost.Scheme = "http"  // TODO: what should this be?
			config.BaseURL = apiHost
		}

	}

	return config, nil

}
