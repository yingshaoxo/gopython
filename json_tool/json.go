// We do json operations here
package json_tool

import (
	"encoding/json"

	"github.com/yingshaoxo/gopython/string_tool"
)

func Convert_struct_object_to_map(an_object any) (map[string]interface{}, error) {
	var a_dict map[string]interface{}

	json_bytes, err := json.Marshal(an_object)
	if err != nil {
		return a_dict, err
	}

	err = json.Unmarshal(json_bytes, &a_dict)
	if err != nil {
		return a_dict, err
	}

	return a_dict, nil
}

func Convert_json_string_to_map(json_string string) (map[string]interface{}, error) {
	var result map[string]interface{}
	err := json.Unmarshal([]byte(json_string), &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func Convert_bytes_json_data_to_map(data []byte) (map[string]interface{}, error) {
	var result map[string]interface{}
	err := json.Unmarshal(data, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func Convert_map_to_json_string(a_dict any) string {
	json_bytes, err := json.Marshal(a_dict)
	if err != nil {
		return ""
	}
	return string_tool.Bytes_to_string(json_bytes)
}

func Convert_struct_object_to_json_string(an_object any) string {
	json_bytes, err := json.Marshal(an_object)
	if err != nil {
		return ""
	}

	return string_tool.Bytes_to_string(json_bytes)
}

func Convert_map_to_struct_object[T any](a_dict any, the_memory_reference_to_that_object *T) {
	var json_string = Convert_map_to_json_string(a_dict)
	json.Unmarshal([]byte(json_string), the_memory_reference_to_that_object)
}

func Convert_json_string_to_struct_object[T any](json_string string, the_memory_reference_to_that_object *T) {
	json.Unmarshal([]byte(json_string), the_memory_reference_to_that_object)
}
