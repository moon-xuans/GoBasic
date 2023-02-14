package main

import (
	"fmt"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"strings"
)

// 段子爬虫

func HttpGets(url string) (result string, err error) {
	resp, err1 := http.Get(url)
	if err1 != nil {
		err = err1
		return
	}

	defer resp.Body.Close()

	// 读取网页body内容
	buf := make([]byte, 1024*4)
	for {
		n, _ := resp.Body.Read(buf)
		if n == 0 {
			break
		}
		result += string(buf[:n])
	}
	return
}

// SpiderOne 开始爬取每一个笑话，每一个段子 title, content, err := SpiderOne(url)
func SpiderOne(url string) (title, content string, err error) {
	// 开始爬取页面内容
	result, err1 := HttpGets(url)
	if err1 != nil {
		err = err1
		return
	}

	// 取关键信息
	// 取标题 <h1>标题</h1>只取1个
	re1 := regexp.MustCompile(`<h1>(?s:(.*?))</h1>`)
	if re1 == nil {
		err = fmt.Errorf("%s", "regexp.MustCompile err")
		return
	}
	// 取标题
	tmpTitle := re1.FindAllStringSubmatch(result, 1) // 最后一个参数为1，只过滤第一个
	for _, data := range tmpTitle {
		title = data[1]
		title = strings.Replace(title, "\t", "", -1)
		break
	}

	re2 := regexp.MustCompile(`<div class = "content-txt pt10>(?s:(.*?))<a id = "prev> href = "`)
	if re2 == nil {
		err = fmt.Errorf("%s", "regexp.MustCompile err2")
		return
	}

	// 取内容
	tmpContent := re2.FindAllStringSubmatch(result, -1)
	for _, data := range tmpContent {
		content = data[1]
		content = strings.Replace(content, "\t", "", -1)
		content = strings.Replace(content, "\n", "", -1)
		content = strings.Replace(content, "\r", "", -1)
		content = strings.Replace(content, "<br />", "", -1)
		break
	}

	return
}

// 把内容写入到文件
func SaveFile(i int, fileTitle []string, fileContent []string) {
	// 新建文件
	f, err := os.Create(strconv.Itoa(i) + ".txt")
	if err != nil {
		fmt.Println("os.Create err = ", err)
		return
	}

	defer f.Close()

	// 写内容
	n := len(fileTitle)
	for i := 0; i < n; i++ {
		// 写标题
		f.WriteString(fileTitle[i])
		f.WriteString("\n")
		// 写内容
		f.WriteString(fileContent[i])
		f.WriteString("\n")

		f.WriteString("\n================\n")
	}
}

func Spider(i int, page chan int) {
	// 明确爬取的url
	// https://www.pengfu.com/xiaohua_1.html
	url := "https://www.pengfu.com/xiaohua_" + strconv.Itoa(i) + ".html"
	fmt.Printf("正在爬取第%d个网页:%s\n", i, url)

	// 开始爬取页面内容
	result, err := HttpGets(url)
	if err != nil {
		fmt.Println("HttpGet err = ", err)
		return
	}
	//fmt.Println("r = ", result)
	//取,<h1 class="bp-b"><a href = "一个段子url链接">
	// 解释表达式
	re := regexp.MustCompile(`<h1 class = "dp-b"><a href="(?s:(.*?))">`)
	if re == nil {
		fmt.Println("regexp.MustCompile err")
		return
	}

	// 取关键信息
	joyUrls := re.FindAllStringSubmatch(result, -1)

	fileTitle := make([]string, 0)
	fileContent := make([]string, 0)

	// 取网址
	// 第一个返回下标，第二个返回内容
	for _, data := range joyUrls {
		//fmt.Println("url = ", data)

		// 开始爬取每一个笑话，每一个段子
		title, content, err := SpiderOne(data[1])
		if err != nil {
			fmt.Println("SpiderOne err = ", err)
			continue
		}
		//fmt.Printf("title = #%s#", title)
		//fmt.Printf("content = #%s#", content)

		fileTitle = append(fileTitle, title)
		fileContent = append(fileContent, content)
	}

	// 把内容写入到文件
	SaveFile(i, fileTitle, fileContent)
	
	page <- i
}

func Work(start, end int) {
	fmt.Printf("准备爬取第%d页到%d页的网址\n", start, end)

	page := make(chan int)

	for i := start; i <= end; i++ {
		// 定义一个函数，爬主页面
		go Spider(i, page)
	}

	for i := start; i <= end; i++ {
		fmt.Printf("第%d个页面爬取成功\n", <-page)
	}
}

func main() {
	var start, end int
	fmt.Printf("请输入起始页( >= 1):")
	fmt.Scan(&start)
	fmt.Printf("请输入终止页( >= 起始页) :")
	fmt.Scan(&end)

	Work(start, end)

}
