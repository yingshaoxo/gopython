package main

import (
	"fmt"
	"testing"

	variable_tool "github.com/yingshaoxo/gopython/variable_tool"
)

func Test_nullable(t *testing.T) {
	hi := "hi"
	ok := "ok"

	a_string := variable_tool.Nullable[*string]{Value: &hi, Is_null: true}
	fmt.Println(*a_string.Value)
	fmt.Println(a_string.Is_null)

	a_string.Value = &ok
	a_string.Is_null = false
	fmt.Println(*a_string.Value)
	fmt.Println(a_string.Is_null)
}
