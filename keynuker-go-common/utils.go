// Licensed to the Apache Software Foundation (ASF) under one or more contributor license agreements;
// and to You under the Apache License, Version 2.0.  See LICENSE in project root for full license + copyright.

package keynuker_go_common

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/tleyden/ow"
	"runtime"
	"time"
)

var UseDockerSkeleton bool

func GenerateDocId(docIdPrefix, keyNukerOrg string) string {
	return fmt.Sprintf("%s_%s", docIdPrefix, keyNukerOrg)
}

func RegistorOrInvokeActionStdIo(callback ow.OpenWhiskCallback) {

	log.Printf("UseDockerSkeleton: %v", UseDockerSkeleton)

	if UseDockerSkeleton {
		InvokeActionStdIo(WrapCallbackWithLogSentinel("ActionProxy", callback))
	} else {
		ow.RegisterAction(WrapCallbackWithLogSentinel("CustomDocker", callback))
	}
}

func LogMemoryUsage() {
	go func() {
		for {
			var m runtime.MemStats
			runtime.ReadMemStats(&m)
			log.Printf("\nAlloc = %v KB\nStackSys = %v KB\nSys = %v KB \nNumGC = %v\n\n", m.Alloc / 1024, m.StackSys / 1024, m.Sys / 1024, m.NumGC)
			time.Sleep(1 * time.Second)
		}
	}()
}

func WrapCallbackWithLogSentinel(invocationMethod string, callback ow.OpenWhiskCallback) ow.OpenWhiskCallback {

	return func(input json.RawMessage) (interface{}, error) {
		log.Printf("-- OpenWhiskCallback via %s Started --", invocationMethod)
		defer log.Printf("-- OpenWhiskCallback via %s Finished --", invocationMethod)
		return callback(input)
	}
}

func InvokeActionStdIo(callback ow.OpenWhiskCallback) {

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
		log.Printf(format, v...)
		b.numInvocations += 1
	}
}

func CreateBoundedLogger(maxInvocations int) Logger {
	return &BoundedLogger{
		numInvocations: 0,
		maxInvocations: maxInvocations,
	}

}


