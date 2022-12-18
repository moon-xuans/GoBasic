package main

import "fmt"

// 多个defer的执行顺序——先进后出

func test3(x int) {
	result := 100 / x
	fmt.Println("result = ", result)
}

func main() {
	defer fmt.Println("aaaaaa")

	defer fmt.Println("bbbbbbb")

	// 调用一个函数，导致内存出问题
	defer test3(0)

	defer fmt.Println("cccccccc")
}
