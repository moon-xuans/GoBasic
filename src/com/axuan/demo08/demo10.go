package main

import "fmt"

// 普通变量的方法集

type Person10 struct {
	name string // 名字
	sex  byte   // 性别
	age  int    // 年龄
}

func (p Person10) SetInfoValue() {
	fmt.Println("SetInfoValue")
}

func (p *Person10) SetInfoPointer() {
	fmt.Println("SetInfoPointer")
}

func main() {
	p := Person10{"mike", 'm', 18}
	p.SetInfoPointer() // func (p *Person10) SetInfoPointer()
	// 内部，先把p，转换成&p再调用，(&p).SetInfoPointer()
}
