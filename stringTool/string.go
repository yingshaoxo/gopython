package stringTool

import (
	"strconv"
	"strings"
)

func IntToString(number int) string {
	return strconv.Itoa(number)
}

func Int64ToString(number int64) string {
	return strconv.FormatInt(number, 10)
}

func StringToInt64(numberString string) (int64, error) {
	return strconv.ParseInt(numberString, 10, 64)
}

func Float64ToInt64(number float64) int64 {
	return int64(number)
}

func Float64ToInt(number float64) int {
	return int(number)
}

func BytesToString(data []byte) string {
	return string(data)
}

func HasSubstring(data string, subString string) bool {
	return strings.Contains(data, subString)
}
