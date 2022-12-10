package main

import (
	"fmt"
	"time"
)

// break和continue的区别
func main() {
	i := 0

	// for后面不写任何东西，这个循环条件永远为真，死循环
	for {
		i++
		time.Sleep(time.Second) // 延迟1s

		if i == 5 {
			// 跳出循环，如果嵌套多个循环，跳出最近的那个内循环
			//break

			// 跳出本次循环，下一次继续
			continue
		}
		fmt.Println("i = ", i)
	}
}
