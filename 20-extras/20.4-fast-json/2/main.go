package main

import (
	"encoding/json"
	"fmt"

	"github.com/valyala/fastjson"
)

type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	var p fastjson.Parser
	jsonData := `{"user": {"name": "John Doe", "age": 30}}`

	value, err := p.Parse(jsonData)
	if err != nil {
		panic(err)
	}
	userJSON := value.Get("user").String()
	var user User
	if err := json.Unmarshal([]byte(userJSON), &user); err != nil {
		panic(err)
	}
	fmt.Printf("User name: %s\n", user.Name)
	fmt.Printf("User age: %d\n", user.Age)
}
