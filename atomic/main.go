package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

//原子操作
var x int64
var wg sync.WaitGroup
var lock sync.Mutex

func add() {
	// lock.Lock()
	// x = x + 1
	// lock.Unlock()
	atomic.AddInt64(&x, 1)
	wg.Done()
}
func main() {
	for i := 0; i < 10000; i++ {
		wg.Add(1)
		go add()

	}

	wg.Wait()
	fmt.Println(x)
}
