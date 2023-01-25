package main

import "fmt"

// recover的使用

func test05a() {
	fmt.Println("aaaaaaaaa")
}

func test05b(x int) {
	// 设置recover
	defer func() {
		// recover() // 可以打印panic的错误信息
		fmt.Println(recover())
		if err := recover(); err != nil { // 产生了panic异常
			fmt.Println(err)
		}
	}() // (),调用此匿名函数

	var a [10]int
	a[x] = 111 // 当x为20时候，导致数组越界，产生一个panic，导致程序崩溃
}

func test05c() {
	fmt.Println("cccccccccccccccc")
}

func main() {
	test05a()
	test05b(20)
	test05c()
}
