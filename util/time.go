package util

import (
	"strconv"
	"time"
)

// TimestampToTime 时间戳字符串转时间
func TimestampToTime(timeStr string) time.Time {
	timestamp, _ := strconv.ParseInt(timeStr, 10, 64) // 字符串转int64
	return time.Unix(timestamp, 0)
}

// TimeFormat 输出标准格式时间字符串
func TimeFormat(tm time.Time) string {
	return tm.Format("2006-01-02 15:04:05")
}
