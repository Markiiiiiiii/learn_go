package taillog

import (
	"fmt"
	"learn_golang/loganget/etcd"
	"time"
)

//建立一个tailobj的管理者，将所有etcd中获取到的path存储到管理者这个切片中，
//在taillog中热加载tailobj的对象
type tailLogMgr struct {
	logEntry    []*etcd.LogEntry
	tskMap      map[string]*TailTask
	newConfChan chan []*etcd.LogEntry
}

var tskMgr *tailLogMgr

func Init(logEntryConf []*etcd.LogEntry) {
	tskMgr = &tailLogMgr{
		logEntry:    logEntryConf, //把当前的日志收集项配置信息保存起来
		tskMap:      make(map[string]*TailTask, 16),
		newConfChan: make(chan []*etcd.LogEntry), //无缓冲区的通道，无数据时阻塞，有更新时才启用
	}
	for _, logEntry := range logEntryConf {
		// logEntry.Path 要收集的日志路径
		//初始化的时候起了多少个tailtask，都要记下来，后续方便做判断是否是新增
		tailObj := NewTailTask(logEntry.Path, logEntry.Topic)
		mk := fmt.Sprintf("%s_%s", logEntry.Path, logEntry.Topic)
		tskMgr.tskMap[mk] = tailObj
		// 3.2发往kafka

	}
	go tskMgr.run()
}

//监听自己的newconfchan,有了新的配置过来后做对应的处理

func (t *tailLogMgr) run() {
	for {
		select {
		// 3.更新配置
		case newConf := <-t.newConfChan:
			//1.新增配置
			for _, conf := range newConf {
				mk := fmt.Sprintf("%s_%s", conf.Path, conf.Topic)
				_, ok := t.tskMap[mk]
				if ok {
					continue
				} else {
					tailObj := NewTailTask(conf.Path, conf.Topic)
					t.tskMap[mk] = tailObj
				}
			}
			for _, c1 := range t.logEntry { //从原配置中依次拿出配置项与新的配置中逐一进行比较，
				isDelete := true
				for _, c2 := range newConf {
					if c2.Path == c1.Path && c2.Topic == c1.Topic {
						isDelete = false
						continue
					}
					if isDelete {
						mk := fmt.Sprintf("%s_%s", c1.Path, c1.Topic)
						t.tskMap[mk].cancleFunc()
					}
				}
			}
			fmt.Println("get new config ", newConf)
			//找出原来的t.logentry中有，现在newconf中没有的，要删掉

		default:
			time.Sleep(time.Second)
		}
	}
}

//向外暴露一个函数，向外暴露tskMgr的newConfchan
//给一个内部私有的字段传值
func NewConfChan() chan<- []*etcd.LogEntry {
	return tskMgr.newConfChan
}
