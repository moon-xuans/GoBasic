package main

import "fmt"

// 结构体成员的使用：指针变量

type Student4 struct {
	id   int
	name string
	sex  byte // 字符类型
	age  int
	addr string
}

func main() {
	// 1.指针有合法指向后，才操作成员
	var s Student4
	var p1 *Student4
	p1 = &s

	// 通过指针操作成员, p1.id 和 (*p1).id完全等价，只能使用.运算符
	p1.id = 1
	(*p1).name = "mike"
	p1.sex = 'm'
	p1.age = 18
	p1.addr = "bj"
	fmt.Println("p1 = ", p1)

	// 2.通过new申请一个结构体
	p2 := new(Student4)
	p2.id = 1
	p2.name = "mike"
	p2.sex = 'm'
	p2.age = 18
	p2.addr = "bj"
	fmt.Println("p2 = ", p2)
}
