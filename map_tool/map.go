// We handle dictionary/map/key-value related problems here
package map_tool

func Check_if_a_key_is_in_the_map(key string, mapData map[string]interface{}) bool {
	_, ok := mapData[key]
	return ok
}

func Get_map_value_by_giving_a_key(key string, mapData map[string]interface{}) interface{} {
	value, ok := mapData[key]
	if ok {
		return value
	}
	return nil
}
