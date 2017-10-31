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

// UseDockerSkeleton: true or false.
//
// - True to use https://hub.docker.com/r/tleyden5iwx/openwhisk-dockerskeleton/ (default)
// - False to directly build an image and push to dockerhub
//
// There are two reasons you might want to set this to False:
//   1. Want full control of all the code, as opposed to trusting the code in https://hub.docker.com/r/tleyden5iwx/openwhisk-dockerskeleton/
//   2. Suspect there is an issue with the actionproxy.py wrapper code in openwhisk-dockerskeleton, and want to compare behavior.
//
// If you set to False, you will need to have docker locally installed and a few extra environment
// variables set.  This needs to match the value in install.py.
//func UseDockerSkeleton() bool {
//
//	val, ok := os.LookupEnv(EnvVarKeyNukerInstallUseDockerSkeleton)
//	if !ok {
//		return true
//	}
//	if strings.ToLower(val) == "false" {
//		return false
//	}
//	return true
//
//
//}
