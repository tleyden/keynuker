// Licensed to the Apache Software Foundation (ASF) under one or more contributor license agreements;
// and to You under the Apache License, Version 2.0.  See LICENSE in project root for full license + copyright.

package keynuker

import (
	"fmt"
	"testing"

	"log"
	"math/rand"
	"time"

	"bytes"
	"io/ioutil"

	"sort"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/couchbaselabs/go.assert"
)

var letterRunes = []rune("123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ")
var numKeys = 350

func TestAwsKeyScanner(t *testing.T) {

	leakedKey := "FakeAccessKey"
	leakedKey2 := "FakeAccessKey2"
	accessKeyMetadata := []FetchedAwsAccessKey{
		{
			AccessKeyId: aws.String(leakedKey),
			UserName:    aws.String("fakeuser@test.com"),
		},
		{
			AccessKeyId: aws.String(leakedKey2),
			UserName:    aws.String("fakeuser2@test.com"),
		},
	}

	// Single leaked key
	log.Printf("--------------------- Single leaked key")
	eventContent := []byte(fmt.Sprintf(`"patch":"@@ -1,2 +1,2 @@\n \n-Oops I just leaked %s!`, leakedKey))
	leakedKeys, _, err := Scan(accessKeyMetadata, eventContent)
	assert.True(t, err == nil)
	assert.True(t, len(leakedKeys) == 1)
	assert.True(t, contains(accessKeyMetadata[0:1], leakedKeys))

	// Two non-adjacent leaked keys
	log.Printf("--------------------- Two non-adjacent leaked keys")
	eventContent = []byte(fmt.Sprintf(`"patch":"@@ -1,2 +1,2 @@\n \n-Oops I just leaked %s %s!`, leakedKey, leakedKey2))
	leakedKeys, _, err = Scan(accessKeyMetadata, eventContent)
	assert.True(t, err == nil)
	assert.True(t, len(leakedKeys) == 2)
	assert.True(t, contains(accessKeyMetadata, leakedKeys))

	// Two non-adjacent leaked keys
	log.Printf("--------------------- Two adjacent leaked keys")
	eventContent = []byte(fmt.Sprintf(`"patch":"@@ -1,2 +1,2 @@\n \n-Oops I just leaked %s%s!`, leakedKey, leakedKey2))
	leakedKeys, _, err = Scan(accessKeyMetadata, eventContent)
	assert.True(t, err == nil)
	assert.True(t, len(leakedKeys) == 2)
	assert.True(t, contains(accessKeyMetadata, leakedKeys))

	// Two leaked keys, leaked twice each -- should de-dupe
	log.Printf("--------------------- Two leaked keys, leaked twice each -- should de-dupe")
	eventContent = []byte(fmt.Sprintf(`"patch":"@@ -1,2 +1,2 @@\n \n-Oops I just leaked %s%s and again %s%s!`,
		leakedKey, leakedKey2, leakedKey2, leakedKey))
	leakedKeys, _, err = Scan(accessKeyMetadata, eventContent)
	assert.True(t, err == nil)
	assert.True(t, len(leakedKeys) == 2)
	assert.True(t, contains(accessKeyMetadata, leakedKeys))

	// No leaked key
	log.Printf("--------------------- No leaked key")
	eventContent = []byte(fmt.Sprintf(`"patch":"@@ -1,2 +1,2 @@\n \n- I'm not a leaker'`))
	noLeakedKeys, _, err2 := Scan(accessKeyMetadata, eventContent)
	assert.True(t, err2 == nil)
	assert.True(t, len(noLeakedKeys) == 0)

}

func TestAwsKeyScannerOverlappingKeys(t *testing.T) {

	leakedKeyInner := "FakeAccessKey"
	leakedKeyOuter := "OuterFakeAccessKeyOuter"
	accessKeyMetadata := []FetchedAwsAccessKey{
		{
			AccessKeyId: aws.String(leakedKeyInner),
			UserName:    aws.String("fakeuser@test.com"),
		},
		{
			AccessKeyId: aws.String(leakedKeyOuter),
			UserName:    aws.String("fakeuser2@test.com"),
		},
	}

	// Leak the outer key, expect to find both
	eventContent := []byte(fmt.Sprintf(`"patch":"@@ -1,2 +1,2 @@\n \n-Oops I just leaked %s!`, leakedKeyOuter))
	leakedKeys, _, err := Scan(accessKeyMetadata, eventContent)
	assert.True(t, err == nil)
	assert.True(t, len(leakedKeys) == 2)

	assert.True(t, contains(accessKeyMetadata, leakedKeys))

}

func contains(expected, actual []FetchedAwsAccessKey) bool {

	if len(expected) != len(actual) {
		return false
	}

	actualKeys := []string{}
	expectedKeys := []string{}

	for i, fetchedAwsAccessKey := range actual {
		actualKeys = append(actualKeys, *fetchedAwsAccessKey.AccessKeyId)
		expectedKeys = append(expectedKeys, *expected[i].AccessKeyId)
	}

	sort.Strings(actualKeys)
	sort.Strings(expectedKeys)

	for i, expectedKey := range expectedKeys {
		actualKey := actualKeys[i]
		if expectedKey != actualKey {
			return false
		}
	}

	return true
}

