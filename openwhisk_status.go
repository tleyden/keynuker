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
	"gopkg.in/mailgun/mailgun-go.v1"
	"github.com/pkg/errors"
	"regexp"
	"strconv"
)

type ParamsMonitorActivations struct {

	// MailerParams
	MailerParams

	// This is the name of the KeyNuker "org/tenant".  Defaults to "default", but allows to be extended multi-tenant.
	KeyNukerOrg string

	// The FROM address that will be used for any notifications
	EmailFromAddress string `json:"email_from_address"`

	// Optionally specify the Keynuker admin email to be CC'd about any leaked/nuked keys
	KeynukerAdminEmailCCAddress string `json:"admin_email_cc_address"`
}

func SendMonitorNotifications(params ParamsMonitorActivations, activationStatus map[string]interface{}) (deliveryId string, err error) {

	mailer := mailgun.NewMailgun(
		params.MailerParams.Domain,
		params.MailerParams.ApiKey,
		params.MailerParams.PublicApiKey,
	)

	messageBody := fmt.Sprintf(
		"Failed activations: %+v",
		activationStatus,
	)

	message := mailgun.NewMessage(
		params.EmailFromAddress,
		"KeyNuker Monitoring: failed activations 💥",
		messageBody,
		params.KeynukerAdminEmailCCAddress,
	)

	mes, id, err := mailer.Send(message)
	if err != nil {
		return "", fmt.Errorf("Error sending message: %v.  Mes: %v", err, mes)
	}

	return id, nil

}


func SendReportNotifications(params ParamsMonitorActivations, report RecentActivationsReportOutput) (deliveryId string, err error) {

	mailer := mailgun.NewMailgun(
		params.MailerParams.Domain,
		params.MailerParams.ApiKey,
		params.MailerParams.PublicApiKey,
	)

	var messageBody string
	if len(report.FailedActivations) > 0 {
		messageBody += fmt.Sprintf("Failed activations: %+v\n", report.FailedActivations)
	}
	messageBody += fmt.Sprintf("Num bytes scanned: %d", report.TotalNumBytesScanned)

	message := mailgun.NewMessage(
		params.EmailFromAddress,
		"KeyNuker Activity Report 🔐",
		messageBody,
		params.KeynukerAdminEmailCCAddress,
	)

	mes, id, err := mailer.Send(message)
	if err != nil {
		return "", fmt.Errorf("Error sending message: %v.  Mes: %v", err, mes)
	}

	return id, nil

}


type RecentActivationsReportInput struct {
	MaxActivationsToScan int

}

type RecentActivationsReportOutput struct {
	FailedActivations []whisk.Activation
	TotalNumBytesScanned int64
}

// A more generalized version of OpenWhiskRecentActivationsStatus
// TODO #1: Moving to structured logging (logrus?) will make this a lot more tenable.  Either that or json stats.
func OpenWhiskRecentActivationsReport(input RecentActivationsReportInput) (output RecentActivationsReportOutput, err error) {

	whiskConfig, err := WhiskConfigFromEnvironment()
	if err != nil {
		return output, errors.Wrapf(err,"Error getting whisk config from environment")
	}

	client, err := whisk.NewClient(http.DefaultClient, whiskConfig)
	if err != nil {
		return output, errors.Wrapf(err, "Error creating whisk.NewClient")
	}

	// TODO: is this needed?
	// output.FailedActivations = []whisk.Activation{}

	// This must limited to a small number, otherwise it will exceed memory limits and get killed abruptly
	pageSize := 25

	// Keep track
	skipOffset := 0

	for {

		// Check to see if we've already scanned far enough back
		numActivationsScanned := skipOffset
		if numActivationsScanned >= input.MaxActivationsToScan {
			// return what we have so far (should be no failures)
			return output, nil
		}

		listActivationsOptions := &whisk.ActivationListOptions{
			Docs:  true, // Need to include this to get the activation doc body, which ends up using lots of memory
			Limit: pageSize,
			Skip:  skipOffset,
		}

		// Make REST call to OpenWhisk API to load list of activations
		activations, _, err := client.Activations.List(listActivationsOptions)

		if err != nil {
			return output, errors.Wrapf(err, "client.Activations.List with options %+v returned error", listActivationsOptions)
		}

		if len(activations) == 0 {
			// Looks like we hit the end of list of total avaialable activations
			return output, nil
		}

		// Loop over activations and look for failures and total up bytes scanned by scanning logs
		for _, activation := range activations {
			if activation.Name == "monitor-activations" {
				continue
			}
			if activation.Response.Success == false {
				output.FailedActivations = append(output.FailedActivations, activation)
			}
			bytesScanned, err := CalculateBytesScanned(activation.Logs)
			if err != nil {
				return output, errors.Wrapf(err, "Error calculating bytes scanned from activation %v logs", activation.ActivationID)
			}
			output.TotalNumBytesScanned += bytesScanned
		}

	}

	return output, nil
}

