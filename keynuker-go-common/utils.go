// Licensed to the Apache Software Foundation (ASF) under one or more contributor license agreements;
// and to You under the Apache License, Version 2.0.  See LICENSE in project root for full license + copyright.

package keynuker_go_common

import (
	"encoding/json"
	"fmt"
	"os"
)

// Same function signature as github.com/jthomas/ow callback, but that's not exposed so redefine it here
type OpenWhiskCallback func(json.RawMessage) (interface{}, error)

func GenerateDocId(docIdPrefix, keyNukerOrg string) string {
	return fmt.Sprintf("%s_%s", docIdPrefix, keyNukerOrg)
}

func StdioFlagPresent() bool {

	// arg0 - program name
	// arg1 - input json (in the case we are using stdio interface)
	// arg2 - "true" string to further indicate we are using stdio interface
	if len(os.Args) >= 3 && os.Args[2] == "true" {
		return true
	}
	return false

}

func InvokeActionStdIo(callback OpenWhiskCallback) {

	// native actions receive one argument, the JSON object as a string
	arg := os.Args[1]

	docWrapper, err := callback([]byte(arg))
	if err != nil {
		panic(fmt.Sprintf("Error calling callback.  Err: %v", err))
	}

	outputBytes, err := json.Marshal(docWrapper)
	if err != nil {
		panic(fmt.Sprintf("Error marshalling outputBytes: %v", err))
	}

	// Write result doc to stdout
	fmt.Printf("%s", string(outputBytes))

}
