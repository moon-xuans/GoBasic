package main

import "fmt"

// 普通变量的方法集

type Person9 struct {
	name string // 名字
	sex  byte   // 性别
	age  int    // 年龄
}

func (p Person9) SetInfoValue() {
	fmt.Println("SetInfoValue")
}

func (p *Person9) SetInfoPointer() {
	fmt.Println("SetInfoPointer")
}

func main() {
	// 结构体变量是一个指针变量，它能够调用哪些方法，这些方法就是一个集合，简称方法集
	p := &Person9{"mike", 'm', 18}
	//p.SetInfoValue() // func (p *Person9).SetInfoPointer
	(*p).SetInfoPointer() // 把(*p)转换成p后再调用，等价于上面

	// 内部做的转换，先把指针p，转换成*p后再调用
	//(*p).SetInfoValue()
	//p.SetInfoValue()
}
