package quickfixdev

import (
	"time"
)

func ConvertFIXDate(yyyymmdd string) time.Time {
	t0, _ := time.Parse("20060102T15:04:05Z07:00", yyyymmdd+"T12:00:00Z")
	return t0
}
