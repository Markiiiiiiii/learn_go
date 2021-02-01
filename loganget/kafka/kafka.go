package kafka

import (
	"fmt"

	"github.com/Shopify/sarama"
)

//往kafka里写日志模块
var (
	clinet sarama.SyncProducer //声明一个全局的连接kafka的生产者客户端

)

//初始化clinet
func Init(addrs []string) (err error) {
	config := sarama.NewConfig()                              //新建一个配置
	config.Producer.RequiredAcks = sarama.WaitForAll          //发送完数据后需要leader和follow都确认ack
	config.Producer.Partitioner = sarama.NewRandomPartitioner //轮询新选出一个partition
	config.Producer.Return.Successes = true                   //成功交付了消息将在sucssce channel 返回一个ture

	//连接kafka
	clinet, err = sarama.NewSyncProducer(addrs, config) //[]string{}可以使是集群地址组
	if err != nil {
		fmt.Println("producer closed,err:", err)
		return
	}
	return
}

func SendToKafka(topic, msgData string) {
	//构建一个消息结构体
	msg := &sarama.ProducerMessage{}
	msg.Topic = topic
	msg.Value = sarama.StringEncoder(msgData) //对消息进行编码

	// 发送到kafka

	pid, offset, err := clinet.SendMessage(msg)
	if err != nil {
		fmt.Println("send msg faild,err:", err)
		return
	}
	fmt.Printf("pid:%v,offset:%v\n", pid, offset)

}
