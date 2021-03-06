package main2

import (
	"context"
	"fmt"
	"time"

	"go.etcd.io/etcd/clientv3"
)

func main2() {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"192.168.1.102:2379"},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		fmt.Println("content to etcd faild ,err:", err)
		return
	}
	fmt.Println("content to etcd server!")
	defer cli.Close()
	//PUT一个k-v
	ctx, cancle := context.WithTimeout(context.Background(), time.Second)
	value := `[{"path":"./access.log","topic":"web_log"},{"path":"./redis.log","topic":"redis_log"}]`
	_, err = cli.Put(ctx, "/logangent/collect_config", value)
	cancle()
	if err != nil {
		fmt.Println("put to etcd faild,err:", err)
		return
	}

	//watch 监视一个KEY的变化
	// rch := cli.Watch(context.Background(), "abcd")
	// for wrsep := range rch {
	// 	for _, ev := range wrsep.Events {
	// 		fmt.Printf("Type:%s Key:%s Value:%s\n", ev.Type, ev.Kv.Key, ev.Kv.Value)
	// 	}
	// }
}
