package main

import "fmt"

// 定义函数时，在函数名后面()定义的参数叫形参
// 参数传递时，只能由实参传递给形参，不能反过来，单向传递

func MyFunc01(a int) {
	fmt.Println("a = ", a)
}

func MyFunc02(a int, b int) {
	fmt.Printf("a = %d, b = %d\n", a, b)
}

func MyFunc03(a, b int) {
	fmt.Printf("a = %d, b = %d\n", a, b)
}

func MyFunc04(a int, b string, c float64) {
}

func MyFunc05(a, b string, c float64, d, e int) {
}

func main() {
	// 有参无返回值函数调用：函数名(所需参数)
	// 调用参数传递的参数叫形参
	MyFunc01(666)

	MyFunc02(666, 777)
}
