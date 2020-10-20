package main

import (
	"net"
	"fmt"
	"strings"
)

// 专门用于客户端数据通信: read/write
func HandlerConnect(conn net.Conn)  {
	defer conn.Close()
	buf :=make([]byte, 4096)

	// 获取成功连接的客户端 IP+port
	fmt.Println("[", conn.RemoteAddr(), "]: 客户端连接成功！！" )

	// 服务器的每一个子go程，都能与对应客户端 反复通信
	for {
		// 读客户端发送数据
		n, err := conn.Read(buf)
		fmt.Println("buf=", buf[:n])

		if string(buf[:n]) == "exit\n" {
			fmt.Println("服务器收到客户端申请，关闭连接！")
			return
		}
		if n == 0 {
			fmt.Println("服务器检测到，对端粗暴关闭！本端对应go程退出，关闭连接")
			return
		}
		if err != nil {
			fmt.Println("Read err:", err)
			return
		}
		fmt.Println("服务器读到客户端发送：", string(buf[:n]))

		// 处理客户端发送的数据 （小 - 大）
		upperStr := strings.ToUpper(string(buf[:n]))

		// 回写数据给客户端
		conn.Write([]byte(upperStr))
	}
}

func main()  {
	// 创建监听套接字
	listener, err := net.Listen("tcp", "127.0.0.1:8007")
	if err != nil {
		fmt.Println("Listen err:", err)
		return
	}
	defer listener.Close()

	// 循环  阻塞监听客户端连接，返回一个 通信套接字
	for  {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Accept err:", err)
			return
		}
		// 创建 go程 专门用于客户端数据通信
		go HandlerConnect(conn)
	}
}
