package main

import "fmt"

// 函数递归调用的流程
func test01(a int) {
	if a == 1 { // 函数终止调用的条件，非常重要
		fmt.Println("a = ", a)
		return // 终止函数调用
	}
	test01(a - 1)
	fmt.Println("a = ", a)
}

func main() {
	test01(3)
	fmt.Println("main")
}
