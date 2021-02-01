package main

import (
	"fmt"

	"github.com/Shopify/sarama"
)

func main() {
	config := sarama.NewConfig()                              //新建一个配置
	config.Producer.RequiredAcks = sarama.WaitForAll          //发送完数据后需要leader和follow都确认ack
	config.Producer.Partitioner = sarama.NewRandomPartitioner //轮询新选出一个partition
	config.Producer.Return.Successes = true                   //成功交付了消息将在sucssce channel 返回一个ture

	//构建一个消息结构体
	msg := &sarama.ProducerMessage{}
	msg.Topic = "web_log"
	msg.Value = sarama.StringEncoder("this is second test log") //对消息进行编码

	//连接kafka
	clinet, err := sarama.NewSyncProducer([]string{"192.168.1.102:9092"}, config) //【】string{}可以使是集群地址组
	if err != nil {
		fmt.Println("producer closed,err:", err)
		return
	}
	defer clinet.Close()

	//发送消息
	pid, offset, err := clinet.SendMessage(msg)
	if err != nil {
		fmt.Println("send msg faild,err:", err)
		return
	}
	fmt.Printf("pid:%v,offset:%v\n", pid, offset)
}
