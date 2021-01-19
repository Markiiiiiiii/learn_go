package main

import (
	"fmt"
	"reflect"
)

type Studens struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	a := Studens{
		Name: "张三",
		Age:  24,
	}
	k := reflect.TypeOf(a)
	// v := reflect.ValueOf(a)

	// fmt.Println(k.Name(), k.Kind())
	// //使用for遍历结构体中的所有字段信息
	// for i := 0; i <= k.NumField(); i++ {
	// 	field := k.Field(i)
	// 	fmt.Printf("name:%s,index:%d,type:%v,json tag:%v\n", field.Name, field.Index, field.Type, field.Tag.Get("json"))
	// }
	//通过字段名获取指定结构体字段信息
	if s, ok := k.FieldByName("Age"); ok {
		fmt.Printf("name:%s,index:%d,type:%v,json tag:%v\n", s.Name, s.Index, s.Type, s.Tag.Get("json"))
	}
}
