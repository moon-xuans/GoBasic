package main

import "fmt"

// 不定参数类型

// ...int类型这样的类型 ...type不定参数类型
// 注意：不定参数，一定(只能)放在形参中的最后一个参数
func MyFunc2(args ...int) { // 传递的实参可以是0或多个
	fmt.Println("len(args) = ", len(args)) // 获取用户传递参数的个数
	for i := 0; i < len(args); i++ {
		fmt.Printf("args[%d] = %d\n", i, args[i])
	}

	fmt.Println("----------------")

	// 返回2个值，第一个是下标，第二个是下标所对应的数
	for i, data := range args {
		fmt.Printf("args[%d] = %d\n", i, data)
	}
}

func main01() {
	MyFunc2()
	fmt.Println("++++++++++++++++++++++")
	MyFunc2(1)
	fmt.Println("++++++++++")
	MyFunc2(1, 2, 3)

}

// 固定参数一定要传参，不定参数根据需求传递
//func MyFunc3(a int, args ...int, b int) {
//
//}
//
//func main() {
//	MyFunc3(1)
//}
