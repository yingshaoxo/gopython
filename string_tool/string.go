package string_tool

import (
	"math/big"
	"strconv"
	"strings"
)

func Int_to_string(number int) string {
	return strconv.Itoa(number)
}

func Int32_to_string(number int32) string {
	return strconv.FormatInt(int64(number), 10)
}

func Int64_to_string(number int64) string {
	return strconv.FormatInt(number, 10)
}

func String_to_int(number string, default_value_when_error_happen int) int {
	value, err := strconv.Atoi(number)
	if err != nil {
		return default_value_when_error_happen
	}
	return value
}

func String_to_int32(number_string string) (int32, error) {
	result, err := strconv.ParseInt(number_string, 10, 64)
	if err != nil {
		return 0, err
	}
	return int32(result), nil
}

func String_to_int64(number_string string) (int64, error) {
	return strconv.ParseInt(number_string, 10, 64)
}

func Float64_to_int(number float64) int {
	return int(number)
}

func Float64_to_int32(number float64) int32 {
	return int32(number)
}

func Float64_to_int64(number float64) int64 {
	return int64(number)
}

func Int_to_float64(number int) float64 {
	return float64(number)
}

func Int_to_float32(number int) float32 {
	return float32(number)
}

func Float64_to_string(number float64, numbers_after_decimal int) string {
	return big.NewFloat(number).Text('f', numbers_after_decimal)
	// numbers_after_decimal_string := Int_to_string(numbers_after_decimal)
	// return fmt.Sprintf(
	// 	"%."+numbers_after_decimal_string+"f",
	// 	number,
	// )
}

func Float32_to_string(number float32, numbers_after_decimal int) string {
	return big.NewFloat(float64(number)).Text('f', numbers_after_decimal)
}

func String_to_float64(number string, default_value_when_error_happen float64) float64 {
	result, err := strconv.ParseFloat(strings.TrimSpace(number), 64)
	if err != nil {
		return default_value_when_error_happen
	}
	return result
}

func String_to_float32(number string) (float32, error) {
	result, err := strconv.ParseFloat(strings.TrimSpace(number), 64)
	if err != nil {
		return 0, err
	}
	return float32(result), nil
}

func Bytes_to_string(data []byte) string {
	return string(data)
}

func String_to_bytes(data string) []byte {
	return []byte(data)
}

func Has_substring(data string, subString string) bool {
	return strings.Contains(data, subString)
}
