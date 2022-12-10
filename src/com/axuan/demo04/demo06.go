package main

import "fmt"

// goto的使用
func main() {
	// break is not in a loop, switch, or select
	//break

	// continue is not in a loop
	//continue

	// goto可以用在任何地方，但是不能跨函数使用
	fmt.Println("1111111111111")

	// goto是关键字，End是用户起的名字，它叫标签
	goto End

	fmt.Println("2222222222222")
End:
	fmt.Println("33333333333333333")
}
