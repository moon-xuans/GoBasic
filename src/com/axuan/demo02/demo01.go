package main

func test() (a, b, c int) {
	return 1, 2, 3
}

//  变量

func main() {
	/*
		// 自动推导类型
		var a int              // 命名时，一个名字必须以一个字母或下划线开头
		fmt.Println("a = ", a) // go语言每个语句后面不要求加上分号表示语句结束
		a = 10
		fmt.Println("a = ", a)
		var b int = 10
		fmt.Println("b = ", b)
		// 对于声明变量时需要进行初始化，var关键字可以保留，但不再是必要的元素
		// 自动推导类型，必须初始化
		c := 30
		// %T打印变量所属的类型
		fmt.Printf("type is %T", c)


	*/
	/*
		// 出现在 := 左侧的变量不应该是已经被声明过的， := 定义时必须初始化
		var v4 int
		v4 := 2 // err
	*/

	/*
		// printf和println的区别
		num1, num2, num3 := 1, 2, 3
		fmt.Println("num1 = ", num1, "num2 = ", num2, "num3 = ", num3)
		fmt.Printf("num1 = %d num2 = %d num3 = %d \n", num1, num2, num3)
	*/

	/*
		// 多重赋值和匿名变量
		a, b := 10, 20
		var temp int
		temp = a
		a = b
		b = temp
		fmt.Printf("a = %d, b = %d \n", a, b)

		num1, num2 := 10, 20
		num1, num2 = num2, num1
		fmt.Printf("num1 = %d, num2 = %d \n", num1, num2)

		i, j := 10, 20
		fmt.Printf("i = %d, j = %d \n", i, j)
		i = 10
		j = 20

		// _匿名变量，丢弃数据不处理，_匿名变量配合函数返回值使用，才有优势
		var tmp int
		tmp, _ = i, j // 这样写，可以直接tmp = j即可，一般情况下匿名变量，是搭配函数使用的
		fmt.Println("tmp = ", tmp)

		var c, d, e int
		c, d, e = test()
		fmt.Printf("c = %d, d = %d, e = %d\n", c, d, e)
		_, d, e = test()
		fmt.Printf("d = %d, e = %d \n", d, e)

	*/

}
