package main

import "fmt"

// 全局变量
func test6() {
	fmt.Println("test a = ", a)
}

// 定义在函数外部的变量是局部变量
// 全局变量在任何地方都能使用
var a int

func main() {
	a = 10
	fmt.Println("a = ", a)

	test6()
}