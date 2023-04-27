// We do json operations here
package json_tool

import (
	"encoding/json"

	"github.com/yingshaoxo/gopython/string_tool"
	"github.com/yingshaoxo/gopython/variable_tool"
)

func Convert_dirty_map_into_pure_map(an_object_instance any) any {
	if an_object_instance == nil {
		return nil
	}

	if variable_tool.Is_the_variable_a_struct_object(an_object_instance) {
		a_map, _ := Convert_struct_object_to_map(an_object_instance)
		return Convert_dirty_map_into_pure_map(a_map)
	}

	if variable_tool.Is_the_variable_a_list_object(an_object_instance) {
		var new_list = make([]any, 0)
		switch t := an_object_instance.(type) {
		case []any:
			for _, value := range t {
				new_list = append(new_list, Convert_dirty_map_into_pure_map(value))
			}
		}
		return new_list
	}

	if variable_tool.Is_the_variable_a_dict_object(an_object_instance) {
		var new_dict = make(map[string]any)
		for key, value := range an_object_instance.(map[string]any) {
			new_dict[key] = Convert_dirty_map_into_pure_map(value)
		}
		return new_dict
	}

	return an_object_instance
}

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

func Convert_map_to_struct_object(a_dict any, the_memory_reference_to_that_object any) {
	var json_string = Convert_map_to_json_string(a_dict)
	json.Unmarshal([]byte(json_string), the_memory_reference_to_that_object)
}

func Convert_json_string_to_struct_object(json_string string, the_memory_reference_to_that_object any) {
	json.Unmarshal([]byte(json_string), the_memory_reference_to_that_object)
}
