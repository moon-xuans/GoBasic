package main

import (
	"fmt"
	"net"
	"strings"
	"time"
)

// 并发聊天服务器

type Client struct {
	C    chan string // 用户发送数据的管道
	Name string      // 用户名
	Addr string      // 网络地址
}

// 保存在线用户 cliAddr =>Client
var onlineMap map[string]Client

var message = make(chan string)

// 新开一个协程，转发消息，只要有消息来了，遍历map，给map每个成员都发送此消息
func Manager() {
	// 给map分配空间
	onlineMap = make(map[string]Client)

	for {
		msg := <-message // 没有消息前，这里会阻塞

		// 遍历map，给map每个成员都发送此消息
		for _, cli := range onlineMap {
			cli.C <- msg
		}
	}
}

func WriteMsgToClient(cli Client, conn net.Conn) {
	for msg := range cli.C { // 给当前客户端发送消息
		conn.Write([]byte(msg + "\n"))
	}
}

func MakeMsg(cli Client, msg string) (buf string) {
	buf = "[" + cli.Addr + "]" + cli.Name + ":" + msg
	return
}

// 处理用户连接
func HandleConnection(conn net.Conn) {
	defer conn.Close()

	// 获取客户端的网络地址
	cliAddr := conn.RemoteAddr().String()

	// 创建一个结构体，默认，用户名和网络地址一样
	cli := Client{make(chan string), cliAddr, cliAddr}
	// 把结构体添加到map
	onlineMap[cliAddr] = cli
	// 新开一个协程，专门给当前客户端发送消息
	go WriteMsgToClient(cli, conn)
	// 提示我是谁
	cli.C <- MakeMsg(cli, "I am here")

	// 广播某个在线
	message <- MakeMsg(cli, "login")

	isQuit := make(chan bool)  // 对方是否主动退出
	hasData := make(chan bool) // 是否有数据

	// 新建一个协程，接收用户发送过来的数据
	go func() {
		buf := make([]byte, 2048)
		for {
			n, err := conn.Read(buf)
			if n == 0 {
				isQuit <- true
				fmt.Println("conn.Read err = ", err)
				return
			}

			msg := string(buf[:n-1]) // 通过windows nc测试，多一个换行
			if len(msg) == 3 && msg == "who" {
				// 遍历map，给当前用户发送所有成员
				conn.Write([]byte("user list:\n"))
				for _, tmp := range onlineMap {
					msg = tmp.Addr + ":" + tmp.Name + "\n"
					conn.Write([]byte(msg))
				}
			} else if len(msg) >= 8 && msg[:6] == "rename" {
				name := strings.Split(msg, "|")[1]
				cli.Name = name
				onlineMap[cliAddr] = cli
				conn.Write([]byte("rename ok\n"))
			} else {
				// 转发此内容
				message <- MakeMsg(cli, msg)
			}

			hasData <- true
		}
	}()

	for {
		// 通过select检测channel的流动
		select {
		case <-isQuit:
			delete(onlineMap, cliAddr)        // 当前用户从map移除
			message <- MakeMsg(cli, "logout") // 广播谁下线了
			return
		case <-hasData:

		case <-time.After(10 * time.Second): // 10s后
			delete(onlineMap, cliAddr)                    // 当前用户从map移除
			message <- MakeMsg(cli, "time out leave out") // 广播谁下线了
			return
		}

	}
}

func main() {
	// 监听
	listener, err := net.Listen("tcp", ":8000")
	if err != nil {
		fmt.Println("net.Listen err = ", err)
		return
	}

	defer listener.Close()

	// 新开一个协程，转发消息，只要有消息来了，变量map，给map每个成员都发送此消息
	go Manager()

	// 主协程，循环阻塞等待用户连接
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("listener.Accept err = ", err)
			continue
		}

		go HandleConnection(conn) // 处理用户连接
	}
}
