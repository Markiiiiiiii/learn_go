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

var jobChan = make(chan *job, 100)       //声明一个结构体类的通道用于接受数据后再发送数据
var resultChan = make(chan *result, 100) //声明一个结构体类型的通道，用于接受数据

func numMaker(x chan<- *job) { //传入一个Job类型的结构体通道，这个通道用于接受值
	defer wg.Done()
	for {
		y := rand.Int63() //生成一个int64的随机数
		newJob := &job{
			value: y, //将随机数保存入结构体中
		}
		x <- newJob //每次都把结构体传入到通道中
		time.Sleep(time.Millisecond * 500)
	}
}
func goWorker(x <-chan *job, z chan<- *result) { //传入一个job类型的结构体通道，用于发送数据，传入一个result类型的结构体用于接收数据
	defer wg.Done()
	for {
		number := <-x     //把job通道中的每个job结构体都另存一个变量，为了避免出现修改的情况
		sum := int64(0)   //把sum强制传化为一个int64类型
		n := number.value //取出结构体中的值
		for n > 0 {
			sum += n % 10 //模运算取余后相加
			n = n / 10
		}
		newReslut := &result{ //将相加的值存入到result结构体中
			job:    number,
			result: sum,
		}
		z <- newReslut //将result结构体传入到通道中
	}
}
func main() {
	wg.Add(1)
	go numMaker(jobChan) //开一个并发goruntine生成随机数
	wg.Add(24)
	for i := 0; i < 24; i++ {
		go goWorker(jobChan, resultChan) //开24个并发goruntine对随机数的各位数相加
	}
	for result := range resultChan {
		fmt.Printf("value:%d sum:%d\n", result.job.value, result.result)
	}
	wg.Wait()
}
