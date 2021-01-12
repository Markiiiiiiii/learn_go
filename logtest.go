package main

import (
	"fmt"
	"os"
)

//logger.Debug
//logger.Waring
// logger.Info
// logger.Trace
// logger.Error("日志的内容")
// 需求：
// 1.可以往不同的输出位置记录日志
// 2.日志分为5种级别
func main() {
	fmt.Fprint(os.Stdout, "这是一条日志")
	fileObj, err := os.OpenFile("./test.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0777)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Fprintln(fileObj, "这是一条日志")
}
