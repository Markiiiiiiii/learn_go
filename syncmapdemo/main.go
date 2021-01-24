package main

import (
	"fmt"
	"strconv"
	"sync"
)

func main() {
	var m2 sync.Map
	var wg sync.WaitGroup
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(n int) {
			key := strconv.Itoa(n)
			m2.Store(key, n)
			if value, ok := m2.Load(key); ok {
				fmt.Println(key, value)
			}
			wg.Done()
		}(i)
	}
	wg.Wait()
}
