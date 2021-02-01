package main

import (
	"fmt"
	"learn_golang/loganget/kafka"
	tail "learn_golang/loganget/taillog"
	"time"

	"github.com/unknwon/goconfig"
)

var CfgInfo map[string]string

func run() {
	// 1.读日志
	for {
		select {
		case line := <-tail.ReadChan():
			// 2.发送到kafka
			kafka.SendToKafka(CfgInfo["topic"], line.Text)
		default:
			time.Sleep(time.Second)
		}

	}

}

func main() {
	CfgInfo = make(map[string]string, 10)
	//0.加载配置文件
	cfg, err := goconfig.LoadConfigFile("./conf/config.ini")
	if err != nil {
		fmt.Println("load goconfig file false. err:", err)
		return
	}
	for _, v := range cfg.GetSectionList() { //获取所有分区列表
		info, err := cfg.GetSection(v)
		if err != nil {
			fmt.Println("get section faild,err:", err)
			return
		}
		for k, x := range info {
			CfgInfo[k] = x
		}
	}
	// fmt.Println(CfgInfo)
	//1. 初始化kafka连接
	err = kafka.Init([]string{CfgInfo["address"]})
	if err != nil {
		fmt.Println("goconfigt kafka faild ,err:", err)
		return
	}
	//2.初始化日志收集模块开始收集日志
	err = tail.Init(CfgInfo["path"])
	if err != nil {
		fmt.Println("Init taillog faild,err:", err)
		return
	}
	run()
}
