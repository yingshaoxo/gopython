package string_tool

import (
	"strconv"
	"strings"
)

func Int_to_string(number int) string {
	return strconv.Itoa(number)
}

func String_to_int(number string, default_value_when_error_happen int) int {
	value, err := strconv.Atoi(number)
	if err != nil {
		return default_value_when_error_happen
	}
	return value
}

func Int64_to_string(number int64) string {
	return strconv.FormatInt(number, 10)
}

func String_to_int64(number_string string) (int64, error) {
	return strconv.ParseInt(number_string, 10, 64)
}

func Float64_to_int64(number float64) int64 {
	return int64(number)
}

func Float64_to_int(number float64) int {
	return int(number)
}

func Bytes_to_string(data []byte) string {
	return string(data)
}

func Has_substring(data string, subString string) bool {
	return strings.Contains(data, subString)
}
