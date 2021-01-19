package main

import "fmt"

func main() {
	a := []string{"xxs", "sda", "sda"}
	// fmt.Printf("%v,%d,%d\n", a, len(a), cap(a))
	// a = append(a, "llll")
	// fmt.Printf("%v,%d,%d\n", a, len(a), cap(a))
	// b := []string{"sad", "撒", "大萨达所"}
	// a = append(a, b...)
	// fmt.Printf("%v,%d,%d\n", a, len(a), cap(a))
	b := a
	var c = make([]string, 3, 3)
	copy(c, a)
	fmt.Println(a, b, c)
	c[2] = "lol"
	fmt.Println(a, b, c)
	c = append(c[:1], c[2:]...)
	fmt.Println(c)
}
