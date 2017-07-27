// Licensed to the Apache Software Foundation (ASF) under one or more contributor license agreements;
// and to You under the Apache License, Version 2.0.  See LICENSE in project root for full license + copyright.

package main

import (
	"encoding/json"

	"github.com/tleyden/keynuker"
	"github.com/tleyden/keynuker/keynuker-github"
	"github.com/tleyden/keynuker/keynuker-go-common"
	"github.com/tleyden/ow"
)

// Scan Github user events for AWS keys

func main() {

	switch keynuker_go_common.StdioFlagPresent() {
	case true:
		keynuker_go_common.InvokeActionStdIo(OpenWhiskCallback)
	default:
		ow.RegisterAction(OpenWhiskCallback)
	}

}

func OpenWhiskCallback(value json.RawMessage) (interface{}, error) {

	var params keynuker.ParamsScanGithubUserEventsForAwsKeys

	err := json.Unmarshal(value, &params)
	if err != nil {
		return nil, err
	}

	if err := params.Validate(); err != nil {
		return params, err
	}

	params = params.WithDefaults()

	fetcher := keynuker_github.NewGoGithubUserEventFetcher(params.GithubAccessToken)

	scanner := keynuker.NewGithubUserEventsScanner(fetcher)

	docWrapper, err := scanner.ScanAwsKeys(params)
	if err != nil {
		return nil, err
	}

	return docWrapper, nil
}
