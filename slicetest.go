package main

import (
	"fmt"
	"io"
	"os"
)

func main() {

	b, err := os.Open("./bot.go")
	if err != nil {
		fmt.Println("don't open this file", err)
		return
	}
	defer b.Close()
	tmp := [128]byte{}
	for {
		n, err := b.Read(tmp[:])
		if err == io.EOF {
			fmt.Println("readend")
			return
		}
		if err != nil {
			fmt.Println("file read error")
			return
		}
		if n < 128 {
			return
		}
	}
}
