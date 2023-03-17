// We handle dictionary/map/key-value related problems here
package dict_tool

func Check_if_a_key_is_in_the_dict(a_dict map[string]any, key string) bool {
	_, ok := a_dict[key]
	return ok
}

func Get_dict_value_by_giving_a_key(a_dict map[string]any, key string) any {
	value, ok := a_dict[key]
	if ok {
		return value
	}
	return nil
}
