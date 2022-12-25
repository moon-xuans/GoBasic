package main

import "fmt"

// 不要操作没有合法指向的内存
func main() {
	var p *int
	p = nil
	fmt.Println("p = ", p)

	//*p = 666 // err，因为p没有合法指向

	var a int
	p = &a
	*p = 666
	fmt.Println("a = ", a)
}
