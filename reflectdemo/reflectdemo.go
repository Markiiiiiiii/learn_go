package main

import (
	"fmt"
	"reflect"
)

type myInt int64

// typeof用来获取传入参数的类型
func reflectTypeInfo(x interface{}) {
	v := reflect.TypeOf(x)
	fmt.Printf("type:%v,kind:%v\n", v.Name(), v.Kind())
}

// valueof用来获取传入参数的值
func reflectValueOf(x interface{}) {
	v := reflect.ValueOf(x)
	k := v.Kind()
	fmt.Println(v)
	fmt.Println(k)
}

func main() {
	var a float32 = 3.1415926
	reflectTypeInfo(a)
	var b int64 = 93814321
	reflectTypeInfo(b)
	c := "test string"
	reflectValueOf(c)
}
