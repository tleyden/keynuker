package keynuker_go_common

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReadLimited1(t *testing.T) {

	// Create a reader that returns 10000 bytes
	// Call ReadLimited with a 100 byte limit
	// Ensure that only 100 bytes is read

	readerSize := 10000

	buffer := bytes.NewBuffer([]byte{})
	for i := 0; i < readerSize; i++ {
		buffer.Write([]byte{uint8(i)})
	}
	bufferBytes := buffer.Bytes()

	result, err := ReadLimited(buffer, 100)
	assert.True(t, len(result) == 100)
	assert.True(t, result[0] == bufferBytes[0])
	assert.True(t, err == nil)

}

func TestReadLimited2(t *testing.T) {

	// Create a reader that returns 10000 bytes
	// Call ReadLimited with a 100 byte limit
	// Ensure that only 100 bytes is read

	readerSize := 10000

	buffer := bytes.NewBuffer([]byte{})
	for i := 0; i < readerSize; i++ {
		buffer.Write([]byte{uint8(i)})
	}
	bufferBytes := buffer.Bytes()

	result, err := ReadLimited(buffer, readerSize*2)
	assert.True(t, len(result) == readerSize)
	assert.True(t, result[0] == bufferBytes[0])
	assert.True(t, err == nil)

}

func TestReadLimited3(t *testing.T) {

	// Create a reader that returns 10000 bytes
	// Call ReadLimited with a 100 byte limit
	// Ensure that only 100 bytes is read

	readerSize := 100

	buffer := bytes.NewBuffer([]byte{})
	for i := 0; i < readerSize; i++ {
		buffer.Write([]byte{uint8(i)})
	}
	bufferBytes := buffer.Bytes()

	result, err := ReadLimited(buffer, 10)
	assert.True(t, len(result) == 10)
	assert.True(t, result[0] == bufferBytes[0])
	assert.True(t, err == nil)

}
