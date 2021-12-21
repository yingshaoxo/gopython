package mapTool

func CheckIfAKeyIsInTheMap(key string, mapData map[string]interface{}) bool {
	_, ok := mapData[key]
	return ok
}

func GetMapValueByKey(key string, mapData map[string]interface{}) interface{} {
	value, ok := mapData[key]
	if ok {
		return value
	}
	return nil
}
