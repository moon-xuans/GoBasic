package main

import (
	"fmt"
	"os"
)

// 获取命令行参数

func main() {
	// 接收拥护传递的参数，都是以字符串方式传递
	list := os.Args

	n := len(list)
	fmt.Println("n = ", n)

	for i, data := range list {
		fmt.Printf("list[%d] = %s\n", i, data)
	}
}
