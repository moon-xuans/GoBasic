package main // 要生成Go可执行程序，必须建立一个名字为main的包，并且在该包中包含一个main()的函数

func test() (a, b, c int) {
	return 1, 2, 3
}

func main() { // 强制要求左花括号{
	/*
		go build ... 只会编译代码，不能运行可执行文件
		go run ... 只会运行，不会生成可执行文件
	*/

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

	/*
		// bool类型
		// 1.声明变量，没有初始化，零值(初始化)为false
		var a bool
		fmt.Println("a = ", a)

		a = true
		fmt.Println("a = ", a)

		// 2.自动推导类型
		var b = false
		fmt.Println("b = ", b)

		c := false
		fmt.Println("c = ", c)
	*/

	/*
		// 浮点型

		// 声明变量
		var f1 float32
		f1 = 3.14
		fmt.Println("f1 = ", f1)

		// 自动推导类型
		f2 := 3.14
		fmt.Printf("f2 type is %T \n", f2) // f2 type is float64

		// float64存储小数比float32更准确
	*/

	/*
		// 字符类型
		var ch byte // 声明字符类型
		ch = 97
		//fmt.Println("ch = ", ch)
		// 格式化输出，%c以字符方式打印，%d以整型方式打印
		fmt.Printf("ch = %c \n", ch)

		ch = 'a' // 字符，单引号
		fmt.Printf("%c %d \n", ch, ch)

		// 大写转小写，小写转大写，大小写相差32，小写大
		fmt.Printf("大写：%d, 小写: %d\n", 'A', 'a')
		fmt.Printf("大写转小写: %c\n", 'A'+32)
		fmt.Printf("小写转大写: %c\n", 'a'-32)

		// '\'以反斜杠开头的字符是转义字符，'\n'代表换行
		fmt.Printf("hello go %c", '\n')
		fmt.Printf("hello moon")

	*/

	/*
		// 字符串类型
		var str1 string // 声明变量
		str1 = "abc"
		fmt.Println("str1 = ", str1)

		// 自动推导类型
		str2 := "mike"
		fmt.Printf("str2类型是%T\n", str2)

		// 内建函数，len()可以测字符串的长度，有多少个字符
		fmt.Println("len(str2) = ", len(str2))

	*/
}
