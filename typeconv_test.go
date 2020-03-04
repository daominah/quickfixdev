package quickfixdev

import (
	"testing"
	"time"
)

func TestConvertFIXDate(t *testing.T) {
	t0 := ConvertFIXDate("20200203")
	if t0.Year() != 2020 || t0.Month() != time.February || t0.Day() != 3 {
		t.Error(t0)
	}
	t0 = ConvertFIXDate("20190229") // non exist day
	if t0.Day() == 29 {
		t.Error(t0)
	}
}
