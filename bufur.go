package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func useBufiopenfile() {
	file, err := os.Open("./bot.go")
	if err != nil {
		fmt.Println("err")
		return
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	for {
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			return
		}
		if err != nil {
			fmt.Println("errr")
			return
		}
		fmt.Println(line)
	}
}
func useIoutilopenfile() {
	ret, err := ioutil.ReadFile("./bot.go")
	if err != nil {
		return
	}
	fmt.Println(string(ret))

}
func main() {
	// useBufiopenfile()
	useIoutilopenfile()
}
