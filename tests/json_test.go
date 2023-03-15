package main

import (
	"fmt"
	"testing"

	"github.com/yingshaoxo/gopython/json_tool"
)

// func Convert_map_to_struct_object(a_dict map[string]interface{}, the_memory_reference_to_that_object *any) {
// 	var json_string = Convert_map_to_json_string(a_dict)
// 	json.Unmarshal([]byte(json_string), the_memory_reference_to_that_object)
// }

// func Convert_json_string_to_struct_object(json_string string, the_memory_reference_to_that_object *any) {
// 	json.Unmarshal([]byte(json_string), the_memory_reference_to_that_object)
// }

type User struct {
	Name string
}

func Test_json_tool(t *testing.T) {
	data := make(map[string]interface{})
	data["name"] = "yingshaoxo"

	var json_string = json_tool.Convert_map_to_json_string(data)

	var user User
	json_tool.Convert_json_string_to_struct_object(json_string, &user)
	// json_tool.Convert_map_to_struct_object(data, &user)

	fmt.Println(json_tool.Convert_struct_object_to_json_string(user))
}
