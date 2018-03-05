// Licensed to the Apache Software Foundation (ASF) under one or more contributor license agreements;
// and to You under the Apache License, Version 2.0.  See LICENSE in project root for full license + copyright.

package keynuker

import (
	"context"
	"fmt"

	f "github.com/fauna/faunadb-go/faunadb"
	"github.com/go-kivik/kivik"

	_ "github.com/go-kivik/couchdb" // The CouchDB driver

	"strings"
	"log"
)

// Write an object specified in params to the underlying database, and return the
// written object back.
func WriteDocToDb(params ParamsWriteDoc) (interface{}, error) {

	switch {
	case params.IsCloudantDb():
		return WriteDocToCloudant(params)
	case params.IsFaunaDb():
		return WriteDocToFauna(params)
	default:
		return nil, fmt.Errorf("Unrecognized DB host: %v", params.Host)
	}

}

func CreateDBClient(params ParamsWriteDoc) (db *kivik.DB, err error) {

	ctx := context.TODO()

	// NOTE regarding http vs https.  I was previously using https, but I needed to revert to http
	// due to the following issue w/ Cloudant: https://gist.github.com/tleyden/d8c69f6cee3a2b32a7a1609539a1dade (filed to Cloudant support)
	dataSourceName := fmt.Sprintf(
		"https://%s:%s@%s",
		params.Username,
		params.Password,
		params.Host,
	)

	client, err := kivik.New(ctx, "couch", dataSourceName)
	if err != nil {
		return nil, err
	}

	return client.DB(ctx, params.DbName)

}

func WriteDocToCloudant(params ParamsWriteDoc) (interface{}, error) {

	ctx := context.TODO()

	db, err := CreateDBClient(params)
	if err != nil {
		return nil, err
	}

	// After maxTries, give up.. this far too many CAS errors than expected under normal operations
	maxTries := 100
	numTries := 0
	var lastErr error

	// Cas loop where we get the latest rev of the doc
	for {

		if numTries >= maxTries {
			return nil, fmt.Errorf("Giving up after %d tries.  Last error: %v", numTries, lastErr)
		}

		numTries++

		fetchedDocRow := db.Get(ctx, params.DocId)

		// Try to get existing doc
		fetchedDoc := map[string]interface{}{}
		scanDocErr := fetchedDocRow.ScanDoc(&fetchedDoc)

		if scanDocErr != nil {
			// Assume this will be the first rev of the doc, so do nothing in this case
		} else {
			// Got an existing doc, update the doc being inserted to be based on this rev
			// we had a valid previous rev
			revRaw, ok := fetchedDoc["_rev"]
			if !ok {
				return nil, fmt.Errorf("Doc does not have _rev field.  Doc: %+v", fetchedDoc)
			}
			rev := revRaw.(string)
			params.Doc["_rev"] = rev
		}

		_, err = db.Put(context.TODO(), params.DocId, params.Doc)
		if err != nil {
			// Assume this is a 409 conflict error
			// TODO: check error status and act accordingly, otherwise will end up in toxic busy loop
			log.Printf("Error putting doc %v into db.  Err: %v.  Assuming 409 conflict error, retrying", params.DocId, err)
			lastErr = err
			continue
		}

		break


	}

	// Read the doc back from the DB
	fetchedDocRow := db.Get(ctx, params.DocId)
	fetchedDoc := map[string]interface{}{}
	scanDocErr := fetchedDocRow.ScanDoc(&fetchedDoc)
	if scanDocErr != nil {
		return nil, scanDocErr
	}

	return fetchedDoc, nil

}

func DeleteDoc(params ParamsWriteDoc, rev string) (newRev string, err error) {

	ctx := context.TODO()

	db, err := CreateDBClient(params)
	if err != nil {
		return "", err
	}

	return db.Delete(ctx, params.DocId, rev)

}

type ParamsWriteDoc struct {
	Username string
	Password string
	Host     string
	DbName   string
	Doc      map[string]interface{}
	DocId    string
}

func (p ParamsWriteDoc) IsCloudantDb() bool {
	return strings.Contains(p.Host, "cloudant")
}

func (p ParamsWriteDoc) IsFaunaDb() bool {
	return strings.Contains(p.Host, "fauna")
}

// Experimenting w/ FaunaDB

var (
	data = f.ObjKey("data")
	ref  = f.ObjKey("ref")
)

func WriteDocToFauna(params ParamsWriteDoc) (interface{}, error) {

	client := f.NewFaunaClient(params.Password)

	// Concats a string .. just testing
	res, err := client.Query(f.Concat([]string{"Hello", "World"}))
	if err != nil {
		panic(err)
	}
	fmt.Println(res)

	// TODO: how do I lazily create a class if it doesn't already exist?  Just try it and ignore 400 error?
	//val, err := client.Query(f.CreateClass(f.Obj{"name": params.DocId}))
	//if err != nil {
	//	return nil, err
	//}

	val, err := client.Query(
		f.Create(
			f.Class(params.DocId),
			f.Obj{"data": f.Obj(params.Doc)},
		),
	)

	if err != nil {
		return nil, err
	}

	fmt.Printf("Val: %v\n", val)

	// TODO: I want to read the doc back out of the db, how do I get the ref?

	// Read the doc back
	val, err = client.Query(f.Get(f.Ref("classes/TestDoc/172286141322494465")))
	if err != nil {
		return nil, err
	}

	readTestDoc := map[string]interface{}{}
	//err = val.Get(&readTestDoc)
	//if err != nil {
	//	return nil, err
	//}

	_ = val.At(data).Get(&readTestDoc)

	// fmt.Printf("readTestDoc: %+v", readTestDoc)

	readTestDoc2 := f.ObjectV{}
	_ = val.At(data).Get(&readTestDoc2)

	readTestDoc2Bytes, err := readTestDoc2.MarshalJSON()
	if err != nil {
		return nil, err
	}
	fmt.Printf("readTestDoc2Bytes: %v", string(readTestDoc2Bytes))

	return readTestDoc, nil
}
