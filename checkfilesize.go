package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("./test.log")
	if err != nil {

		return
	} //获取文件对象的类型
	fmt.Printf("%T\n", file)
	fileinfo, err := file.Stat()
	if err != nil {

		return
	}
	// 获取文件对象的详细信息
	fmt.Println(fileinfo.Size())
}
