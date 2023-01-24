package main

import "fmt"

// 接口的定义和实现

type Humaner interface {
	sayhi()
}

type Student15 struct {
	name string
	age  int
}

// Student实现了此方法
func (tmp *Student15) sayhi() {
	fmt.Printf("Student[%s, %d] sayhi\n", tmp.name, tmp.age)
}

type Teacher15 struct {
	addr  string
	group string
}

// Teacher实现了此方法
func (tmp *Teacher15) sayhi() {
	fmt.Printf("Teacher[%s, %s] sayhi\n", tmp.addr, tmp.group)
}

type MyStr string

// MyStr实现了此方法
func (tmp *MyStr) sayhi() {
	fmt.Printf("MyStr[%s] sayhi\n", *tmp)
}

func WhoSayHi(i Humaner) {
	i.sayhi()
}

func main() {
	// 多态的表现

	s := &Student15{"mike", 666}
	t := &Teacher15{"bj", "go"}
	var str MyStr = "hello mike"

	// 调用同一函数，不同表现，多态，多种形态
	WhoSayHi(s)
	WhoSayHi(t)
	WhoSayHi(&str)

	// 创建一个切片
	x := make([]Humaner, 3)
	x[0] = s
	x[1] = t
	x[2] = &str

	// 第一个返回下标，第二个返回下标所对应的值
	for _, i := range x {
		i.sayhi()
	}
}

func main01() {
	// 定义接口类型的变量
	var i Humaner

	// 只是实现了此接口方法的类型，那么这个类型的变量(接收者类型)就可以给i赋值
	s := &Student15{"mike", 666}
	i = s
	i.sayhi()

	t := &Teacher15{"bj", "go"}
	i = t
	i.sayhi()

	var str MyStr = "hello mike"
	i = &str
	i.sayhi()
}
