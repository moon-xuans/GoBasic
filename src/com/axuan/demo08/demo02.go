package main

import "fmt"

// 成员的操作

type Person2 struct {
	name string // 名字
	sex  byte   // 性别
	age  int    // 年龄
}

type Student2 struct {
	Person2 // 只有类型，没有名字，匿名字段，继承了Person的成员
	id      int
	addr    string
}

func main() {
	s1 := Student2{Person2{"mike", 'm', 18}, 1, "bj"}
	s1.name = "yoyo"
	s1.sex = 'f'
	s1.age = 22
	s1.id = 666
	s1.addr = "sz"

	s1.Person2 = Person2{"go", 'm', 18}

	fmt.Println(s1.name, s1.sex, s1.age, s1.id, s1.addr)
}
