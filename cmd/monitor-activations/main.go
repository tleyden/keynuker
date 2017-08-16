package main

import (
	"encoding/json"
	"log"

	"github.com/tleyden/keynuker/keynuker-go-common"
	"github.com/tleyden/keynuker"
)

func main() {

	keynuker_go_common.InvokeActionStdIo(OpenWhiskCallback)
}

func OpenWhiskCallback(value json.RawMessage) (interface{}, error) {

	log.Printf("monitor-activations called with %v", string(value))

	activationsStatus := keynuker.OpenWhiskRecentActivationsStatus()

	marshalled, _ := json.MarshalIndent(activationsStatus, "", "    ")
	log.Printf("%v\n", marshalled)

	return activationsStatus, nil

}
