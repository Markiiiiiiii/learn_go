package main

import "fmt"

// const (
// 	a = iota
// 	b
// 	c
// 	d

// )
const (
	_          = iota
	KB float64 = 1 << (10 * iota)
	MB
)

func main() {
	// fmt.Println(a, b, c, d)
	fmt.Println(KB, MB)
}
