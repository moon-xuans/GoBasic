package main

import "fmt"

// new函数的使用
func main() {
	var p *int

	// p是*int，指向int类型
	p = new(int)

	*p = 666
	fmt.Println("*p = ", *p)

	q := new(int) // 自动推导类型
	*q = 777
	fmt.Println("*q = ", *q)
}
