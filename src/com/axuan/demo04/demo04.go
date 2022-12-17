package main

import "fmt"

// 一个返回值

// 有返回值的函数需要通过return中断函数，通过return返回
func myFunc01() int {
	return 666
}

// 给返回值起一个变量名，go推荐写法
func myFunc02() (result int) {
	return 666
}

// 给返回值起一个变量名，go推荐写法
func myFunc03() (result int) {
	result = 666
	return
}

func main() {
	// 无参有返回值函数调用
	var a int
	a = myFunc01()
	fmt.Println("a = ", a)

	b := myFunc01()
	fmt.Println("b = ", b)

	c := myFunc03()
	fmt.Println("c = ", c)
}
