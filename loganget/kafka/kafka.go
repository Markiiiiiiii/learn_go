package kafka

import (
	"fmt"
	"time"

	"github.com/Shopify/sarama"
)

//往kafka里写日志模块
var (
	clinet      sarama.SyncProducer //声明一个全局的连接kafka的生产者客户端
	logDataChan chan *logData
)

type logData struct {
	topic string
	data  string
}

//初始化clinet
func Init(addrs []string, maxSize int) (err error) {
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
	//初始化全局logDataChan
	logDataChan = make(chan *logData, maxSize)
	//开启后台的goroutine取数据发往kafka
	go sendToKafka()
	return
}

//真正向kafka发送的函数
func sendToKafka() {
	for {
		select {
		case ld := <-logDataChan:

			//构建一个消息结构体
			msg := &sarama.ProducerMessage{}
			msg.Topic = ld.topic
			msg.Value = sarama.StringEncoder(ld.data) //对消息进行编码

			// 发送到kafka

			pid, offset, err := clinet.SendMessage(msg)
			if err != nil {
				fmt.Println("send msg faild,err:", err)
				return
			}
			fmt.Printf("pid:%v,offset:%v\n", pid, offset)
		default:
			time.Sleep(time.Millisecond * 50)
		}
	}

}

//给外部暴露一个函数，该函数只把日志数据发送到一个内部的channel里
func SendToChan(topic, data string) {
	msg := &logData{
		topic: topic,
		data:  data,
	}
	logDataChan <- msg
}
