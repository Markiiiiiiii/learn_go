package es

import (
	"context"
	"fmt"
	"strings"

	"github.com/olivere/elastic/v7"
)

var clinet *elastic.Client

//初始化es，准备接受kafka那边发来的数据

func Init(address string) (err error) {
	if !strings.HasPrefix(address, "http://") {
		address = "http://" + address
	}
	clinet, err = elastic.NewClient(elastic.SetURL(address))
	if err != nil {
		fmt.Println(">> Elasic content faild,err:", err)
		return
	}
	fmt.Println("Connect to ES success!")
	return
}

//发送数据到es
func SendToES(indexStr string, data interface{}) error {
	//链式操作
	put1, err := clinet.Index().
		Index(indexStr).
		// Type(typeStr).
		BodyJson(data).
		Do(context.Background())
	if err != nil {
		return err
	}
	fmt.Printf("Index: %s \nType: %s\nUser: %s\n", put1.Index, put1.Type, put1.Id)
	return err
}
