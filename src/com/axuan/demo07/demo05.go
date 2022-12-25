package main

import "fmt"

// 指针做函数参数
func swap2(p1, p2 *int) {
	*p1, *p2 = *p2, *p1
}
func main() {
	a, b := 10, 20

	// 通过一个函数交换a和b的内容
	swap2(&a, &b)
	fmt.Printf("main: a = %d, b = %d\n", a, b)
}
