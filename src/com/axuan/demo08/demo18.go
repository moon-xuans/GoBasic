package main

import "fmt"

// 空接口

func xxx(arg ...interface{}) {

}

func main() {
	// 空接口万能类型，保存任何类型的值
	var i interface{} = 1
	fmt.Println("i = ", i)

	i = "abc"
	fmt.Println("i = ", i)
}
