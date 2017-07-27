// Licensed to the Apache Software Foundation (ASF) under one or more contributor license agreements;
// and to You under the Apache License, Version 2.0.  See LICENSE in project root for full license + copyright.

package keynuker_github

import (
	"regexp"

	"github.com/aws/aws-sdk-go/service/iam"
)

type AwsKeyScanner struct {
}

func NewAwsKeyScanner() *AwsKeyScanner {
	return &AwsKeyScanner{}
}

// Scan the content to see if any of the aws keys are present
func (aks AwsKeyScanner) Scan(accessKeyMetadata []iam.AccessKeyMetadata, content []byte) (leaks []iam.AccessKeyMetadata, nearbyContent []byte, err error) {

	for _, keyMetadata := range accessKeyMetadata {

		key := *keyMetadata.AccessKeyId

		// TODO: this is grossly inefficient.  Convert to a lexer based approach
		r, err := regexp.Compile(key)
		if err != nil {
			return nil, nil, err
		}

		if r.Match(content) {
			leaks = append(leaks, keyMetadata)
		}

	}

	return leaks, nil, nil

}
