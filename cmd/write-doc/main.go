// Licensed to the Apache Software Foundation (ASF) under one or more contributor license agreements;
// and to You under the Apache License, Version 2.0.  See LICENSE in project root for full license + copyright.

package main

import (
	"encoding/json"

	_ "github.com/flimzy/kivik/driver/couchdb" // The CouchDB driver
	"github.com/tleyden/keynuker"
	"github.com/tleyden/keynuker/keynuker-go-common"
)

// This writes/updates a doc in CouchDB.  If the doc doesn't exist, it will write it.  If it does exist,
// it will "clobber" it and create a new revision based on the updated content.

func main() {
	keynuker_go_common.RegistorOrInvokeActionStdIo(OpenWhiskCallback)
}

func OpenWhiskCallback(value json.RawMessage) (interface{}, error) {

	var params keynuker.ParamsWriteDoc

	err := json.Unmarshal(value, &params)
	if err != nil {
		return nil, err
	}

	docWrapper, err := keynuker.WriteDocToDb(params)
	if err != nil {
		return nil, err
	}

	return docWrapper, nil
}
