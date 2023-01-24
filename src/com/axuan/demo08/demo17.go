package main

import "fmt"

// 接口转换

type Humaner17 interface { // 子集
	sayhi()
}

type Personer17 interface { // 超集
	Humaner17 // 匿名字段，继承了sayhi()
	sing(lrc string)
}

type Student17 struct {
	name string
	id   int
}

// Student实现了sayhi()
func (tmp *Student17) sayhi() {
	fmt.Printf("Student[%s, %d] sayhi\n", tmp.name, tmp.id)
}

func (tmp *Student17) sing(lrc string) {
	fmt.Println("Student在唱着：", lrc)
}

func main() {
	// 超集可以转换为子集，反过来不可以
	var iPro Personer17 // 超集
	iPro = &Student17{"mike", 666}

	var i Humaner17 // 子集

	// iPro = i // err
	i = iPro
	i.sayhi()
}
