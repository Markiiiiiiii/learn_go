package main

import (
	"fmt"
	"sync"
)

var a chan int
var wg sync.WaitGroup

func noBuffChannel() {
	fmt.Println(a) //未初始化时是nil
	a = make(chan int)
	wg.Add(1)
	go func() {
		defer wg.Done()
		x := <-a
		fmt.Println("接受到通道a中的值", x)
	}()
	a <- 10
	fmt.Println("10已发送到无缓冲区的channel中了")
	fmt.Println(a) //输出的是一个内存地址
	wg.Wait()
}
func buffChannel() {
	fmt.Println(a)
	a := make(chan int, 16)
	a <- 10
	fmt.Println("10发送到缓冲区channel中了")
	x := <-a
	fmt.Println(x)
}

func main() {
	buffChannel()
}
