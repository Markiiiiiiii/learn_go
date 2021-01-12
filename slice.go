package main

import "fmt"

type person struct {
	name   string
	gender string
}

func fixn(x *person) {
	x.gender = "女"
}
func (p *person) fixname() {
	fmt.Println("name fix ")
}

func main() {
	var p person
	p.name = "张三"
	p.gender = "男"
	fixn(&p)
	fmt.Println(p)
	//结构体指针1
	var p2 = new(person)
	fmt.Printf("%T %p %p\n", p2, p2, &p2)
	//结构体指针2
	var p3 = person{
		name:   "李四",
		gender: "男",
	}
	fmt.Println(p3)
	//	值列表的方式初始化
	p4 := person{
		"王五",
		"女",
	}
	fmt.Println(p4)

	p6 := &person{
		name: "无名氏",
	}
	p6.fixname()

	fmt.Printf("%p", p6)
}
