package main

import (
	"fmt"
	"runtime"
)

// GOMAXPROCS的使用

func main() {
	n := runtime.GOMAXPROCS(4) // 指定以4核运算
	fmt.Println("n = ", n)

	for {
		go fmt.Print(1)

		fmt.Print(0)
	}
}
