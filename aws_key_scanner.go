// Licensed to the Apache Software Foundation (ASF) under one or more contributor license agreements;
// and to You under the Apache License, Version 2.0.  See LICENSE in project root for full license + copyright.

package keynuker

import (
	"regexp"

	"log"

	"github.com/derekparker/trie"
)

type ScannerState int

const (
	ScannerStateInToken ScannerState = iota
	ScannerStateOutsideToken
)

func Scan(accessKeysToScan []FetchedAwsAccessKey, content []byte) (leaks []FetchedAwsAccessKey, nearbyContent []byte, err error) {
	return ScanViaRegexLoop(accessKeysToScan, content)
}

// This is grossly inefficient but it passes all of the tests
func ScanViaRegexLoop(accessKeysToScan []FetchedAwsAccessKey, content []byte) (leaks []FetchedAwsAccessKey, nearbyContent []byte, err error) {

	for _, keyMetadata := range accessKeysToScan {

		key := *keyMetadata.AccessKeyId

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

// Scan the input in a single pass and use a trie prefix match to figure out if any aws keys match.
// This is approx 2x faster than ScanViaRegexLoop, but it still feels slow.
// TODO: try using lexmachine and see if it's a lot faster.  Do some post-processing to deal with nested tokens (where one token contains another, which is one of the unit tests)
func ScanViaTrie(accessKeysToScan []FetchedAwsAccessKey, content []byte) (leaks []FetchedAwsAccessKey, nearbyContent []byte, err error) {

	debug := false

	uniqueLeaks := map[string]FetchedAwsAccessKey{}

	leaks = []FetchedAwsAccessKey{}

	trie := trie.New()
	exactMatch := map[string]FetchedAwsAccessKey{}

	for _, keyMetadata := range accessKeysToScan {

		// Get the access key id
		key := *keyMetadata.AccessKeyId

		// Add it to the trie to be indexed
		trie.Add(key, keyMetadata)
		exactMatch[key] = keyMetadata

	}

	runes := []rune(string(content))

	currentScannerState := ScannerStateOutsideToken

	scanStartPointer := 0
	scanEndPointer := 1

	for {

		if scanStartPointer >= len(runes) {
			break
		}

		if scanEndPointer > len(runes) {
			scanEndPointer = len(runes)
		}

		currentStateName := ""
		if currentScannerState == ScannerStateInToken {
			currentStateName = "ScannerStateInToken"
		} else {
			currentStateName = "ScannerStateOutsideToken"
		}
		if debug {
			log.Printf("runes[%d:%d] - state: %s", scanStartPointer, scanEndPointer, currentStateName)
		}

		currentRunes := runes[scanStartPointer:scanEndPointer]
		if debug {
			log.Printf("currentRunes: |%v|", string(currentRunes))
		}

		switch trie.HasKeysWithPrefix(string(currentRunes)) {
		case true:

			if debug {
				log.Printf("trie.HasKeysWithPrefix: |%v|.  Switching state -> ScannerStateInToken", string(currentRunes))
			}

			currentScannerState = ScannerStateInToken

			accessKeyMeta, found := exactMatch[string(currentRunes)]
			if found {
				if debug {
					log.Printf("trie exact match: %v", *accessKeyMeta.AccessKeyId)
				}
				uniqueLeaks[string(currentRunes)] = accessKeyMeta
			}

			// Keep consuming runes by expanding the scanning window to be one rune larger
			scanEndPointer += 1

		case false:  // trie.HasKeysWithPrefix(string(currentRunes)) == false

			switch currentScannerState {
			case ScannerStateInToken:

				// There are no prefix matches, and we were in a possible token, so we're out of the token now

				currentScannerState = ScannerStateOutsideToken

				// Suppose the goal is to find these two keys [ROX, OA5]
				// And the input text is: "ROA5XYZ"
				// It will be in a token match state for the "RO" since the prefix will match "ROX",
				// but when it gets to the "A" in "ROA5", the "ROA" prefix will not match anything.
				// At this point, slide the start of the window one rune to the right so it's now aligned to "OA5"
				scanStartPointer += 1

			case ScannerStateOutsideToken:

				// There are no prefix matches, and we weren't in a token

				if scanEndPointer > (scanStartPointer + 1) {

					// This can happen if we just recently partially or fully matched a token, and have a wide scanning window,
					// but then on reading the next rune it doesn't match any tokens, but still has this wide scanning window.
					// Shrink the window back down until it's reading a single character.  It does do some (probably)
					// unecessary scanning
					scanEndPointer -= 1

				} else {

					// Normal case
					// Slide the scanning window one rune to the right
					scanStartPointer += 1
					scanEndPointer += 1
				}

			}

		}

	}

	for _, accessKeyMeta := range uniqueLeaks {
		leaks = append(leaks, accessKeyMeta)
	}

	return leaks, nil, nil
}
