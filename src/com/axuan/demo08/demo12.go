package main

import "fmt"

// 方法的重写

type Person12 struct {
	name string // 名字
	sex  byte   // 性别，字符类型
	age  int    // 年龄
}

// Person类型，实现了一个方法
func (tmp *Person12) PrintInfo() {
	fmt.Printf("name=%s, sex=%c, age=%d\n", tmp.name, tmp.sex, tmp.age)
}

type Student12 struct {
	Person12 // 匿名字段
	id       int
	addr     string
}

// Student也实现了一个方法，这个方法和Person方法同名，这种方法叫重写
func (tmp *Student12) PrintInfo() {
	fmt.Println("Student: tmp = ", tmp)
}

func main() {
	s := Student12{Person12{"mike", 'm', 18}, 666, "bj"}
	// 就近原则：先找本作用域的方法，找不到再用继承的方法
	s.PrintInfo() // 到底调用的是Person，还是Student

}
