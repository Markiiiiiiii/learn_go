package main

import (
	"flag"
	"fmt"
	"time"
)

func main() {

	name := flag.String("name", "李四", "请输入名字")
	cTime := flag.Duration("time", time.Second, "time")

	// var name string
	// flag.StringVar(&name, "name", "例子", "enter name")
	flag.Parse()
	// fmt.Println(name)
	fmt.Println(*name)
	fmt.Println(*cTime)
	fmt.Println(flag.Args())  //返回命令行参数后的其他参数，【】string类型
	fmt.Println(flag.NArg())  //返回命令行后其他参数
	fmt.Println(flag.NFlag()) //有几个标志位
}
