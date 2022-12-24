package main

import (
	"GoLearning/src/com/axuan/demo06/calc"
	"fmt"
)

// init函数作用：如果导入这一个包，他会在函数执行前，先执行init方法
func init() {
	fmt.Println("this is main init")
}

func main() {
	/*
		// 同级目录
		// 1.分文件编程(多个源文件)，必须放在src目录
		// 2.设置GOPATH环境变量
		// 3.同一个目录，包名必须一样
		// 4.go env查看go相关的环境路径
		// 5.同一个目录，调用别的文件的函数，直接调用即可，无需包名引用

		test()

	*/

	// 不同目录
	// 不同目录，包名不一样
	// 调用不同包里面的函数，格式：包名.函数名()
	// 调用别的包的函数，这个包函数名字如果首字母是小写，无法让别人调用，要想别人能调用，必须首字母大写
	fmt.Println("res = ", calc.Add(10, 20))

	fmt.Println("res = ", calc.Minus(10, 5))

	// 如果有多个文件或多个包
	// 配置GOPATH环境变量，配置src同级目录的绝对路径
	// 自动生成bin或pkg目录，需要使用go install命令(了解),除了要配置GOPATH环境变量，还要配置GOBIN环境变量
	// src:放源代码；bin：放可执行程序；pkg：放平台相关的库
}
