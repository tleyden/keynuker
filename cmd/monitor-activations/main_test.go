package main

import (
	"encoding/json"
	"testing"
)

func TestOpenWhiskCallback(t *testing.T) {

	OpenWhiskCallback(json.RawMessage{})

}
