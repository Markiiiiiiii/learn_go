package main

import (
	"fmt"
	"path"
	"runtime"
)

func getInfo() {
	pc, file, line, ok := runtime.Caller(0)
	if !ok {
		fmt.Println("runtime.caller() faild ")
		return
	}
	fmt.Println(runtime.FuncForPC(pc).Name())
	fmt.Println(file)            //定位文件的绝对路径
	fmt.Println(path.Base(file)) //只获取文件名
	fmt.Println(line)            //定位行号

}
func main() {
	getInfo()
}
