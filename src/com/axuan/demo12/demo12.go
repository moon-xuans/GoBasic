package main

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
)

// 并发版爬虫

// 爬取网页内容
func HttpGet(url string) (result string, err error) {
	resp, err1 := http.Get(url)
	if err1 != nil {
		err = err1
		return
	}

	defer resp.Body.Close()

	// 读取网页body内容
	buf := make([]byte, 1024*4)
	for {
		n, err := resp.Body.Read(buf)
		if n == 0 {
			fmt.Println("resp.Body.Read err = ", err)
			break
		}
		result += string(buf[:n])
	}
	return
}

func SpiderPage(i int, page chan int) {
	url := "https://tieba.baidu.com/f?kw=%E5%8E%9F%E7%A5%9E%E5%86%85%E9%AC%BC&ie=utf-8&pn=" +
		strconv.Itoa((i-1)*50)
	fmt.Printf("正在爬第%d页网页:%s\n", i, url)

	// 2) 爬(将所有的网站的内容全部爬下来)
	result, err := HttpGet(url)
	if err != nil {
		fmt.Println("HttpGet err = ", err)
		return
	}

	// 把内容写入到文件
	fileName := strconv.Itoa(i) + ".html"
	f, err1 := os.Create(fileName)
	if err1 != nil {
		fmt.Println("os.Create err1 = ", err1)
		return
	}

	f.WriteString(result) // 写内容

	f.Close() // 关闭文件

	page <- i
}

func DoWork(start, end int) {
	fmt.Printf("正在爬取 %d 到 %d的页面\n", start, end)
	page := make(chan int)

	// 明确目标 (要知道你准备在哪个范围或者网站取搜索)
	//https://tieba.baidu.com/f?kw=%E5%8E%9F%E7%A5%9E%E5%86%85%E9%AC%BC&ie=utf-8&pn=0 // 下一页+50
	for i := start; i <= end; i++ {
		go SpiderPage(i, page)
	}

	for i := start; i <= end; i++ {
		fmt.Printf("第%d个页面爬取完成\n", <-page)
	}
}

func main() {
	var start, end int
	fmt.Printf("请输入起始页( >= 1):")
	fmt.Scan(&start)
	fmt.Printf("请输入终止页( >= 起始页) :")
	fmt.Scan(&end)

	DoWork(start, end)
}
