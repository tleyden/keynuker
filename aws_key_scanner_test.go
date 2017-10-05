// Licensed to the Apache Software Foundation (ASF) under one or more contributor license agreements;
// and to You under the Apache License, Version 2.0.  See LICENSE in project root for full license + copyright.

package keynuker

import (
	"fmt"
	"testing"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/couchbaselabs/go.assert"
	"bytes"
)

func TestAwsKeyScanner(t *testing.T) {

	leakedKey := "FakeAccessKey"
	accessKeyMetadata := []FetchedAwsAccessKey{
		{
			AccessKeyId: aws.String(leakedKey),
			UserName:    aws.String("fakeuser@test.com"),
		},
	}

	// Single leaked key
	eventContent := []byte(fmt.Sprintf(`"patch":"@@ -1,2 +1,2 @@\n \n-Oops I just leaked %s!`, leakedKey))
	leakedKeys, _, err := Scan(accessKeyMetadata, eventContent)
	assert.True(t, err == nil)
	assert.True(t, len(leakedKeys) == 1)
	assert.True(t, *leakedKeys[0].AccessKeyId == leakedKey)

	// No leaked key
	eventContent = []byte(fmt.Sprintf(`"patch":"@@ -1,2 +1,2 @@\n \n- I'm not a leaker'`))
	noLeakedKeys, _, err2 := Scan(accessKeyMetadata, eventContent)
	assert.True(t, err2 == nil)
	assert.True(t, len(noLeakedKeys) == 0)

}

func BenchmarkAwsKeyScanner(b *testing.B) {

	// Setup keys and content
	leakedKeyPrefix := "FakeAccessKey"
	accessKeyMetadata := []FetchedAwsAccessKey{}

	contentBuffer := bytes.NewBufferString("")

	numIterations := 1000
	for i := 0; i < numIterations; i++ {
		leakedKey := fmt.Sprintf("%s-%d", i, leakedKeyPrefix)
		fetchedAwsAccessKey := FetchedAwsAccessKey{
			AccessKeyId: aws.String(leakedKey),
			UserName:    aws.String("fakeuser@test.com"),
		}
		accessKeyMetadata = append(accessKeyMetadata, fetchedAwsAccessKey)

		contentBuffer.WriteString(fmt.Sprintf(`"patch":"@@ -1,2 +1,2 @@\n \n-Oops I just leaked %s! `, leakedKey))
	}

	eventContent := contentBuffer.Bytes()

	// Reset benchmark timer so that the results aren't skewed by the time taken during setup
	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		_, _, err := Scan(accessKeyMetadata, eventContent)
		if err != nil {
			b.Fatalf("Unexpected error scanning content.  Err: %v", err)
		}

	}
}