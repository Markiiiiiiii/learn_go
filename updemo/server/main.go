package main

import (
	"fmt"
	"net"
	"strings"
)

//udp
func main() {
	conn, err := net.ListenUDP("udp", &net.UDPAddr{
		IP:   net.IPv4(127, 0, 0, 1),
		Port: 40000,
	})
	//不能在这里写defer conn.close(),因为没有判断是否出现错误，如果出现错误在此关闭会出错
	if err != nil {
		fmt.Println("udp faild,err:", err)
		return
	}
	defer conn.Close()
	var data [1024]byte
	for {
		n, addr, err := conn.ReadFromUDP(data[:])
		if err != nil {
			fmt.Println("read faild,err:", err)
			return
		}
		fmt.Println(data[:n])
		reply := strings.ToUpper(string(data[:n]))
		conn.WriteToUDP([]byte(reply), addr)
	}
}
