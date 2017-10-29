// Licensed to the Apache Software Foundation (ASF) under one or more contributor license agreements;
// and to You under the Apache License, Version 2.0.  See LICENSE in project root for full license + copyright.

package main

import (
	"encoding/json"

	"github.com/tleyden/keynuker"
	"github.com/tleyden/keynuker/keynuker-go-common"
)

func main() {
	defer keynuker_go_common.LogMemoryUsage()
	keynuker_go_common.LogMemoryUsageLoop()
	keynuker_go_common.RegistorOrInvokeActionStdIo(OpenWhiskCallback)
}

func OpenWhiskCallback(value json.RawMessage) (interface{}, error) {

	var params keynuker.ParamsNukeLeakedAwsKeys

	err := json.Unmarshal(value, &params)
	if err != nil {
		return nil, err
	}

	if err := params.Validate(); err != nil {
		return keynuker.ParamsNukeLeakedAwsKeys{}, err
	}

	docWrapper, err := keynuker.NukeLeakedAwsKeys(params)
	if err != nil {
		return nil, err
	}

	return docWrapper, nil
}
