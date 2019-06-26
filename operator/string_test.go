package operator

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	testString0 = "a"
	testString1 = ""
	testArray01 = []string{"a", "aa", "b"}
	testArray02 = []string{"a0", "aa", "b"}
	testArray03 = []string{}
	testArray04 = []string{"a", "aa", "b"}
	testArray05 = []string{"b", "aa", "a"}
	testArray06 = []string{"a", "aa", "b", "b", "b"}
)

func TestStringInSliceIncluded(t *testing.T) {
	result := StringInSlice(testString0, testArray01)
	assert.Equal(t, true, result, "StringInSlice function wrong result on included case")
}

func TestStringInSliceExcluded(t *testing.T) {
	result := StringInSlice(testString0, testArray02)
	assert.Equal(t, false, result, "StringInSlice function wrong result on excluded case")
}

func TestStringInSliceEmpty(t *testing.T) {
	result := StringInSlice(testString0, testArray03)
	assert.Equal(t, false, result, "StringInSlice function wrong result on empty slice")
	result = StringInSlice(testString1, testArray03)
	assert.Equal(t, false, result, "StringInSlice function wrong result on empty string in empty slice")
}

func TestStringSliceEqualEqual(t *testing.T) {
	result := StringSliceEqual(testArray01, testArray01)
	assert.Equal(t, true, result, "StringSliceEqual function cannot return equal on same slices")
	result = StringSliceEqual(testArray01, testArray04)
	assert.Equal(t, true, result, "StringSliceEqual function cannot return equal on same slice elements")
	result = StringSliceEqual(testArray03, []string{})
	assert.Equal(t, true, result, "StringSliceEqual function cannot return equal on same empty slices")
	result = StringSliceEqual(nil, nil)
	assert.Equal(t, true, result, "StringSliceEqual function cannot return equal on same nil slices")
}

func TestStringSliceEqualNotEqual(t *testing.T) {
	result := StringSliceEqual(testArray01, testArray02)
	assert.Equal(t, false, result, "StringSliceEqual function cannot return not equal on different slices")
	result = StringSliceEqual(testArray01, testArray03)
	assert.Equal(t, false, result, "StringSliceEqual function cannot return not equal on comparing to empty slice")
	result = StringSliceEqual(testArray01, testArray05)
	assert.Equal(t, false, result, "StringSliceEqual function cannot return not equal on same slice elements but different ordering")
	result = StringSliceEqual(testArray01, testArray06)
	assert.Equal(t, false, result, "StringSliceEqual function cannot return not equal on one slice being subset of the other")
	result = StringSliceEqual(testArray01, nil)
	assert.Equal(t, false, result, "StringSliceEqual function cannot return not equal on one nil slice")
	result = StringSliceEqual(testArray03, nil)
	assert.Equal(t, false, result, "StringSliceEqual function cannot return not equal on one nil and one empty slice")
}
