package jsonSwitch

import (
	"encoding/json"
	"fmt"
	"log"
)

// 定义结构体，变量名注意大写，因为跨到json包了
type Address struct {
	Street string `json:"street"`
	City   string `json:"city"`
	State  string `json:"state"`
}

type Person struct {
	Name    string  `json:"name"`
	Age     int     `json:"age"`
	Address Address `json:"address"`
}

func main() {
	person := Person{
		Name: "Alice",
		Age:  25,
		Address: Address{
			Street: "123 Main St",
			City:   "Anytown",
			State:  "CA",
		},
	}
	//序列化结构体
	jsonString, err := json.Marshal(person)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(jsonString))
}

//{"name":"Alice","age":25,"address":{"street":"123 Main St","city":"Anytown","state":"CA"}}
