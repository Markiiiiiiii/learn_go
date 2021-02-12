package kafka

import (
	"fmt"

	"github.com/Shopify/sarama"
	"learn_go.com/logtransfer/es"
)

//LogData ...
type LogData struct {
	Data string `json:"data"`
}

//初始化Kafka消费者从kafka取数据发给es
func Init(addrs []string, topic string) error {

	consumer, err := sarama.NewConsumer(addrs, nil)
	if err != nil {
		fmt.Println("don't content to kafka server! err:", err)
		return err
	}
	partitionList, err := consumer.Partitions(topic) //获取到该条topic所存储的每个分区
	if err != nil {
		fmt.Println("don't get partitions, err:", err)
		return err
	}
	fmt.Println(partitionList)
	for partition := range partitionList {
		// 遍历topic所所存储的所有分区
		// 针对每个分区创建一个对应的分区消费者
		pc, err := consumer.ConsumePartition(topic, int32(partition), sarama.OffsetNewest)
		if err != nil {
			fmt.Println("don't get partition,err:", err)
			return err
		}
		// defer pc.AsyncClose()
		go func(sarama.PartitionConsumer) {
			for msg := range pc.Messages() {
				fmt.Printf("Partition:%d Offset:%d Key:%v Value:%s\n", msg.Partition, msg.Offset, msg.Key, string(msg.Value))
				//直接发给ES
				ld := map[string]interface{}{
					"data": string(msg.Value),
				}
				// fmt.Printf("%#v\n", ld)
				// if err != nil {
				// 	fmt.Println("unmarshal faild ,err:", err)
				// 	continue
				// }
				err = es.SendToES(topic, ld)
			}
		}(pc)
	}

	// config := sarama.NewConfig()                              //新建一个配置
	// config.Producer.RequiredAcks = sarama.WaitForAll          //发送完数据后需要leader和follow都确认ack
	// config.Producer.Partitioner = sarama.NewRandomPartitioner //轮询新选出一个partition
	// config.Producer.Return.Successes = true                   //成功交付了消息将在sucssce channel 返回一个ture

	// //连接kafka
	// clinet, err = sarama.NewSyncProducer(addrs, config) //[]string{}可以使是集群地址组
	// if err != nil {
	// 	fmt.Println("producer closed,err:", err)
	// 	return
	// }
	// //初始化全局logDataChan
	// logDataChan = make(chan *logData, maxSize)
	// //开启后台的goroutine取数据发往kafka
	// go sendToKafka()
	// return
	return err
}
