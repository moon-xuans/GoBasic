package main

import "fmt"

// 方法的继承

type Person11 struct {
	name string // 名字
	sex  byte   // 性别，字符类型
	age  int    // 年龄
}

// Person类型，实现了一个方法
func (tmp *Person11) PrintInfo() {
	fmt.Printf("name=%s, sex=%c, age=%d\n", tmp.name, tmp.sex, tmp.age)
}

type Student11 struct {
	Person11 // 匿名字段
	id       int
	addr     string
}

func main() {
	s := Student11{Person11{"mike", 'm', 18}, 666, "bj"}
	s.PrintInfo()
}
