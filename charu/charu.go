package main

import (
	"fmt"
	"io"
	"os"
)

func chaRu() {
	//读取源文件
	yuanFile, err := os.OpenFile("./test.txt", os.O_RDWR, 0777)
	if err != nil {
		fmt.Printf("File open faild %v\n", err)
		return
	}
	//创建临时文件
	tmpFile, err := os.Create("./temp.tmp")
	if err != nil {
		fmt.Printf("Create file faild %v\n", err)
		return
	}

	//截取插入位置之前的数据
	strArray := [6]byte{}
	if _, err := yuanFile.Read(strArray[:]); err != nil { //读取文件返回值为读取了几个字节
		// if err1 != nil {
		fmt.Printf("The file don't read ,error:%v", err)
		return
	}
	// 写入临时文件
	tmpFile.Write(strArray[:])
	//写入插入内容
	tmpFile.WriteString("\n李瞎子")
	//移动光标到指定位置
	yuanFile.Seek(6, 0)
	// 读取剩余源文件内容
	lastStrarray := [1024]byte{}
	for {
		n, err := yuanFile.Read(lastStrarray[:])
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println("read faild")
		}
		tmpFile.Write(lastStrarray[:n])
	}
	// 写入剩余内容
	tmpFile.Close()
	yuanFile.Close()
	// 重命名临时文件名为源文件名
	os.Rename("./temp.tmp", "./test.txt")
}
func main() {
	chaRu()
}
