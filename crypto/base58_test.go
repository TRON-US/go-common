package crypto

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	testHexString    = "41a4a9bd6deb6f5b685cde2b27238f68a1e763160d"
	testBase58String = "TQysAFw3jfhEPtUWiYDPqwBxiYuvti27GF"
)

// Test decode58 function.
func TestDecode58CheckSuccess(t *testing.T) {
	decode, err := Decode58Check(&testBase58String)
	assert.NoError(t, err, "Decode58Check address failed")
	assert.Equal(t, testHexString, *decode, "Decode58Check function failed")
}

func TestDecode58CheckNil(t *testing.T) {
	decode, err := Decode58Check(nil)
	assert.Nil(t, decode, "Decode58Check should return nil result")
	assert.Nil(t, err, "Decode58Check should return nil error")
}

func TestDecode58CheckLength(t *testing.T) {
	short := "hello"
	_, err := Decode58Check(&short)
	assert.Equal(t, ErrDecodeLength, err, "Decode58Check should return length error")
}

func TestDecode58CheckFormat(t *testing.T) {
	invalid := testBase58String[:len(testBase58String)-1] + "G"
	_, err := Decode58Check(&invalid)
	assert.Equal(t, ErrDecodeCheck, err, "Decode58Check should return check error")
}

// Test encode58 function.
func TestEncode58Check(t *testing.T) {
	encode, err := Encode58Check(&testHexString)
	assert.NoError(t, err, "Encode58Check address failed")
	assert.Equal(t, testBase58String, *encode, "Encode58Check function failed")
}

func TestEncode58CheckNil(t *testing.T) {
	encode, err := Encode58Check(nil)
	assert.Nil(t, encode, "Encode58Check should return nil result")
	assert.Nil(t, err, "Encode58Check should return nil error")
}

func TestEncode58CheckHex(t *testing.T) {
	invalid := "31ZZZZZ"
	_, err := Encode58Check(&invalid)
	assert.Error(t, err, "Encode58Check hex check failed")
}

func TestEncode58CheckLen(t *testing.T) {
	encode, err := Encode58CheckLen(&testHexString, len(testBase58String))
	assert.NoError(t, err, "Encode58CheckLen address failed")
	assert.Equal(t, testBase58String, *encode, "Encode58CheckLen function failed")
	_, err = Encode58CheckLen(&testHexString, 0)
	assert.Error(t, err, "Encode58CheckLen len failed")
	_, err = Encode58CheckLen(&testHexString, 1000)
	assert.Error(t, err, "Encode58CheckLen len failed")
}

func TestEncode58CheckLenNil(t *testing.T) {
	encode, err := Encode58CheckLen(nil, 0)
	assert.Nil(t, encode, "Encode58CheckLen should return nil result")
	assert.Nil(t, err, "Encode58CheckLen should return nil error")
	encode, err = Encode58CheckLen(nil, 10)
	assert.Nil(t, encode, "Encode58CheckLen should return nil result on wrong len")
	assert.Nil(t, err, "Encode58CheckLen should return nil error on wrong len")
}

func TestEncode58CheckLenHex(t *testing.T) {
	invalid := "31ZZZZZ"
	_, err := Encode58CheckLen(&invalid, 0)
	assert.Error(t, err, "Encode58CheckLen hex check failed")
	_, err = Encode58CheckLen(&invalid, len(invalid))
	assert.Error(t, err, "Encode58CheckLen hex check failed")
}
