package conversion

import (
	"fmt"
	"regexp"
	"time"
)

const (
	// This is really the PG-dumped timestamp format
	PGTimeFormatMilli = "2006-01-02 15:04:05.999"
	PGTimeFormatMicro = "2006-01-02 15:04:05.999999"
)

// EpochStamp2MilliTime converts a UNIX epoch time in millisecond (int64) to a local timestamp
func EpochStamp2MilliTime(timestamp int64) time.Time {
	return time.Unix(0, timestamp*int64(time.Millisecond))
}

// EpochStamp2MicroTime converts a UNIX epoch time in microsecond (int64) to a local timestamp
func EpochStamp2MicroTime(timestamp int64) time.Time {
	return time.Unix(0, timestamp*int64(time.Microsecond))
}

// Time2MilliStamp converts a local timestamp to a UNIX epoch time in millisecond (int64)
func Time2MilliStamp(t time.Time) int64 {
	return t.UnixNano() / int64(time.Millisecond)
}

// Time2MircoStamp converts a local timestamp to a UNIX epoch time in mircosecond (int64)
func Time2MircoStamp(t time.Time) int64 {
	return t.UnixNano() / int64(time.Microsecond)
}

// TimeString2MilliTime converts a postgres raw (millisecond) time string into a local time
func TimeString2MilliTime(s string) (time.Time, error) {
	return time.Parse(PGTimeFormatMilli, s)
}

// TimeString2MicroTime converts a postgres raw (microsecond) time string into a local time
func TimeString2MicroTime(s string) (time.Time, error) {
	return time.Parse(PGTimeFormatMicro, s)
}

func EpochStamp2PGMilliTimeString(tsEpoch int64) string {
	tsString := fmt.Sprintf("%s", EpochStamp2MilliTime(tsEpoch).UTC())
	re := regexp.MustCompile(` [[:word:]]{3}`)
	tsString = re.ReplaceAllString(tsString, "")
	return tsString
}
