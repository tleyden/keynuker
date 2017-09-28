// Licensed to the Apache Software Foundation (ASF) under one or more contributor license agreements;
// and to You under the Apache License, Version 2.0.  See LICENSE in project root for full license + copyright.

package keynuker

import (
	"regexp"
)


// Scan the content to see if any of the aws keys are present
func Scan(accessKeyMetadata []FetchedAwsAccessKey, content []byte) (leaks []FetchedAwsAccessKey, nearbyContent []byte, err error) {

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
