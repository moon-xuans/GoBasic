package main

import (
	"fmt"
)

// 非结构体匿名字段

type mystr string // 自定义类型，给一个类型改名

type Person4 struct {
	name string // 名字
	sex  byte   // 性别
	age  int    // 年龄
}

type Student4 struct {
	Person4 // 结构体匿名字段
	int     // 基础类型的匿名字段
	mystr
}

func main() {
	s := Student4{Person4{"mike", 'm', 18}, 666, "hehehe"}
	fmt.Printf("s = %+v\n", s)

	s.Person4 = Person4{"go", 'm', 22}

	fmt.Println(s.name, s.age, s.sex, s.int, s.mystr)
	fmt.Println(s.Person4, s.int, s.mystr)
}
