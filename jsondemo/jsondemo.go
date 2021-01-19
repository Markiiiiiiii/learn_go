package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	str := `{"name":"张三","age":17}`
	var p person
	err := json.Unmarshal([]byte(str), &p)
	if err != nil {
		fmt.Printf("err:%v", err)
		return
	}
	fmt.Println(p)
}
