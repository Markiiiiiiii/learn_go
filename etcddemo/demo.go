package main

import (
	"time"

	"go.etcd.io/etcd/clientv3"
)

func main() {
	clientv3.New(clientv3.Config{
		Endpoints:   []string{"192.168.1.102：2379"},
		DialTimeout: 5 * time.Second,
	})
}
