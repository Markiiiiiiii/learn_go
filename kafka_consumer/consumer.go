package main

import (
	"fmt"

	"github.com/Shopify/sarama"
)

//kafka消费者实例
func main() {
	consumer, err := sarama.NewConsumer([]string{"192.168.101.76:9092"}, nil)
	if err != nil {
		fmt.Println("don't content to kafka server! err:", err)
		return
	}
	partitionList, err := consumer.Partitions("web_log") //获取到该条topic所存储的每个分区
	if err != nil {
		fmt.Println("don't get partitions, err:", err)
		return
	}
	fmt.Println(partitionList)
	for partition := range partitionList {
		// 遍历topic所所存储的所有分区
		// 针对每个分区创建一个对应的分区消费者
		pc, err := consumer.ConsumePartition("web_log", int32(partition), sarama.OffsetNewest)
		if err != nil {
			fmt.Println("don't get partition,err:", err)
			return
		}
		defer pc.AsyncClose()
		go func(sarama.PartitionConsumer) {
			for msg := range pc.Messages() {
				fmt.Printf("Partition:%d\n Offset:%d\n Key:%v\n Value:%s\n", msg.Partition, msg.Offset, msg.Key, string(msg.Value))
			}
		}(pc)
		select {}
	}
}
