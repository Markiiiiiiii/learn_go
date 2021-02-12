package main

import (
	"fmt"

	"learn_go.com/logtransfer/es"
	"learn_go.com/logtransfer/kafka"

	"gopkg.in/ini.v1"
	"learn_go.com/logtransfer/conf"
)

func main() {
	//0.加载配置文件
	var cfg conf.LogTransferCfg
	err := ini.MapTo(&cfg, "./conf/cfg.ini") //在一个函数中修改变量一定要传入指针
	if err != nil {
		fmt.Println("init config faild,err:", err)
		return
	}
	// 1.1初始化
	//1.1初始化 es连接的clinet
	//1.1.1对外提供一个往ES写入数据的一个函数
	err = es.Init(cfg.ESCfg.Address, cfg.ESCfg.ChanSize, cfg.ESCfg.Nums)
	if err != nil {
		fmt.Println("[ES] package init faild ,err:", err)
		return
	}
	//1.2.1链接kafka，创建分区的消费者
	//1.2.2每个分区的消费者分别取得数据 通过SendToEs发给ES
	err = kafka.Init([]string{cfg.KafkaCfg.Address}, cfg.KafkaCfg.Topic)
	if err != nil {
		fmt.Println("[kafka] package init faild, err:", err)
		return
	}

	// 2.从kafka取日志数据
	// 3.发往ES
	select {}
}
