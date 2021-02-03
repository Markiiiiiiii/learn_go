package taillog

import "learn_golang/loganget/etcd"

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
		NewTailTask(logEntry.Path, logEntry.Topic)
		// 3.2发往kafka

	}
}

//监听自己的newconfchan,有了新的配置过来后做对应的处理
//1.新增配置
// 2.删除配置
// 3.更新配置
func (t *tailLogMgr) run() {

}
