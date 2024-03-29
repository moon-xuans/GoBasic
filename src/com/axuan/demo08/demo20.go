package main

import "fmt"

// 类型断言：switch.go

type Student20 struct {
	name string
	id   int
}

func main() {
	i := make([]interface{}, 3)
	i[0] = 1                      // int
	i[1] = "hello go"             // string
	i[2] = Student20{"mike", 666} // Student

	// 类型查询，类型断言
	for index, data := range i {
		switch value := data.(type) {
		case int:
			fmt.Printf("x[%d] 类型为int， 内容为%d\n", index, value)
		case string:
			fmt.Printf("x[%d] 类型为string， 内容为%s\n", index, value)
		case Student20:
			fmt.Printf("x[%d] 类型为Student， 内容为name = %s, id = %d\n", index, value.name, value.id)
		}
	}
}
