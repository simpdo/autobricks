package util

//GetStrValue 指定key的字符串值
func GetStrValue(m map[string]interface{}, key string) (string, bool) {
	_, ok := m[key]
	if !ok {
		return "", false
	}

	val, ok := m[key].(string)
	if !ok {
		return "", false
	}

	return val, true
}
