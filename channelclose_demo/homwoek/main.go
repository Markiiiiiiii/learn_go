package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

//job ...
type job struct {
	value int64
}

//result ...
type result struct {
	job    *job
	result int64
}

var wg sync.WaitGroup
var jobChan = make(chan *job, 100)
var resultChan = make(chan *result, 100)

func numMaker(x chan<- *job) {
	defer wg.Done()
	for {
		y := rand.Int63()
		newJob := &job{
			value: y,
		}
		x <- newJob
		time.Sleep(time.Millisecond * 500)
	}
}
func goWorker(x <-chan *job, z chan<- *result) {
	defer wg.Done()
	for {
		number := <-x
		sum := int64(0)
		n := number.value
		for n > 0 {
			sum += n % 10
			n = n / 10
		}
		newReslut := &result{
			job:    number,
			result: sum,
		}
		resultChan <- newReslut
	}
}
func main() {
	wg.Add(1)
	go numMaker(jobChan)
	wg.Add(24)
	for i := 0; i < 24; i++ {
		go goWorker(jobChan, resultChan)
	}
	for result := range resultChan {
		fmt.Printf("value:%d sum:%d\n", result.job.value, result.result)
	}
	wg.Wait()
}
