package main

import "fmt"

type person struct {
	name string
	age  int
}
type myInt int

func newPerson(name string, age int) *person {
	return &person{
		name: name,
		age:  age,
	}
}

//方法和接受者
//非指针类型的接受者实际上是在这个函数体内复制了一个新的值
func (p person) changeAge1(newAge int) {
	p.age = newAge
	fmt.Println(p.age)
}

//指针类型的接受者
//在调用这个方法的时候是在内存的指针位置修改了值
func (p *person) changeAge2(newAge int) {
	p.age = newAge
	fmt.Println(p.age)
}

//对自定义类型建立方法
func (m myInt) printSelf(x int) {
	fmt.Println(x)
}
func main() {
	p1 := newPerson("李四", 19)
	p1.changeAge1(30)
	p1.changeAge2(30)
	fmt.Println(p1)
	var num myInt
	num.printSelf(10)
}
