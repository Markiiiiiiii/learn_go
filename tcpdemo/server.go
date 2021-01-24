package main

import (
	"fmt"
	"net"
	"sync"
)

var wg sync.WaitGroup

func processConn(conn net.Conn) {
	defer wg.Done()
	var tmp [128]byte
	n, err := conn.Read(tmp[:])
	if err != nil {
		fmt.Println("don't read data,err", err)
		return
	}
	fmt.Println(string(tmp[:n]))
}
func main() {

	lintener, err := net.Listen("tcp", "127.0.0.1:20000")
	if err != nil {
		fmt.Println("don't listen port,err:", err)
		return
	}
	for {
		conn, err := lintener.Accept()
		if err != nil {
			fmt.Println("don't acccpt data,err:", err)
			return
		}
		wg.Add(1)
		go processConn(conn)
	}
	wg.Wait()
}
