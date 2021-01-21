package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func hello(i int) {
	fmt.Println("hello", i)
}
func f() {
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 10; i++ {
		r1 := rand.Int()
		r2 := rand.Intn(10)
		fmt.Println(r1, r2)
	}
}
func f1(i int) {
	time.Sleep(time.Millisecond * time.Duration(rand.Intn(300)))
	fmt.Println(i)
	wg.Done() //结束
}

var wg sync.WaitGroup

func main() {
	// for i := 0; i < 1000; i++ {
	// 	go hello(i)
	// }
	// fmt.Println("finsh main ")
	// wg.Add(10)
	for i := 0; i < 10; i++ {
		wg.Add(1) //计数器减1
		go f1(i)
	}
	wg.Wait() //等待wg的计数器减为0
}
