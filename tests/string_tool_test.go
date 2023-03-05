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
