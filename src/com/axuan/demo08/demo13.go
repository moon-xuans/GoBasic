package main

import "fmt"

// 方法值

type Person13 struct {
	name string // 名字
	sex  byte   // 性别，字符类型
	age  int    // 年龄
}

func (p Person13) SetInfoValue() {
	fmt.Printf("SetInfoValue:%p, %v\n", &p, p)
}

func (p *Person13) SetInfoPointer() {
	fmt.Printf("SetInfoPointer:%p, %v\n", p, p)
}

func main() {
	p := Person13{"mike", 'm', 18}
	fmt.Printf("main:%p, %v\n", &p, p)

	p.SetInfoPointer() // 传统调用方式

	// 保存方式入口地址
	pFunc := p.SetInfoPointer // 这个就是方法值，调用函数时，无需再传递接收者，隐藏了接收者
	pFunc()

	vFunc := p.SetInfoValue
	vFunc() // 等价于p.SetInfoValue()

}
