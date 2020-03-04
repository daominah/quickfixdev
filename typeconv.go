package quickfixdev

import (
	"time"
)

// ConvertFIXDate get golang time_Time data type from FIX date data type.
func ConvertFIXDate(yyyymmdd string) time.Time {
	t0, _ := time.Parse("20060102T15:04:05Z07:00", yyyymmdd+"T12:00:00Z")
	return t0
}
