package main

import "fmt"

// 同名字段

type Person3 struct {
	name string // 名字
	sex  byte   // 性别
	age  int    // 年龄
}

type Student3 struct {
	Person3 // 只有类型，没有名字，匿名字段，继承了Person的成员
	id      int
	addr    string
	name    string // 和Person同名了
}

func main() {
	// 声明(定义一个变量)
	var s Student3

	// 默认规则(就近原则)，如果能在本作用域找到此成员，就操作此成员
	// 如果没找到，找继承的字段

	s.name = "mike" // 操作的是Student中的name
	s.sex = 'm'
	s.age = 18
	s.addr = "bj"

	// 显式调用
	s.Person3.name = "yoyo" // Person的name

	fmt.Printf("s = %+v\n", s)
}
