package main

import (
	"fmt"
	"time"
)

// 有缓冲的channel

func main() {
	ch := make(chan int, 3)

	// len(ch)缓冲区剩余数据个数，cap(ch)缓冲区大小
	fmt.Printf("len(ch) = %d, cap(ch) = %d\n", len(ch), cap(ch))

	// 新建协程
	go func() {
		for i := 0; i < 10; i++ {
			ch <- i // 往chan写内容
			fmt.Printf("子协程[%d]:len(ch) = %d, cap(ch) = %d\n", i, len(ch), cap(ch))
		}
	}()

	// 延时
	time.Sleep(2 * time.Second)

	for i := 0; i < 10; i++ {
		num := <-ch // 读管道中内容，没有内容前，阻塞
		fmt.Println("num = ", num)
	}
}
