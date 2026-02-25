package util

import "time"

// DateTime returns the formats time.Now parsed, accordingly to time.DateTime constant.
func DateTime() time.Time {
	parsedTime, _ := time.Parse(time.DateTime, time.Now().Format(time.DateTime))
	return parsedTime
}
