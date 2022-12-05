package string_tool

import (
	"encoding/json"

	"github.com/yingshaoxo/gopython/string_tool"
)

func Convert_struct_object_to_map(an_object interface{}) (map[string]interface{}, error) {
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

func Convert_map_to_json_string(a_dict map[string]interface{}) string {
	json_bytes, err := json.Marshal(a_dict)
	if err != nil {
		return ""
	}
	return string_tool.Bytes_to_string(json_bytes)
}

func Convert_struct_object_to_json_string(an_object interface{}) string {
	json_bytes, err := json.Marshal(an_object)
	if err != nil {
		return ""
	}

	return string_tool.Bytes_to_string(json_bytes)
}

func Convert_bytes_json_data_to_map(data []byte) (map[string]interface{}, error) {
	var result map[string]interface{}
	err := json.Unmarshal(data, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}
