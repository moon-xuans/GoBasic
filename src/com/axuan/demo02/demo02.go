package main

// 常量
func main() {
	/*
		// 常量的使用
		// 变量：程序运行期间，可以改变的量，变量声明需要var
		// 常量：程序运行期间，不可以改变的量，常量声明需要const
		const a int = 10
		// a = 20 // err，常量不允许修改
		fmt.Println("a = ", a)

		const b = 11.2 // 可以省略类型，go会自动推导，并且没有使用:=
		fmt.Println("b = ", b)

	*/

	/*

		// 多个变量或常量定义
		var (
			a int     = 1
			b float64 = 2.0
		)
		fmt.Println("a = ", a, "b = ", b)
		const (
			i int     = 1
			j float64 = 3.14
		)
		fmt.Println("i = ", i, "j = ", j)
		// 当然这两个都可以自动推导类型
		var (
			a = 1
			b = 2.0
		)
		const (
			i = 1
			j = 3.14
		)
		fmt.Println("a = ", a, "b = ", b)
		fmt.Println("i = ", i, "j = ", j)

	*/

	/*
		// iota枚举

		// 1.iota常量自动生成器，每个一行，自动累加1
		// 2.iota给常量赋值使用
		const (
			a = iota
			b = iota
			c = iota
		)
		fmt.Printf("a = %d, b = %d, c = %d \n", a, b, c)

		// 3.iota遇到const，重置为0
		const d = iota
		fmt.Printf("d = %d", d)

		// 4.可以只写一个iota
		const (
			a1 = iota
			b1
			c1
		)
		fmt.Printf("a1 = %d, b1 = %d, c1 = %d\n", a1, b1, c1)

		// 5.如果是同一行，值都一样
		const (
			i          = iota
			j1, j2, j3 = iota, iota, iota
			k          = iota
		)
		fmt.Printf("i = %d, j1 = %d, j2 = %d, j3 = %d, k = %d", i, j1, j2, j3, k)
	*/
}
