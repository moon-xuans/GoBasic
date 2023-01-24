package main

import "fmt"

// 接口的继承

type Humaner16 interface { // 子集
	sayhi()
}

type Personer16 interface { // 超集
	Humaner16 // 匿名字段，继承了sayhi()
	sing(lrc string)
}

type Student16 struct {
	name string
	id   int
}

// Student实现了sayhi()
func (tmp *Student16) sayhi() {
	fmt.Printf("Student[%s, %d] sayhi\n", tmp.name, tmp.id)
}

func (tmp *Student16) sing(lrc string) {
	fmt.Println("Student在唱着：", lrc)
}

func main() {
	// 定义一个接口类型的变量
	var i Personer16
	s := &Student16{"mike", 666}
	i = s

	i.sayhi() // 继承过来的方法
	i.sing("嘻嘻嘻")
}
