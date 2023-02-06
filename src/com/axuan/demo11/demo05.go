package main

import (
	"fmt"
	"runtime"
)

// Goexit的使用

func test() {
	defer fmt.Println("cccccc")

	// return // 终止此函数
	runtime.Goexit() // 终止所在的协程

	fmt.Println("ddddddddddddd")
}

func main() {
	// 创建新建的协程
	go func() {
		fmt.Println("aaaaaaaaaa")

		// 调用了别的函数
		test()

		fmt.Println("bbbbbbbbbbb")
	}()

	// 特地写一个死循环，目的不让主协程结束
	for {

	}
}
