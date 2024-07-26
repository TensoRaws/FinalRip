package util

import (
	"github.com/bytedance/sonic"
)

// StructToString 结构体转字符串
func StructToString(s interface{}) string {
	// v, _ := json.Marshal(s)
	v, _ := sonic.Marshal(s)
	return string(v)
}

// StringToStruct 字符串转结构体
func StringToStruct(str string, s interface{}) error {
	// return json.Unmarshal([]byte(str), s)
	return sonic.Unmarshal([]byte(str), s)
}

// StructToMap 结构体转 map[string]interface{}，请勿在有数字情况下使用，请使用反射
func StructToMap(s interface{}) (map[string]interface{}, error) {
	// 使用 sonic 将结构体序列化为 JSON
	jsonBytes, err := sonic.Marshal(s)
	if err != nil {
		return nil, err
	}

	// 将 JSON 反序列化为 map[string]interface{}
	var result map[string]interface{}
	err = sonic.Unmarshal(jsonBytes, &result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

// CheckStringInSlice 检查字符串是否在切片中
func CheckStringInSlice(item string, slice []string) bool {
	for _, s := range slice {
		if s == item {
			return true
		}
	}
	return false
}
