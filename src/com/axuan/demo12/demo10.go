package main

import (
	"fmt"
	"net/http"
)

// http 服务器

// w，给客户端回复数据
// req，读取客户端发送的数据
func myHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("r.Method = ", r.Method)
	fmt.Println("r.URL = ", r.URL)
	fmt.Println("r.Header = ", r.Header)
	fmt.Println("r.Body = ", r.Body)

	fmt.Fprintln(w, "hello world") // 给客户端回复数据
}

func main() {
	// 注册处理函数，用户连接，自动调用指定的处理函数
	http.HandleFunc("/", myHandler)

	// 在指定的地址进行监听 开启一个HTTP
	http.ListenAndServe(":8000", nil)
}
