package stringTool

import "strconv"

func IntToString(number int) string {
	return strconv.Itoa(number)
}

func Int64ToString(number int64) string {
	return strconv.FormatInt(number, 10)
}
