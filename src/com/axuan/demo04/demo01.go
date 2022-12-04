package main

import "fmt"

// if语句
func main() {

	/*
		// if的使用
		s := "mike"

		// if和{就是条件，条件通常都是关系运算符
		if s == "mike" {
			fmt.Println("bingo")
		}

		// if支持1个初始化语句，初始化语句和判断条件以分号分割
		if a := 10; a == 10 {
			fmt.Println("a == 10")
		}

	*/

	/*
		// if和elseif 的使用
		// 1.
		a := 10
		if a == 10 {
			fmt.Println("a == 10")
		} else {
			fmt.Println("a != 10")
		}

		// 2.
		if a := 10; a == 10 {
			fmt.Println("a == 10")
		} else {
			fmt.Println("a != 10")
		}

		// 3.
		a = 8
		if a == 10 {
			fmt.Println("a == 10")
		} else if a > 10 {
			fmt.Println(" a > 10")
		} else if a < 10 {
			fmt.Println("a < 10")
		} else {
			fmt.Println("这是不可能的")
		}

		// 4.
		if a := 8; a == 10 {
			fmt.Println("a == 10")
		} else if a > 10 {
			fmt.Println(" a > 10")
		} else if a < 10 {
			fmt.Println("a < 10")
		} else {
			fmt.Println("这是不可能的")
		}
	*/

	// 多个if和elseif的区别
	// 这种好，因为如果满足一个条件后，后面的都不会继续去判断了
	a := 10
	if a == 10 {
		fmt.Println("a == 10")
	} else if a > 10 {
		fmt.Println("a > 10")
	} else if a < 10 {
		fmt.Println("a < 10")
	}

	b := 10
	if b == 10 {
		fmt.Println("b == 10")
	}
	if b > 10 {
		fmt.Println("b > 10")
	}
	if b < 10 {
		fmt.Println("b < 10")
	}
}
