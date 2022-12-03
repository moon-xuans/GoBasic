package main // 要生成Go可执行程序，必须建立一个名字为main的包，并且在该包中包含一个main()的函数
import "fmt"

// 第一个go程序
func main() { // 强制要求左花括号{
	/*
		go build ... 只会编译代码，不能运行可执行文件
		go run ... 只会运行，不会生成可执行文件
	*/

	fmt.Println("hello go")
}
