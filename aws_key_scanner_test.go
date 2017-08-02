// Licensed to the Apache Software Foundation (ASF) under one or more contributor license agreements;
// and to You under the Apache License, Version 2.0.  See LICENSE in project root for full license + copyright.

package keynuker

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/iam"
	"github.com/couchbaselabs/go.assert"
	"testing"
)

func TestAwsKeyScanner(t *testing.T) {

	leakedKey := "FakeAccessKey"
	accessKeyMetadata := []iam.AccessKeyMetadata{
		iam.AccessKeyMetadata{
			AccessKeyId: aws.String(leakedKey),
			UserName:    aws.String("fakeuser@test.com"),
		},
	}

	keyScanner := NewAwsKeyScanner()

	// Single leaked key
	eventContent := []byte(fmt.Sprintf(`"patch":"@@ -1,2 +1,2 @@\n \n-Oops I just leaked %s!`, leakedKey))
	leakedKeys, _, err := keyScanner.Scan(accessKeyMetadata, eventContent)
	assert.True(t, err == nil)
	assert.True(t, len(leakedKeys) == 1)
	assert.True(t, *leakedKeys[0].AccessKeyId == leakedKey)

	// No leaked key
	eventContent = []byte(fmt.Sprintf(`"patch":"@@ -1,2 +1,2 @@\n \n- I'm not a leaker'`))
	noLeakedKeys, _, err2 := keyScanner.Scan(accessKeyMetadata, eventContent)
	assert.True(t, err2 == nil)
	assert.True(t, len(noLeakedKeys) == 0)

}
