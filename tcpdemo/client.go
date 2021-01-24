package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	//建立连接
	conn, err := net.Dial("tcp", "127.0.0.1:20000")
	if err != nil {
		fmt.Println("dial faild,err:", err)
		return
	}
	//发送数据通讯
	// var msg string
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Println("请输入：")
		text, _ := reader.ReadString('\n')
		text = strings.TrimSpace(text)
		// fmt.Scanln(&msg)
		if text == "exit" {
			break
		}
		conn.Write([]byte(text))
	}
	conn.Close()
}
