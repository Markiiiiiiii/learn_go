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

	switch k {
	case reflect.Int:
		fmt.Printf("this a int ,value:%d,kind:%v\n", int(v.Int()), k)
	case reflect.Float32:
		fmt.Printf("this a float32,value:%f,kinde:%v/n", float32(v.Float()), k)

	}
}

// 通过反射设置变量的值
func reflectSetValue1(x interface{}) { //错误示范
	v := reflect.ValueOf(x)
	if v.Kind() == reflect.Int64 {
		v.SetInt(3000) //修改的是副本 会引发Panic错误
	}
}

func reflectSetValue2(x interface{}) { //正确示范
	v := reflect.ValueOf(x)
	if v.Elem().Kind() == reflect.Int64 { //必须使用elem()方法获取指针对应的值才能修改
		v.Elem().SetInt(3000) //修改的是指针
	}
}

func main() {
	// 	var a float32 = 3.1415926
	// 	reflectTypeInfo(a)
	// var b int64 = 93814321
	// 	reflectTypeInfo(b)
	// c := 123
	// reflectValueOf(c)
	// reflectSetValue1(&b)
	// reflectSetValue2(&b)
	// fmt.Println(b)

	var a *int
	fmt.Println("var a *int IsNil?", reflect.ValueOf(a).IsNil()) //判断是否为空指针

	fmt.Println("var a *int IsValid?", reflect.ValueOf(a).IsValid()) //判断是否是有效的返回值
	fmt.Println("var a *int IsZero?", reflect.ValueOf(a).IsZero())   //判断是否是是零值
	b := struct{}{}
	fmt.Println("查看是否不存在一个结构体成员：", reflect.ValueOf(b).FieldByName("abc").IsValid())
	fmt.Println("查看是否不存在一个结构体方法：", reflect.ValueOf(b).MethodByName("abc").IsValid())
	c := map[string]int{}
	fmt.Println("查看Map是否不存在一个键：", reflect.ValueOf(c).MapIndex(reflect.ValueOf("李四")).IsValid())
}
