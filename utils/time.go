package utils

import (
	"time"
)

const datetimeNow = "2006-01-02 15:03:04"

var (
	Microsecond = 1
	Millisecond = 1000 * Microsecond
	Second      = 1000 * Millisecond
	Minute      = 60 * Second
	Hour        = 60 * Minute
	day         = 24 * Hour
)

// NowTime 获取当前时间
func NowTime() time.Time {
	return time.Now()
}
