package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	x      = 0
	wg     sync.WaitGroup
	rwlock sync.RWMutex
)

func write() {
	defer wg.Done()
	rwlock.Lock()
	x = x + 1
	time.Sleep(time.Millisecond * 5)
	rwlock.Unlock()
}
func read() {
	defer wg.Done()
	rwlock.RLock()
	fmt.Println(x)
	time.Sleep(time.Millisecond)
	rwlock.RUnlock()
}

func main() {
	startTime := time.Now()
	for i := 0; i < 100; i++ {
		go write()
		wg.Add(1)
	}
	time.Sleep(time.Second)
	for i := 0; i < 10000; i++ {
		go read()
		wg.Add(1)
	}
	fmt.Println(time.Now().Sub(startTime))
}
