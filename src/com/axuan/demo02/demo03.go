package main

// 基本数据类型

func main() {
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

	/*
		// 字符和字符串的区别
		var ch byte
		var str string

		// 字符
		// 1.单引号
		// 2.字符，往往都只有一个字符，转义字符除外'\n'
		ch = 'a'
		fmt.Println("ch = ", ch)

		// 字符串
		// 1.双引号
		// 2.字符串由1个或多个字符组成
		// 3.字符串都是隐藏了一个结果符,'\0'
		str = "a" // 由'a'和'\0'组成了一个字符串
		fmt.Println("str = ", str)

		str = "hello go"
		// 只想操作字符串的某个字符，从0开始操作
		fmt.Printf("str[0] = %c, str[1] = %c\n", str[0], str[1])

	*/

	/*
		// 复数类型
		var t complex128 // 声明
		t = 2.1 + 3.14i  // 赋值
		fmt.Println("t = ", t)

		// 自动推导类型
		t2 := 3.3 + 4.4i
		fmt.Printf("t2 type is %T\n", t2)

		// 通过内建函数，取实部和虚部
		fmt.Println("real(t2) = ", real(t2), ", imag(t2) = ", imag(t2))

	*/
}
