package main

import (
	"encoding/json"
	"log"

	"github.com/tleyden/keynuker"
	"github.com/tleyden/keynuker/keynuker-go-common"
)

func main() {
	keynuker_go_common.LogMemoryUsageLoop()
	keynuker_go_common.RegistorOrInvokeActionStdIo(OpenWhiskCallback)
}

func OpenWhiskCallback(value json.RawMessage) (interface{}, error) {

	var params keynuker.ParamsMonitorActivations

	err := json.Unmarshal(value, &params)
	if err != nil {
		return nil, err
	}

	input := keynuker.RecentActivationsReportInput{
		MaxActivationsToScan: 200,
	}
	output, err := keynuker.OpenWhiskRecentActivationsReport(input)

	if err != nil {
		return nil, err
	}

	deliveryId, err := keynuker.SendReportNotifications(params, output)
	if err != nil {
		return nil, err
	}

	log.Printf("Report delivery id: %v", deliveryId)

	return output, nil

}
