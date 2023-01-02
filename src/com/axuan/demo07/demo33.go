package main

import "fmt"

// 结构体比较和赋值

type Student5 struct {
	id   int
	name string
	sex  byte // 字符类型
	age  int
	addr string
}

func main() {
	s1 := Student5{1, "mike", 'm', 18, "bj"}
	s2 := Student5{1, "mike", 'm', 18, "bj"}
	s3 := Student5{2, "mike", 'm', 18, "bj"}
	fmt.Println("s1 == s2 ", s1 == s2)
	fmt.Println("s1 == s3", s1 == s3)

	// 同类型的2个结构体变量可以相互赋值
	var tmp Student5
	tmp = s3
	fmt.Println("tmp = ", tmp)
}
