package timeTool

import "time"

func IsTheFirstTimestampBeforeTheSecond(timestamp1 int64, timestamp2 int64) bool {
	return timestamp1 < timestamp2
}

func GetCurrentTimestamp() int64 {
	return time.Now().Unix()
}
