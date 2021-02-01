package main

import (
	"fmt"
	"time"

	"github.com/hpcloud/tail"
)

func main() {
	fileName := "./my.log"
	config := tail.Config{ //配置项
		ReOpen:    true,                                 //是否重新打开
		Follow:    true,                                 //是否跟踪文件
		Location:  &tail.SeekInfo{Offset: 0, Whence: 2}, //从文件哪个位置开始读取
		MustExist: false,                                //文件不存在是否报错
		Poll:      true,                                 //
	}
	tails, err := tail.TailFile(fileName, config)
	if err != nil {
		fmt.Println("tail file faild,err:", err)
		return
	}
	var (
		msg *tail.Line
		ok  bool
	)
	for {
		msg, ok = <-tails.Lines
		if !ok {
			fmt.Println("tail file close reopen ,filename:", tails.Filename)
			time.Sleep(time.Second)
			continue
		}
		fmt.Println("msg:", msg.Text)
	}
}
