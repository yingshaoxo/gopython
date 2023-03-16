// We handle dictionary/map/key-value related problems here
package map_tool

import "github.com/yingshaoxo/gopython/dict_tool"

func Check_if_a_key_is_in_the_map(key string, a_dict map[string]any) bool {
	return dict_tool.Check_if_a_key_is_in_the_dict(key, a_dict)
}

func Get_map_value_by_giving_a_key(key string, a_dict map[string]any) any {
	return dict_tool.Get_dict_value_by_giving_a_key(key, a_dict)
}
