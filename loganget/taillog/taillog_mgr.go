package taillog

import "learn_golang/loganget/etcd"

//建立一个tailobj的管理者，将所有etcd中获取到的path存储到管理者这个切片中，
//在taillog中热加载tailobj的对象
type tailLogMgr struct {
	logEntry []*etcd.LogEntry
	tskMap   map[string]*TailTask
}

var tskMgr *tailLogMgr

func Init(logEntryConf []*etcd.LogEntry) {
	tskMgr = &tailLogMgr{
		logEntry: logEntryConf, //把当前的日志收集项配置信息保存起来
	}
	for _, logEntry := range logEntryConf {
		// logEntry.Path 要收集的日志路径
		NewTailTask(logEntry.Path, logEntry.Topic)
		// 3.2发往kafka

	}
}
