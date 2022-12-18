package main

import "fmt"

var b byte // 全局变量

func main() {
	var b int // 局部变量

	// 1.不同作用域，允许定义同名变量
	// 2.使用变量的原则，就近原则

	fmt.Printf("%T\n", b)

	{
		var b float32
		fmt.Printf("2: %T\n", b)
	}

	test7()
}

func test7() {
	fmt.Printf("3: %T\n", b)
}