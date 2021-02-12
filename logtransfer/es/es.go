package es

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/olivere/elastic/v7"
)

//LogData ...
type LogData struct {
	Topic string `json:"topic"`
	Data  string `json:"data"`
}

var (
	clinet *elastic.Client
	ch     chan *LogData
)

//初始化es，准备接受kafka那边发来的数据

func Init(address string, chanSize, nums int) (err error) {
	if !strings.HasPrefix(address, "http://") {
		address = "http://" + address
	}
	clinet, err = elastic.NewClient(elastic.SetURL(address))
	if err != nil {
		fmt.Println(">> Elasic content faild,err:", err)
		return
	}
	fmt.Println("Connect to ES success!")
	ch = make(chan *LogData, chanSize)
	for i := 0; i < nums; i++ {
		go SendToES()
	}
	return
}

// 把数据发送到chan中
func SendToESChan(msg *LogData) {
	ch <- msg
}

//发送数据到es
func SendToES() {
	for {
		select {
		case msg := <-ch:
			//链式操作
			put1, err := clinet.Index().
				Index(msg.Topic).
				Type("IP").
				BodyJson(msg).
				Do(context.Background())
			if err != nil {
				fmt.Println("err:", err)
			}
			fmt.Printf("Id:%v Index:%s Type:%s\n", put1.Id, put1.Index, put1.Type)
		default:
			time.Sleep(time.Second)
		}
	}

}
