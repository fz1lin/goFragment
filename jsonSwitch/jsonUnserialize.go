package jsonSwitch

import (
	"encoding/json"
	"fmt"
	"log"
)

type Addresss struct {
	Street string `json:"street"`
	City   string `json:"city"`
	State  string `json:"state"`
}

type Persons struct {
	Name    string  `json:"name"`
	Age     int     `json:"age"`
	Address Address `json:"address"`
}

func main() {
	jsonString := `{"name": "Alice", "age": 25, "address": {"street": "123 Main St", "city": "Anytown", "state": "CA"}}`
	var persons Persons
	err := json.Unmarshal([]byte(jsonString), &persons)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%+v\n", persons)
}
