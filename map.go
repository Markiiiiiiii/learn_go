package main

import "fmt"

func main() {
	// var m1 map[string]int
	m1 := make(map[string]int, 10)
	m1["aaa"] = 10
	m1["bbb"] = 200
	// fmt.Println(m1)
	// value, ok := m1["ccc"]
	// if ok {
	// 	fmt.Println(value)
	// } else {
	// 	fmt.Println("not key")
	// }
	// for k, v := range m1 {
	// 	fmt.Println(k, v)
	// }
	delete(m1, "aaa")
	fmt.Println(m1)
}
