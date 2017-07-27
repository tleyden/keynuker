// Licensed to the Apache Software Foundation (ASF) under one or more contributor license agreements;
// and to You under the Apache License, Version 2.0.  See LICENSE in project root for full license + copyright.

package keynuker_go_common

type AwsParams struct {

	// The aws access key to connect as.  This only needs permissions to list IAM users and access keys,
	// and delete access keys (in the case they are nuked)
	AwsAccessKeyId string

	// The secret access key corresponding to AwsAccessKeyId
	AwsSecretAccessKey string
}
