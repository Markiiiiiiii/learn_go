package main

import (
	"fmt"
	"net"
	"strings"
)

//GetOutBoundIP 获取本机对外IP地址
func GetOutBoundIp() (ip string, err error) {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		fmt.Println("don't Get ip ,err:", err)
		return
	}
	defer conn.Close()
	localaddr := conn.LocalAddr().(*net.UDPAddr)
	fmt.Println(localaddr.String())
	ip = strings.Split(localaddr.IP.String(), ":")[0]
	return
}
func main() {
	ip, err := GetOutBoundIp()
	if err != nil {
		fmt.Println("err:", err)
		return
	}
	fmt.Println(ip)
}
