package main

import (
	"fmt"
	"regexp"
)

// 正则表达式3

func main() {
	buf := `
	dsfasdgsadga
	<div>哈哈</div>
	<div>test</div>
	<div>
		测试
		你好啊
	</div>
dsagsadgasd
`
	// 解释正则表达式，+匹配前一个字符的1次或多次
	reg := regexp.MustCompile(`<div>(?s:(.*?))</div>`)
	if reg == nil {
		fmt.Println("MustCompile err")
		return
	}

	// 提取关键信息
	//result := reg.FindAllString(buf, -1)
	result := reg.FindAllStringSubmatch(buf, -1)

	// 过滤<></>
	for index, text := range result {
		fmt.Printf("text[%d] = %s\n", index, text[1])
	}
}
