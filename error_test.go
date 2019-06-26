package utils

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	testMsg    = "I am an error"
	testErrMsg = "big exception!"
	testError  = fmt.Errorf(testErrMsg)
)

func TestE(t *testing.T) {
	res := E(testMsg, testError)
	expected := fmt.Errorf("%s: [%s]", testMsg, testErrMsg)
	assert.Equal(t, expected, res, "E error wrap message and error failed")
}
