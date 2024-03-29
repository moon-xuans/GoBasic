package main

import (
	"encoding/json"
	"fmt"
)

// json解析到结构体

type IT08 struct {
	Company  string   `json: "company"` // 此字段不会输出到屏幕
	Subjects []string `json:"subjects"` // 二次编码
	IsOk     bool     `json:"isok"`
	Price    float64  `json:"price"`
}

func main() {
	jsonBuf := ` {
 "company": "itcast",
 "subjects": [
  "Go",
  "C++",
  "Python",
  "Test"
 ],
 "isok": true,
 "price": 666.666
}
`
	var tmp IT08                                 // 定义一个结构体变量
	err := json.Unmarshal([]byte(jsonBuf), &tmp) // 第二个参数要地址传递
	if err != nil {
		fmt.Println("err = ", err)
		return
	}
	//fmt.Println("tmp = ", tmp)
	fmt.Printf("tmp = %+v\n", tmp)

	type IT2 struct {
		Subjects []string `json:"subjects"` // 二次编码
	}
	var tmp2 IT2
	err = json.Unmarshal([]byte(jsonBuf), &tmp2) // 第二个参数要地址传递
	if err != nil {
		fmt.Println("err = ", err)
		return
	}
	fmt.Printf("tmp2 = %+v\n", tmp2)
}
