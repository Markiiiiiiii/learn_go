package main

import (
	"context"
	"fmt"
	"time"

	"go.etcd.io/etcd/clientv3"
)

func main() {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"192.168.101.75:2379"},
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
	_, err = cli.Put(ctx, "abcd", "efgh")
	cancle()
	if err != nil {
		fmt.Println("put to etcd faild,err:", err)
		return
	}
	//GET一个K-V
	ctx, cancle = context.WithTimeout(context.Background(), time.Second)
	resp, err := cli.Get(ctx, "abcd")
	cancle()
	if err != nil {
		fmt.Println("get fiald,err:", err)
		return
	}
	for _, ev := range resp.Kvs {
		fmt.Println(ev.Key, ev.Value)
	}
	//watch 监视一个KEY的变化
	rch := cli.Watch(context.Background(), "abcd")
	for wrsep := range rch {
		for _, ev := range wrsep.Events {
			fmt.Printf("Type:%s Key:%s Value:%s\n", ev.Type, ev.Kv.Key, ev.Kv.Value)
		}
	}
}
