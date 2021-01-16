package main

import (
	"mylogger"
	"time"
)

var log mylogger.LogInterface //声明一个全局接口变量

func main() {
	log = mylogger.NewConsoleLog("error") //终端打印日志模式
	// log = mylogger.NewFilelogger("info", "./", "test.log", 10*1024) //文件生成日志模式
	for {
		name := "李四"
		id := 1000
		log.Debug("这是一条debug日志,id:%d,name:%s", id, name)
		log.Info("这是一条Info日志,id:%d,name:%s", id, name)
		log.Fatal("这是一条Fatal日志,id:%d,name:%s", id, name)
		log.Error("这是一条Error日志,id:%d,name:%s", id, name)
		log.Waring("这是一条Waring日志,id:%d,name:%s", id, name)
		time.Sleep(time.Second)
	}
}
