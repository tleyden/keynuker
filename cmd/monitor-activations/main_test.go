package main

import (
	"testing"
	"encoding/json"
)

func TestOpenWhiskCallback(t *testing.T) {

	OpenWhiskCallback(json.RawMessage{})

}
