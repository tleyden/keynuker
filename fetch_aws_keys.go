// Licensed to the Apache Software Foundation (ASF) under one or more contributor license agreements;
// and to You under the Apache License, Version 2.0.  See LICENSE in project root for full license + copyright.

package keynuker

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/iam"
	"github.com/tleyden/keynuker/keynuker-go-common"
	"log"
)

// Look up all the AWS keys associated with the AWS account corresponding to AwsAccessKeyId
// and return a document suitable for sticking into a Cloudant database.
//
// This is meant to be run in the context of an OpenWhisk Action
// (see https://tleyden.github.io/blog/2017/07/02/openwhisk-action-sequences/)
// and so nothing else except the JSON content can be written to standard output.
func FetchAwsKeys(params ParamsFetchAwsKeys) (docWrapper DocumentWrapperFetchAwsKeys, err error) {

	docId := keynuker_go_common.GenerateDocId(
		keynuker_go_common.DocIdPrefixAwsKeys,
		params.KeyNukerOrg,
	)

	// Create output document + wrapepr
	doc := DocumentFetchAwsKeys{
		Id:                docId,
		AwsAccessKeyId:    params.AwsAccessKeyId,
		AccessKeyMetadata: []iam.AccessKeyMetadata{},
	}

	log.Printf("create AWS session with params: %+v", params)

	// Create AWS session
	sess, err := session.NewSession(&aws.Config{
		Credentials: credentials.NewCredentials(
			&credentials.StaticProvider{Value: credentials.Value{
				AccessKeyID:     params.AwsAccessKeyId,
				SecretAccessKey: params.AwsSecretAccessKey,
			}},
		),
	})
	if err != nil {
		return DocumentWrapperFetchAwsKeys{}, fmt.Errorf("Error creating aws session: %v", err)
	}

	// Create IAM client with the session.
	svc := iam.New(sess)

	// Fetch list of IAM users
	iamUsers, err := FetchIAMUsers(svc)
	if err != nil {
		return DocumentWrapperFetchAwsKeys{}, err
	}

	for _, user := range iamUsers {

		listAccessKeysInput := &iam.ListAccessKeysInput{
			UserName: user.UserName,
			MaxItems: aws.Int64(1000),
		}

		listAccessKeysOutput, err := svc.ListAccessKeys(listAccessKeysInput)
		if err != nil {
			return DocumentWrapperFetchAwsKeys{}, fmt.Errorf("Error listing access keys for user: %v.  Err: %v", user, err)
		}

		// Panic if more than 1K results, which is not handled
		if *listAccessKeysOutput.IsTruncated {
			// TODO: remove panic and put in a paginated loop.  Move to tleyden/awsutils + unit tests against mocks
			return DocumentWrapperFetchAwsKeys{}, fmt.Errorf("Output is truncated and this code does not handle it")
		}

		// doc.AccessKeyMetadata = append(doc.AccessKeyMetadata, listAccessKeysOutput.AccessKeyMetadata...)
		for _, accessKeyMetadata := range listAccessKeysOutput.AccessKeyMetadata {
			doc.AccessKeyMetadata = append(doc.AccessKeyMetadata, *accessKeyMetadata)
		}

	}

	docWrapper = DocumentWrapperFetchAwsKeys{
		Doc:   doc,
		DocId: docId,
	}

	return docWrapper, nil

}

func FetchIAMUsers(svc *iam.IAM) (users []*iam.User, err error) {

	// Discover list of IAM users in account
	listUsersInput := &iam.ListUsersInput{
		MaxItems: aws.Int64(1000),
	}
	listUsersOutput, err := svc.ListUsers(listUsersInput)
	if err != nil {
		return nil, fmt.Errorf("Error listing users: %v", err)
	}
	// Panic if more than 1K results, which is not handled
	if *listUsersOutput.IsTruncated {
		// TODO: remove panic and put in a paginated loop.  Move to tleyden/awsutils + unit tests against mocks
		return nil, fmt.Errorf("Output is truncated and this code does not handle it")
	}

	return listUsersOutput.Users, nil

}

type ParamsFetchAwsKeys struct {

	// The aws access key to connect as.  This only needs permissions to list IAM users and access keys,
	// and delete access keys (in the case they are nuked)
	AwsAccessKeyId string

	// The secret access key corresponding to AwsAccessKeyId
	AwsSecretAccessKey string

	// This is the name of the KeyNuker "org/tenant".  Defaults to "default", but allows to be extended multi-tenant.
	KeyNukerOrg string
}

type DocumentFetchAwsKeys struct {
	Id                string `json:"_id"`
	AwsAccountId      string
	AwsAccessKeyId    string
	AccessKeyMetadata []iam.AccessKeyMetadata
}

type DocumentWrapperFetchAwsKeys struct {
	// Serialize into a form that the cloudant db adapter expects
	Doc   DocumentFetchAwsKeys `json:"doc"`
	DocId string               `json:"docid"`
}
