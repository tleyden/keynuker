// Licensed to the Apache Software Foundation (ASF) under one or more contributor license agreements;
// and to You under the Apache License, Version 2.0.  See LICENSE in project root for full license + copyright.

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

	var params keynuker.ParamsPostNukeNotifier

	err := json.Unmarshal(value, &params)
	if err != nil {
		return nil, err
	}

	resultDoc, err := keynuker.SendPostNukeMailgunNotifications(params)
	if err != nil {
		return nil, err
	}

	docId := keynuker_go_common.GenerateDocId(
		keynuker_go_common.DocIdPrefixGithubEventCheckpoints,
		params.KeyNukerOrg,
	)

	docWrapper := keynuker.DocumentWrapperPostNukeNotifier {
		DocId: docId,
		Doc: resultDoc,
	}

	return docWrapper, nil

}
