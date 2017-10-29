// Licensed to the Apache Software Foundation (ASF) under one or more contributor license agreements;
// and to You under the Apache License, Version 2.0.  See LICENSE in project root for full license + copyright.

package main

import (
	"encoding/json"
	"fmt"

	"github.com/tleyden/keynuker"
	"github.com/tleyden/keynuker/keynuker-go-common"
)

func main() {
	keynuker_go_common.LogMemoryUsage()
	keynuker_go_common.RegistorOrInvokeActionStdIo(OpenWhiskCallback)
}

func OpenWhiskCallback(value json.RawMessage) (interface{}, error) {

	var params keynuker.ParamsGithubUserAggregator

	err := json.Unmarshal(value, &params)
	if err != nil {
		return nil, err
	}

	// Must have a github access token
	if params.GithubAccessToken == "" {
		return nil, fmt.Errorf("You must supply the GithubAccessToken")
	}

	// If no keynuker org is specified, use "default"
	if params.KeyNukerOrg == "" {
		params.KeyNukerOrg = keynuker_go_common.DefaultKeyNukerOrg
	}

	docWrapper, err := keynuker.AggregateGithubUsers(params)
	if err != nil {
		return nil, err
	}

	return docWrapper, nil

}
