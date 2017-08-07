// Licensed to the Apache Software Foundation (ASF) under one or more contributor license agreements;
// and to You under the Apache License, Version 2.0.  See LICENSE in project root for full license + copyright.

package keynuker

import (
	"log"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"encoding/json"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/google/go-github/github"
	"github.com/tleyden/keynuker/keynuker-go-common"
)


func TestFindAwsAccount(t *testing.T) {

	monitorAwsAccessKeyId := "TestAwsAccessKeyId"

	targetAwsAccounts := []TargetAwsAccount{
		{
			AwsAccessKeyId:     monitorAwsAccessKeyId,
			AwsSecretAccessKey: "TestAwsSecretAccessKey",
		},
		{
			AwsAccessKeyId:     "TestAwsAccessKeyId2",
			AwsSecretAccessKey: "TestAwsSecretAccessKey2",
		},
	}

	targetAwsAccount, err := FindAwsAccount(targetAwsAccounts, monitorAwsAccessKeyId)
	assert.NoError(t, err, "Unexpected error")
	assert.EqualValues(t, targetAwsAccount.AwsAccessKeyId, monitorAwsAccessKeyId)

}
