package etcd

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"go.etcd.io/etcd/clientv3"
)

var (
	cli *clientv3.Client
)

//需要收集的日志配置信息
type LogEntry struct {
	Path  string `json:"path"`  //日志存放路径
	Topic string `json:"topic"` //日志要发往kafka中的哪个topic
}

//初始化etcd
func Init(addr string, timeout time.Duration) (err error) {

	cli, err = clientv3.New(clientv3.Config{
		Endpoints:   []string{addr},
		DialTimeout: timeout,
	})
	if err != nil {
		fmt.Println("content to etcd faild ,err:", err)
		return
	}
	return
}

//从etcd中根据KEY获取配置项
func GetInfo(key string) (logEntyConf []*LogEntry, err error) {
	ctx, cancle := context.WithTimeout(context.Background(), time.Second)
	resp, err := cli.Get(ctx, key)
	cancle()
	if err != nil {
		fmt.Println("ETCD get key fiald,err:", err)
		return
	}
	for _, ev := range resp.Kvs {
		// fmt.Println(ev.Key, ev.Value)
		err = json.Unmarshal(ev.Value, &logEntyConf)
		if err != nil {
			fmt.Println("umarshal etcd value faild ,err:", err)
			return
		}
	}
	return
}
func WatchConf(key string, newConfch chan<- []*LogEntry) {
	//watch 监视一个KEY的变化
	ch := cli.Watch(context.Background(), key)
	for wrsep := range ch {
		for _, ev := range wrsep.Events {
			fmt.Printf("Type:%s Key:%s Value:%s\n", ev.Type, ev.Kv.Key, ev.Kv.Value)
			//通知taillog.tskmgr
			//1.先判断操作的类型 是否是删除的操作，如果删除的话值就没有了，提前定义一个slice
			var newConf []*LogEntry
			if ev.Type != clientv3.EventTypeDelete {
				//判断是否是删除操作,手动传递一个空的配置项
				err := json.Unmarshal(ev.Kv.Value, &newConf)
				if err != nil {
					fmt.Println("unmarshal faild,err:", err)
					continue
				}
			}
			fmt.Println("get new conf:", newConf)
			newConfch <- newConf
		}
	}
}
