package main

import "fmt"

// 闭包的特点
// 函数的返回值是一个匿名函数，返回一个函数类型
func test22() func() int {
	var x int // 没有初始化，值为0

	return func() int {
		x++
		return x * x
	}
}

func main() {
	// 返回值为一个匿名函数，返回一个函数类型，通过f来调用返回的匿名函数，f来调用闭包函数
	// 它不关心这些捕获了的变量和常量是否已经超出了作用域
	// 所以只有闭包还在使用它，这些变量就还会存在

	f := test22()
	fmt.Println(f()) // 1
	fmt.Println(f()) // 4
	fmt.Println(f()) // 9
	fmt.Println(f()) // 16
	fmt.Println(f()) // 25
}

func test21() int {
	// 函数被调用时,x才分配空间，才初始化为0
	var x int // 没有初始化， 值为0
	x++
	return x * x // 函数调用完毕，x自动释放
}

func main2() {
	fmt.Println(test21())
	fmt.Println(test21())
	fmt.Println(test21())
	fmt.Println(test21())
}
