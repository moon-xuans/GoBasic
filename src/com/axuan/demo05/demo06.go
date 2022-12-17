package main

import "fmt"

// 多个返回值

func myFunc1() (int, int, int) {
	return 1, 2, 3
}

// go官方推荐写法
func myFunc22() (a int, b int, c int) {
	a, b, c = 111, 222, 333
	return
}

func myFunc33() (a, b, c int) {
	a, b, c = 111, 222, 333
	return
}
func main() {
	// 函数调用
	a, b, c := myFunc33()
	fmt.Printf("a = %d, b = %d, c = %d", a, b, c)
}
