package main

import (
	"encoding/json"
	"log"
	"testing"
)

func TestOpenWhiskCallback(t *testing.T) {

	result, err := OpenWhiskCallback(json.RawMessage{})
	log.Printf("result: %v, err: %v", result, err)

}
