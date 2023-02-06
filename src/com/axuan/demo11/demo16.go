package main

import (
	"fmt"
	"time"
)

// 通过timer实现延时功能

func main() {
	<-time.After(2 * time.Second) // 定时2s，阻塞2s，2s后产生一个事件，往channel写内容
	fmt.Println("时间到")
}

func main162() {
	time.Sleep(2 * time.Second)
	fmt.Println("时间到")
}

func main161() {
	// 延时2s后打印一句话
	timer := time.NewTimer(2 * time.Second)
	<-timer.C
	fmt.Println("时间到")
}
