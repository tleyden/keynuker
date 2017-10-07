// Licensed to the Apache Software Foundation (ASF) under one or more contributor license agreements;
// and to You under the Apache License, Version 2.0.  See LICENSE in project root for full license + copyright.

package keynuker_go_common

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"log"
)

// Same function signature as github.com/jthomas/ow callback, but that's not exposed so redefine it here
type OpenWhiskCallback func(json.RawMessage) (interface{}, error)

func GenerateDocId(docIdPrefix, keyNukerOrg string) string {
	return fmt.Sprintf("%s_%s", docIdPrefix, keyNukerOrg)
}

func InvokeActionStdIo(callback OpenWhiskCallback) {

	// read everything available on stdin
	bytes, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		panic(fmt.Sprintf("Error reading stdin.  Err: %v", err))
	}

	docWrapper, err := callback(bytes)
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

type Logger interface {
	Printf(format string, v ...interface{})
}

type BoundedLogger struct {
	numInvocations int
	maxInvocations int
}

func (b *BoundedLogger) Printf(format string, v ...interface{}) {
	if b.numInvocations < b.maxInvocations {
		log.Printf(format, v)
		b.numInvocations += 1
	}
}

func CreateBoundedLogger(maxInvocations int) Logger {
	return &BoundedLogger{
		numInvocations: 0,
		maxInvocations: maxInvocations,
	}

}