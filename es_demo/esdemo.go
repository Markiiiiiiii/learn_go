package main

import (
	"context"
	"fmt"

	"github.com/olivere/elastic/v7"
)

//es demo 插入数据
type student struct {
	Name    string `json:"name"`
	Age     int    `json:"age"`
	Merried bool   `json:"merried"`
}

func main() {
	clinet, err := elastic.NewClient(elastic.SetURL("http://192.168.101.76:9200"))
	if err != nil {
		panic(err)
	}
	fmt.Println("Content to ES server sucess!")
	p1 := student{
		Name:    "binon",
		Age:     22,
		Merried: true,
	}
	//链式操作
	put1, err := clinet.Index().
		Index("student").
		Type("go").
		BodyJson(p1).
		Do(context.Background())
	if err != nil {
		panic(err)
	}
	fmt.Printf("Index: %s \nType: %s\nUser: %s\n", put1.Index, put1.Type, put1.Id)
}
