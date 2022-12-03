package main

import "fmt"

// fmt包的格式化输入输出
func main() {

	/*
		// 格式化输出
		a := 10
		b := "abcd"
		c := 'a'
		d := 3.14
		// %T操作变量所属类型
		fmt.Printf("%T, %T, %T, %T\n", a, b, c, d)

		// %d 整数类型
		// %s 字符串格式
		// %c 字符格式
		// %f 浮点型格式
		fmt.Printf("a = %d, b = %s, c = %c, d = %f\n", a, b, c, d)
		// %v自动匹配格式输出
		fmt.Printf("a = %v, b = %v, c = %v, d = %v", a, b, c, d)

	*/

	// 输入的使用
	var a int // 声明变量
	fmt.Printf("请输入变量a: ")

	// 阻塞等待用户的输入
	//fmt.Scanf("%d", &a) // 注意&
	fmt.Scan(&a)
	fmt.Println("a = ", a)
}
