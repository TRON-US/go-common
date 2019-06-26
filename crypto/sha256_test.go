package crypto

import (
	"encoding/hex"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	testHashEncodeString = "9f86d081884c7d659a2feaa0c55ad015a3bf4f1b2b0b822cd15d6c15b0f00a08"
	testHashDecodeString = "test"
)

// Test hash256 function.
func TestHash(t *testing.T) {
	hash := Hash([]byte(testHashDecodeString))
	assert.Equal(t, testHashEncodeString, hex.EncodeToString(hash), "hash256 function failed")
}
