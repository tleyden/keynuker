package main

import (
	"encoding/json"

	"github.com/tleyden/keynuker"
	"github.com/tleyden/keynuker/keynuker-go-common"
)

func main() {

	keynuker_go_common.InvokeActionStdIo(OpenWhiskCallback)

}

func OpenWhiskCallback(value json.RawMessage) (interface{}, error) {

	activationsStatus := keynuker.OpenWhiskRecentActivationsStatus()

	return activationsStatus, nil

}
