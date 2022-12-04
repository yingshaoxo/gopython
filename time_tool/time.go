package time_tool

import "time"

func Get_current_time_as_timestamp() int64 {
	return time.Now().Unix()
}

func Is_the_first_time_happend_before_the_second_one(timestamp1 int64, timestamp2 int64) bool {
	return timestamp1 < timestamp2
}
