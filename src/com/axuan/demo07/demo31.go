package main

import "fmt"

// 结构体成员的使用：普通变量

// 定义一个结构体类型
type Student3 struct {
	id   int
	name string
	sex  byte // 字符类型
	age  int
	addr string
}

func main() {
	// 定义一个结构体普通变量
	var s Student3

	// 操作成员，需要使用点(.)运算符
	s.id = 1
	s.name = "mike"
	s.sex = 'm' // 字符
	s.age = 18
	s.addr = "bj"
	fmt.Println("s = ", s)
}