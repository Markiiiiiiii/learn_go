package main

import "fmt"

func main() {
	// n := 10
	// fmt.Println(&n)
	// t := &n
	// fmt.Println(&t)
	// fmt.Printf("%T\n", t)
	// m := *t
	// fmt.Println(m)
	// fmt.Printf("%T\n", m)

	// var a *int
	var a = new(int)
	fmt.Println(a)
	*a = 100
	fmt.Println(a)
	fmt.Println(*a)

}
