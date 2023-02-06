package main

import (
	"fmt"
	"time"
)

// 停止和重置定时器

func main() {
	timer := time.NewTimer(3 * time.Second)
	ok := timer.Reset(1 * time.Second)
	fmt.Println("ok = ", ok)

	<-timer.C
	fmt.Println("时间到")
}

func main170() {
	timer := time.NewTimer(3 * time.Second)

	go func() {
		<-timer.C
		fmt.Println("子协程可以打印了，因为定时器的时间到")
	}()

	timer.Stop() // 停止定时器

	for {

	}
}
