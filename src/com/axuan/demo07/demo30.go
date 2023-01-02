package main

import "fmt"

// 结构体指针变量初始化

type Student2 struct {
	id   int
	name string
	sex  byte // 字符类型
	age  int
	addr string
}

func main() {
	// 顺序初始化，每个成员必须初始化，别忘了&
	var p1 *Student2 = &Student2{1, "mike", 'm', 18, "bj"}
	fmt.Println("p1 = ", p1)

	// 指定成员初始化，没有初始化的成员，自动赋值为0
	p2 := &Student2{name: "mike", addr: "bj"}
	fmt.Printf("p2 type is %T\n", p2)
	fmt.Println("p2 = ", p2)
}
