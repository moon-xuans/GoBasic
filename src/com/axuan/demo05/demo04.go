package main

import "fmt"

// 不定参数传递

func myFunc(tmp ...int) {
	for _, data := range tmp {
		fmt.Println("data = ", data)
	}
}

func myFunc2(tmp ...int) {
	for _, data := range tmp {
		fmt.Println("data = ", data)
	}
}

func test(args ...int) {
	// 全部元素传递给myFunc
	//myFunc(args...)

	// 只想把后2个参数传递给另外一个函数使用
	myFunc2(args[:2]...) // args[0] ~ args[2] (不包括数字args[2])，传递过去
	fmt.Println("------------------")
	myFunc2(args[2:]...) // 从args[2]开始(包括本身)，把后面所有元素传递过去
}

func main() {
	test(1, 2, 3, 4)
}
