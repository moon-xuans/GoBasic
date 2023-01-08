package main

import "fmt"

// 值语义和引用语义

type Person7 struct {
	name string // 名字
	sex  byte   // 性别
	age  int    // 年龄
}

func (p Person7) SetInfoValue(n string, s byte, a int) {
	p.name = n
	p.sex = s
	p.age = a
	fmt.Printf("SetInfoValue &p = %p\n", &p) // %p打印地址
}

func (p *Person7) SetInfoPointer(n string, s byte, a int) {
	p.name = n
	p.sex = s
	p.age = a
	fmt.Printf("SetInfoPointer &p = %p\n", p) // %p打印地址
}

func main() {
	s1 := Person7{"go", 'm', 22}
	fmt.Printf("&s1 = %p\n", &s1) // 打印地址

	// 值语义
	//s1.SetInfoValue("mike", 'm', 18)
	//fmt.Println("s1 = ", s1) // 打印内容

	// 引用语义
	(&s1).SetInfoPointer("mike", 'm', 18)
	fmt.Println("s1 = ", s1) // 打印内容
}
