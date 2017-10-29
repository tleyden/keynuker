// Licensed to the Apache Software Foundation (ASF) under one or more contributor license agreements;
// and to You under the Apache License, Version 2.0.  See LICENSE in project root for full license + copyright.

package main

import (
	"encoding/json"

	"github.com/tleyden/keynuker"
	"github.com/tleyden/keynuker/keynuker-go-common"
)

func main() {
	keynuker_go_common.LogMemoryUsage()
	keynuker_go_common.RegistorOrInvokeActionStdIo(OpenWhiskCallback)
}

func OpenWhiskCallback(value json.RawMessage) (interface{}, error) {

	var params keynuker.ParamsFetchAwsKeys

	err := json.Unmarshal(value, &params)
	if err != nil {
		return nil, err
	}

	if params.KeyNukerOrg == "" {
		params.KeyNukerOrg = keynuker_go_common.DefaultKeyNukerOrg
	}

	docWrapper, err := keynuker.FetchAwsKeys(params)
	if err != nil {
		return nil, err
	}

	return docWrapper, nil
}
