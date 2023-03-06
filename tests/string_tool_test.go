package main

import (
	"log"
	"testing"

	"github.com/yingshaoxo/gopython/string_tool"
)

func Test_string_ralated_convertion(t *testing.T) {
	number := 3.14556
	result := string_tool.Float64_to_string(number, 3)
	log.Println(result)
}

func Test_bytes_ralated_convertion(t *testing.T) {
	byte_ := []byte{52, 46, 53}

	result1 := string_tool.StringBytes_to_float64(byte_, 0)
	result2 := string_tool.Float64_to_StringBytes(result1, 5)
	result3 := string_tool.StringBytes_to_float64(result2, 0)

	log.Println(result1)
	log.Println(result2)
	log.Println(result3)

	if result1 == result3 {
		log.Println("ok")
	}
}