// Look for log messages with form:
//    Scanning 1833 bytes
// and extract the number of bytes, and then add them up
func CalculateBytesScanned(logs []string) (int64, error) {

	numBytesAccumulated := int64(0)

	r, err := regexp.Compile(`Scanning (\d*) bytes`)
	if err != nil {
		return 0, errors.Wrapf(err, "Error compiling regex")
	}

	for _, logLine := range logs {
		result := r.FindStringSubmatch(logLine)
		if len(result) > 0 {
			numBytesAsStr := result[1]
			numBytes, err := strconv.Atoi(numBytesAsStr)
			if err != nil {
				return 0, errors.Wrapf(err, "Error converting %s to a number", numBytesAsStr)
			}
			numBytesAccumulated += int64(numBytes)
		}
	}


	return numBytesAccumulated, nil
}

// Connect to OpenWhisk API and scan the list of recent activations and look for any failures.
// If any failures found, return {"status": "failure"}.  Otherwise return {"status": "success"}.
// The idea is that this would be served up by a web action that a monitoring tool could poll
// and send alerts if any failures occurred.
func OpenWhiskRecentActivationsStatus(maxActivationsToScan int) (keynukerStatus map[string]interface{}) {

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

	failedActivations, err := ScanActivationsForFailures(whiskConfig, maxActivationsToScan)
	log.Printf("ScanActivationsForFailures returned %d failedActivations", len(failedActivations))

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
	} else {
		failedActivationIds := []string{}
		for _, failedActivation := range failedActivations {
			failedActivationIds = append(failedActivationIds, failedActivation.ActivationID)
		}
		keynukerStatus["failedActivationIds"] = failedActivationIds
	}

	log.Printf("keynukerStatus: %+v", keynukerStatus)

	return keynukerStatus

}

// Loop over all activations and return the ones that have a whisk.Result with Success == false.
// Stop scanning after maxActivationsToScan activations have been scanned
func ScanActivationsForFailures(whiskConfig *whisk.Config, maxActivationsToScan int) (failedActivations []whisk.Activation, err error) {

	client, err := whisk.NewClient(http.DefaultClient, whiskConfig)
	if err != nil {
		return failedActivations, err
	}

	failedActivations = []whisk.Activation{}

	// This must limited to a small number, otherwise it will exceed memory limits and get killed abruptly
	pageSize := 25

	// Keep track
	skipOffset := 0

	for {

		// Check to see if we've already scanned far enough back
		numActivationsScanned := skipOffset
		if numActivationsScanned >= maxActivationsToScan {
			// return what we have so far (should be no failures)
			log.Printf("numActivationsScanned (%d) >= maxActivationsToScan (%d).  return failedActivations: %v", numActivationsScanned, maxActivationsToScan, failedActivations)

			return failedActivations, nil
		}

		listActivationsOptions := &whisk.ActivationListOptions{
			Docs:  true, // Need to include this to get the activation doc body, which ends up using lots of memory
			Limit: pageSize,
			Skip:  skipOffset,
		}

		log.Printf("List Activations with: %+v", listActivationsOptions)
		// Make REST call to OpenWhisk API to load list of activations
		activations, _, err := client.Activations.List(listActivationsOptions)
		log.Printf("List Activations returned %d activations", len(activations))

		if err != nil {
			log.Printf("List Activations returned err: %v", err)
			return failedActivations, err
		}

		if len(activations) == 0 {
			// Looks like we hit the end of list of total avaialable activations
			return failedActivations, nil
		}

		// Loop over activations and look for failures
		for _, activation := range activations {
			if activation.Name == "monitor-activations" {
				log.Printf("Ignoring monitor-activations activation: %v", activation.ActivationID)
				continue
			}
			if activation.Response.Success == false {
				log.Printf("Detected failed activation: %v", activation.ActivationID)
				failedActivations = append(failedActivations, activation)
			}
		}

		// If we found any failures, just return early
		if len(failedActivations) > 0 {
			log.Printf("len(failedActivations) > 0 (=%d).  Returning: %v", len(failedActivations), failedActivations)

			return failedActivations, nil
		}

		// Go to the next page of data
		skipOffset += pageSize

	}

	// Should never get here
	if len(failedActivations) > 0 {
		log.Printf("len(failedActivations) > 0 (=%d).  Returning: %v", len(failedActivations), failedActivations)
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
			config.Host = val
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

	config.Host = owApiHost
	
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
