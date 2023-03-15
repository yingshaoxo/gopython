package main

import (
	"fmt"
	"testing"

	variable_tool "github.com/yingshaoxo/gopython/variable_tool"
)

func Test_nullable(t *testing.T) {
	hi := "hi"
	ok := "ok"

	a_string := variable_tool.Nullable(&hi)
	a_string.Is_null = true
	fmt.Println(*a_string.Value)
	fmt.Println(a_string.Is_null)

	a_string.Value = &ok
	a_string.Is_null = false
	fmt.Println(*a_string.Value)
	fmt.Println(a_string.Is_null)
}

func Test_result(t *testing.T) {
	hi := "hi"

	a_string := variable_tool.Result(&hi)
	a_string.Error = "error"
	fmt.Println(a_string.Error)
	fmt.Println(*a_string.Value.Value)
}
