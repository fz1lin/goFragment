package jsonSwitch

import (
	"encoding/json"
	"fmt"
)

//JSON转换为map

// jsonStr := `{"name": "Alice", "age": 20}`
func JsonSwitchMap(jsonStr string) map[string]interface{} {
	// 将 JSON 数据转换为 map
	var data map[string]interface{}
	err := json.Unmarshal([]byte(jsonStr), &data)
	if err != nil {
		fmt.Println(err)
	}
	return data
}

//map[age:20 name:Alice]

func MapSwitchJson(mapData map[string]interface{}) string {
	// 将 map 转换为 JSON 字符串
	jsonBytes, err := json.Marshal(mapData)
	if err != nil {
		fmt.Println(err)
	}
	// 将 JSON 字符串转换为字符串类型输出
	jsonString := string(jsonBytes)
	return jsonString
}

//{"age":30,"email":"example.com","name":"John"}
