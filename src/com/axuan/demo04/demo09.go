package main

import "fmt"

// 数字累加
func test11() (sum int) {
	for i := 1; i <= 100; i++ {
		sum += i
	}
	return
}

func test12(i int) int {
	if i == 1 {
		return 1
	}
	return i + test12(i-1)
}

func test13(i int) int {
	if i == 100 {
		return 100
	}
	return i + test13(i+1)
}
func main() {
	var sum int
	//sum = test11()
	//sum = test12(100)
	sum = test13(1)
	fmt.Println("sum = ", sum)
}
