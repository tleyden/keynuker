package main

import (
	"encoding/json"
	"log"

	"github.com/tleyden/keynuker"
	"github.com/tleyden/keynuker/keynuker-go-common"
)

func main() {

	keynuker_go_common.InvokeActionStdIo(OpenWhiskCallback)

}

func OpenWhiskCallback(value json.RawMessage) (interface{}, error) {

	log.Printf("monitor-activations called with: %v", string(value))

	activationsStatus := keynuker.OpenWhiskRecentActivationsStatus()

	log.Printf("activationStatus raw: %+v\n", activationsStatus)

	marshalled, _ := json.MarshalIndent(activationsStatus, "", "    ")

	log.Printf("activationStatus: %v\n", string(marshalled))

	return activationsStatus, nil

}
