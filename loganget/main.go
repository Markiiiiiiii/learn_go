package main

import (
	"fmt"
	"learn_golang/loganget/kafka"
	tail "learn_golang/loganget/taillog"
	"time"

	ini "gopkg.in/ini.vi"
)

func run() {
	// 1.读日志
	for {
		select {
		case line := <-tail.ReadChan():
			// 2.发送到kafka
			kafka.SendToKafka("web_log", line.Text)
		default:
			time.Sleep(time.Second)
		}

	}

}

func main() {
	//0.加载配置文件
	cfg, err := ini.Load("./conf/config.ini")

	//1. 初始化kafka连接
	err = kafka.Init([]string{"192.168.1.102:9092"})
	if err != nil {
		fmt.Println("init kafka faild ,err:", err)
		return
	}
	//2.初始化日志收集模块开始收集日志
	err = tail.Init("./my.log")
	if err != nil {
		fmt.Println("Init taillog faild,err:", err)
		return
	}
	run()
}