func TestAwsKeyScannerLargeFileManyKeys(t *testing.T) {

	largeFileWithKeys, accessKeyMetadata, err := largeFileContentAndKeys()
	if err != nil {
		t.Fatalf("Error: %v", err)
	}

	leaks, _, err := Scan(accessKeyMetadata, largeFileWithKeys)

	if len(leaks) != numKeys {
		uniqueKeys := map[string]string{}
		for _, accessKey := range accessKeyMetadata {
			uniqueKeys[*accessKey.AccessKeyId] = *accessKey.AccessKeyId
		}
		for _, leak := range leaks {
			_, hasKey := uniqueKeys[*leak.AccessKeyId]
			if hasKey {
				delete(uniqueKeys, *leak.AccessKeyId)
			}
		}
		log.Printf("Undetected leaks")
		for key, _ := range uniqueKeys {
			log.Printf("Undetected key: %v", key)
		}
	}

	assert.Equals(t, len(leaks), numKeys)
}

func TestAwsKeyScannerProblematicTest(t *testing.T) {

	problematicTest, err := Asset("testdata/ProblematicText.txt")
	if err != nil {
		t.Fatalf("Err: %v", err)
	}

	accessKeyMetadata := []FetchedAwsAccessKey{}
	fetchedAwsAccessKey := FetchedAwsAccessKey{
		AccessKeyId: aws.String("OAQWO93BD58KJNGGXRF1E2VY9A6599OFMSGWU34OU-205"),
		UserName:    aws.String("fakeuser@test.com"),
	}
	accessKeyMetadata = append(accessKeyMetadata, fetchedAwsAccessKey)

	fetchedAwsAccessKey2 := FetchedAwsAccessKey{
		AccessKeyId: aws.String("ROA5IUM5ZFWZC85PP11C8XCICWW35PPB888WV6ABKE1HM2FNV2ORWKPCV5SHC4KIE8BLC131G8RS437OPF68SYM6T2XCGTZRF2AM24T2SP-202"),
		UserName:    aws.String("fakeuser@test.com"),
	}
	accessKeyMetadata = append(accessKeyMetadata, fetchedAwsAccessKey2)

	leaks, _, err := Scan(accessKeyMetadata, problematicTest)
	assert.Equals(t, len(leaks), 1)

}

func BenchmarkAwsKeyScanner(b *testing.B) {

	largeFileWithKeys, accessKeyMetadata, err := largeFileContentAndKeys()
	if err != nil {
		b.Fatalf("Error: %v", err)
	}
	// Reset benchmark timer so that the results aren't skewed by the time taken during setup
	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		now := time.Now()
		_, _, err := Scan(accessKeyMetadata, largeFileWithKeys)
		delta := time.Since(now)
		log.Printf("Scanned content in %v", delta)
		if err != nil {
			b.Fatalf("Unexpected error scanning content.  Err: %v", err)
		}

	}
}

// Inject the keys into the content and return it
func injectKeys(content []byte, accessKeys []string) []byte {

	if len(content) < len(accessKeys) {
		panic(fmt.Sprintf("len(content) < len(accessKeys)"))
	}

	// If we have more content bytes than keys (usual case), interject the keys
	// roughly evenly within the content

	// Make a new slice that will be returned, which has enough capacity
	// First figure out how much content the accessKeys slice contains.
	accessKeysRunes := 0
	for _, accessKey := range accessKeys {
		accessKeysRunes += len(accessKey)
	}

	contentPointer := 0

	resultContent := bytes.NewBufferString("")

	// Find the "intersperse" distance between normal content and keys to be injected
	intersperseBytes := len(content) / len(accessKeys)

	for _, accessKey := range accessKeys {

		resultContent.WriteString(accessKey)
		endRange := contentPointer + intersperseBytes
		if endRange > len(content) {
			endRange = len(content) - 1
		}
		resultContent.Write(content[contentPointer:endRange])
		contentPointer += intersperseBytes
	}

	return resultContent.Bytes()

}

func largeFileContentAndKeys() (contentWithKeys []byte, accessKeyMetadata []FetchedAwsAccessKey, err error) {

	accessKeyMetadata = []FetchedAwsAccessKey{}
	accessKeys := []string{}

	uniqueKeys := map[string]string{}

	for i := 0; i < numKeys; i++ {
		keyLength := RandIntRange(4, 128)
		leakedKeyPrefix := RandStringRunes(keyLength)
		leakedKey := fmt.Sprintf("%s-%d", leakedKeyPrefix, i)
		// log.Printf("leakedKey: %v", leakedKey)
		fetchedAwsAccessKey := FetchedAwsAccessKey{
			AccessKeyId: aws.String(leakedKey),
			UserName:    aws.String("fakeuser@test.com"),
		}
		accessKeyMetadata = append(accessKeyMetadata, fetchedAwsAccessKey)
		accessKeys = append(accessKeys, leakedKey)
		_, hasKey := uniqueKeys[leakedKey]
		if hasKey {
			return nil, nil, fmt.Errorf("Tried to add duplicate key: %v", leakedKey)
		}
		uniqueKeys[leakedKey] = leakedKey

	}

	largeFile, err := Asset("testdata/largefile.txt")
	if err != nil {
		return nil, nil, fmt.Errorf("Unable to load large commit asset: %v", err)
	}

	// Multiply large file to take it from 1 MB -> 10 MB
	multipliedLargeFile := bytes.NewBuffer(largeFile)
	for i := 0; i < 10; i++ { // TEMP SET to 1.  Reset to 10 for benchmarking
		multipliedLargeFile.Write(largeFile)
	}

	contentWithKeys = injectKeys(multipliedLargeFile.Bytes(), accessKeys)

	ioutil.WriteFile("/tmp/injectedKeys", contentWithKeys, 0644)

	return contentWithKeys, accessKeyMetadata, nil

}

func RandIntRange(min, max int) int {
	return rand.Intn(max-min) + min
}

func RandStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}
