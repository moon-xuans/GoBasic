package main

import "fmt"

func test24a() {
	fmt.Println("aaaaaaaaa")
}

func test24b(x int) {
	var a [10]int
	a[x] = 11 // 当x为20时，导致数组越界，产生一个panic，导致程序崩溃
}

func test24c() {
	fmt.Println("cccccccccccccccc")
}

func main() {
	test24a()
	test24b(20)
	test24c()
}
