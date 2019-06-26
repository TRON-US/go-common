package conversion

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

var (
	testEpochStampMilli   int64 = 1555555555555
	testEpochStampMicro   int64 = 1555555555555555
	testMilliTimeString         = "2019-04-18 02:45:55.555 +0000 UTC"
	testMicroTimeString         = "2019-04-18 02:45:55.555555 +0000 UTC"
	testTimeMilli               = time.Date(2019, time.April, 18, 2, 45, 55, 555000000, time.UTC)
	testTimeMicro               = time.Date(2019, time.April, 18, 2, 45, 55, 555555000, time.UTC)
	testPGTimeStringMilli       = "2019-04-18 02:45:55.555"
	testPGTimeStringMicro       = "2019-04-18 02:45:55.555555"
	testPGMilliTimeString       = "2019-04-18 02:45:55.555 +0000"
)

func TestEpochStamp2MilliTime(t *testing.T) {
	res := EpochStamp2MilliTime(testEpochStampMilli)
	assert.Equal(t, testMilliTimeString, res.UTC().String(), "EpochStamp2MilliTime failed to convert unix time to local timestamp")
}

func TestEpochStamp2MicroTime(t *testing.T) {
	res := EpochStamp2MicroTime(testEpochStampMicro)
	assert.Equal(t, testMicroTimeString, res.UTC().String(), "EpochStamp2MicroTime failed to convert unix time to local timestamp")
}

func TestTime2MilliStamp(t *testing.T) {
	res := Time2MilliStamp(testTimeMilli)
	assert.Equal(t, testEpochStampMilli, res, "Time2MilliStamp failed to convert local timestamp to unix time")
}

func TestTime2MircoStamp(t *testing.T) {
	res := Time2MircoStamp(testTimeMicro)
	assert.Equal(t, testEpochStampMicro, res, "Time2MircoStamp failed to convert local timestamp to unix time")
}

func TestTimeString2MilliTime(t *testing.T) {
	res, err := TimeString2MilliTime(testPGTimeStringMilli)
	assert.NoError(t, err, "TimeString2MilliTime failed to parse postgres timestamp")
	assert.Equal(t, testTimeMilli, res, "TimeString2MilliTime failed to return the right time.Time")
	_, err = TimeString2MilliTime(testMilliTimeString)
	assert.Error(t, err, "TimeString2MilliTime should fail to parse wrong timestamp format")
}

func TestTimeString2MicroTime(t *testing.T) {
	res, err := TimeString2MicroTime(testPGTimeStringMicro)
	assert.NoError(t, err, "TimeString2MicroTime failed to parse postgres timestamp")
	assert.Equal(t, testTimeMicro, res, "TimeString2MicroTime failed to return the right time.Time")
	_, err = TimeString2MicroTime(testMicroTimeString)
	assert.Error(t, err, "TimeString2MicroTime should fail to parse wrong timestamp format")
}

func TestEpochStamp2PGMilliTimeString(t *testing.T) {
	res := EpochStamp2PGMilliTimeString(testEpochStampMilli)
	assert.Equal(t, testPGMilliTimeString, res, "EpochStamp2PGMilliTimeString failed to convert unix time to pg time")
}
