package main

import (
	"encoding/json"

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

	activationsStatus := keynuker.OpenWhiskRecentActivationsStatus()
	if activationsStatus["status"] == "failure" {
		deliveryId, err := keynuker.SendMonitorNotifications(params, activationsStatus)
		if err != nil {
			return nil, err
		}
		activationsStatus["notificationDeliveryId"] = deliveryId
	}

	return activationsStatus, nil

}
