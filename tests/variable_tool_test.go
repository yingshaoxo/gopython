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

func Test_type_string_getting(t *testing.T) {
	hi := "hi"

	var type_string = variable_tool.Get_variable_type_string_representation(hi)
	fmt.Println(type_string)
}

type TestClass struct {
	Name string
}

func Test_get_dict_from_class_object(t *testing.T) {
	var test_class = TestClass{Name: "yingshaoxo"}

	fmt.Println(variable_tool.Get_key_value_dict_from_struct_object(test_class, true))
	fmt.Println(variable_tool.Get_key_value_dict_from_struct_object(test_class, false))
}

func Test_is_null_and_set_to_null(t *testing.T) {
	var hi = variable_tool.Nullable("h").Set_to_null()

	fmt.Println(hi)
	fmt.Println(variable_tool.Is_it_null(hi))
}
