package main

import (
	"fmt"
	"os"
)

//使用结构体设计学生系统
//结构体中保存学生的信息
type studyInfo struct {
	info map[string]int
}

//结构体中查看所有学生的方法
func (s studyInfo) listAll() {
	for k, v := range s.info {
		fmt.Printf("学生名：%s 年龄：%v\n", k, v)
	}
}

//结构体有新增学生的方法
func (s studyInfo) addInfo(name string, ages int) {
	s.info[name] = ages
}

//结构体中删除学生的方法
func (s studyInfo) delectInfo(name string) {
	delete(s.info, name)
}
func main() {
	var p studyInfo
	var ennum, age int
	var name string
	p.info = make(map[string]int)
	for {
		fmt.Print(`
学生管理系统
1 查看所有学生
2 新增学生信息
3 删除学生信息
4 退出
请输入您的要求：
`)
		fmt.Scanln(&ennum)
		switch ennum {
		case 1:
			p.listAll()
		case 2:
			fmt.Println("请输入 学生名 年龄")
			fmt.Scanln(&name, &age)
			p.addInfo(name, age)

		case 3:
			fmt.Println("请输入要删除的 学生名")
			fmt.Scanln(&name)
			p.delectInfo(name)
		case 4:
			os.Exit(1)
		}
		// p.addInfo("张三", 19)
		// p.listAll()
	}
}
