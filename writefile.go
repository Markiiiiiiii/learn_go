package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

func openFilewritmode() {
	fileer, err := os.OpenFile("./test.txt", os.O_RDWR|os.O_APPEND|os.O_CREATE, 0777)
	if err != nil {

		return
	}
	defer fileer.Close()
	fileer.Write([]byte("s"))
	fileer.WriteString("张三李四")
	fileer.Close()
}
func writeFilebufiomode() {
	fileer, err := os.OpenFile("./test.txt", os.O_RDWR|os.O_APPEND|os.O_CREATE, 0777)
	if err != nil {

		return
	}
	defer fileer.Close()
	wr := bufio.NewWriter(fileer)
	wr.WriteString("wihe bufio mode write to file") //写入缓存
	wr.Flush()                                      //将缓存的字符串写入文件
}
func writeFilewithioutilmode() {
	str := "/n write file with ioutil mode/n"
	err := ioutil.WriteFile("/test.txt", []byte(str), 0777)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func main() {
	// openFilewritmode()
	// writeFilebufiomode()
	writeFilewithioutilmode()
}
