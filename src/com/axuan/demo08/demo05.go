package main

import "fmt"

// 结构体指针类型匿名字段

type Person5 struct {
	name string // 名字
	sex  byte   // 性别
	age  int    // 年龄
}

type Student5 struct {
	*Person5 // 只有类型，没有名字，匿名字段，继承了Person的成员
	id       int
	addr     string
}

func main() {
	s1 := Student5{&Person5{"mike", 'm', 18}, 666, "bj"}
	fmt.Println(s1.name, s1.sex, s1.age, s1.id, s1.addr)

	// 先定义变量
	var s2 Student5
	s2.Person5 = new(Person5) // 分配空间
	s2.name = "yoyo"
	s2.sex = 'm'
	s2.age = 10
	s2.id = 222
	s2.addr = "sz"
	fmt.Println(s2.name, s2.sex, s2.age, s1.id, s1.addr)

}
