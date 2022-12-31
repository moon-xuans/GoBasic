package main

import "fmt"

// 数组指针做函数参数

// p指向实现数组，它是指向数组，它是数组指针
// *p代表指针所指向的内存，就是实参a
func modify2(p *[5]int) {
	(*p)[0] = 666
	fmt.Println("modify *a = ", *p)
}

func main() {
	a := [5]int{1, 2, 3, 4, 5} // 初始化

	modify2(&a) // 地址传递

	fmt.Println("main: a = ", a)
}
