package main

import (
	"fmt"
	"learn_golang/loganget/etcd"
	"learn_golang/loganget/kafka"
	"learn_golang/loganget/taillog"
	"strconv"
	"sync"
	"time"

	"github.com/unknwon/goconfig"
)

var CfgInfo map[string]string

// func run() {
// 	// 1.读日志
// 	for {
// 		select {
// 		case line := <-tail.ReadChan():
// 			// 2.发送到kafka
// 			kafka.SendToKafka(CfgInfo["topic"], line.Text)
// 		default:
// 			time.Sleep(time.Second)
// 		}

// 	}

// }

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
	maxSize, _ := strconv.ParseInt(CfgInfo["chan_max_size"], 10, 0)
	err = kafka.Init([]string{CfgInfo["address"]}, int(maxSize))
	if err != nil {
		fmt.Println("goconfigt kafka faild ,err:", err)
		return
	}

	// 2.初始化etcd
	timeout, _ := strconv.ParseInt(CfgInfo["timeout"], 10, 0)

	err = etcd.Init(CfgInfo["etcdaddrees"], time.Duration(timeout)*time.Second)
	if err != nil {
		fmt.Println("etcd conntent fiald,err:", err)
		return
	}

	// 2.1从etcd中拉取日志收集项的配置信息
	logEntryConf, err := etcd.GetInfo(CfgInfo["collect_log_key"])
	if err != nil {
		fmt.Println("etcd.GetInfo failed ,err:", err)
		return
	}
	// 派一个哨兵监视日志收集项的变化（有变化及时通知我的logagent实现热加载配置）

	// for index, value := range logEntryConf {
	// 	fmt.Println(index, value)
	// }
	// 2.2派一个哨兵监视日志收集项的变化（有变化及时通知我的logagent实现热加载配置）
	//3.收集日志发往kafka
	//3.1循环每一个日志收集项，创建一个tailobj
	// for _, logEntry := range logEntryConf {
	// 	// logEntry.Path 要收集的日志路径
	// 	taillog.NewTailTask(logEntry.Path, logEntry.Topic)
	// 	// 3.2发往kafka
	taillog.Init(logEntryConf)
	//NewConfChan()访问了tskmgr的newconfchan这个channal是在init初始化的时候执行的
	newConfChan := taillog.NewConfChan() //从taillog包中获取对外暴露的通道
	var wg sync.WaitGroup
	wg.Add(1)
	go etcd.WatchConf(CfgInfo["collect_log_key"], newConfChan) //哨兵发现最新的配置信息，通知上面的通道
	wg.Wait()
	// }

	// //3.初始化日志收集模块开始收集日志
	// err = tail.Init(CfgInfo["path"])
	// if err != nil {
	// 	fmt.Println("Init taillog faild,err:", err)
	// 	return
	// }
	// run()
}
