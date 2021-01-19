package main

import (
	"fmt"
	"time"
)

func main() {
	nowtime := time.Now()
	fmt.Println(nowtime)
	fmt.Println(nowtime.Clock())
	fmt.Println(nowtime.Date())
	fmt.Println(nowtime.GobEncode())
	fmt.Println(nowtime.Zone())
	fmt.Println(nowtime.Hour())
	fmt.Println(nowtime.Second())
	fmt.Println(nowtime.Minute())
	fmt.Println(nowtime.Year())
	fmt.Println(nowtime.YearDay())
	// 时间戳
	fmt.Println(nowtime.Unix())
	//纳秒时间戳
	fmt.Println(nowtime.UnixNano())
	// 时间间隔常量
	fmt.Println(time.Second)
	//增加时间 增加1小时
	fmt.Println(nowtime.Add(1 * time.Hour).Hour())

	// 定时器
	// 	timer := time.Tick(time.Second)
	// 	for t := range timer {
	// 		fmt.Println(t.Second())
	// 	}
	// 格式化时间，2006-01-02 15：04：05
	fmt.Println(nowtime.Format("2006-01-02"))
	fmt.Println(nowtime.Format("2006-01-02 15:04:05"))
	fmt.Println(nowtime.Format("2006/01/02 03-04-05 PM"))
	fmt.Println(nowtime.Format("2006-01-02 15:04:05.000.000"))
	//格式化字符串为时间格式
	tim, _ := time.Parse("2006-01-02", "2021-01-03")
	fmt.Println(tim.Date())
	//时间差
	fmt.Println(nowtime.Sub(tim))

	//sleep
	// n := 5
	// fmt.Println("开始sleep")
	// // time.Sleep(time.Duration(n) * time.Second) //将变量格式化为duration类型
	// fmt.Println("5秒过去了")

	//按照指定时区格式化时间
	loc, err := time.LoadLocation("Asia/Beijing")
	if err != nil {
		return
	}
	timeObj, err := time.ParseInLocation("2006-01-02 15:04:05", "2019-12-25 23:59:59", loc)
	if err != nil {
		return
	}
	fmt.Println(timeObj)
}
