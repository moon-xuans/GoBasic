package main

import "fmt"

// for的使用
func main() {
	// go循环语句中只有for，没有while 以及 do while

	// for 初始条件; 判断条件; 条件变化{]
	// }

	// 1 + 2 + 3 + ... + 100 累加

	sum := 0

	// 1) 初始化条件 i:= 1
	// 2) 判断条件是否为真， i <= 100, 如果为真，执行循环体，如果为假，跳出循环
	// 3) 条件变化 i++
	// 4) 重复2, 3, 4
	for i := 0; i <= 100; i++ {
		sum += i
	}
	fmt.Println("sum = ", sum)
}
