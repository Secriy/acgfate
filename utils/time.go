package utils

import (
	"strconv"
	"time"
)

// TimestampToTime 字符串转时间
func TimestampToTime(timeStr string) time.Time {
	timestamp, _ := strconv.ParseInt(timeStr, 10, 64) // 字符串转int64
	return time.Unix(timestamp, 0)
}
