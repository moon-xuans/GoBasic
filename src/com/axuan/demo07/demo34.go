package main

import "fmt"

type Student6 struct {
	id   int
	name string
	sex  byte // 字符类型
	age  int
	addr string
}

func test02(p *Student6) {
	p.id = 666
}

func main() {
	s := Student6{1, "mike", 'm', 18, "bj"}

	test02(&s) // 地址传递(引用传递),形参可以改实参
	fmt.Println("main: ", s)
}

func test11(s Student6) {
	s.id = 666
	fmt.Println("test01: ", s)
}

func main11() {
	s := Student6{1, "mike", 'm', 18, "bj"}

	test11(s) // 值传递，形参无法改实参
	fmt.Println("main: ", s)
}
