package util

import "strconv"

func ToString(i interface{}) string {
	switch i.(type) {
	case string:
		return i.(string)
	case int64:
		return strconv.FormatInt(i.(int64), 10)
	default:
		return ""
	}
}
