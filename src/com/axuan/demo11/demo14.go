package main

import "fmt"

// 单向channel的特点

func producer(out chan<- int) {
	for i := 0; i < 10; i++ {
		out <- i * i
	}
	close(out)
}

func consumer(in <-chan int) {
	for num := range in {
		fmt.Println("num = ", num)
	}
}

func main() {
	// 创建一个双向通道
	ch := make(chan int)

	// 生产者，生成数字，写入channel
	// 新开一个协程
	go producer(ch) // channel传参，引用传递

	// 消费者，从channel读取内容，打印
	consumer(ch)
}
