package main

import (
	"mylogger"
	"time"
)

func main() {
	log := mylogger.Newlog("debug")
	name := "李四"
	id := 1000
	for {
		log.Debug("这是一条debug日志,id:%d,name:%s", id, name)
		log.Info("这是一条Info日志")
		log.Fatal("这是一条Fatal日志")
		log.Error("这是一条Error日志")
		log.Waring("这是一条Waring日志")
		time.Sleep(time.Second * 3)
	}
}
