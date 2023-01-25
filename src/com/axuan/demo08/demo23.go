package main

import "fmt"

// 显式调用panic函数

func testa() {
	fmt.Println("aaaaaaaaa")
}

func testb() {
	//fmt.Println("bbbbbbbbbbbbbbbb")
	// 显式调用panic函数，导致程序中断
	panic("this is a panic test")
}

func testc() {
	fmt.Println("cccccccccccccccc")
}

func main() {
	testa()
	testb()
	testc()
}
