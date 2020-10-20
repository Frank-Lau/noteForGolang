package main

import (
	"net"
	"fmt"
)

func main()  {
	fmt.Println("--------------------1")
	// 创建监听套接字
	//listener, err := net.Listen("tcp", "127.0.0.1:8004")
	listener, err := net.Listen("tcp", "192.168.42.51:8004")
	if err != nil {
		fmt.Println("Listen err:", err)
		return
	}
	// 程序退出前，关闭监听套接字
	defer listener.Close()
	fmt.Println("--------------------2")
	// 创建监听	——  造成当前 server 阻塞等待。 一旦解除阻塞，说明成功与客户端建立连接。
	conn, err := listener.Accept()		// 返回用于通信的套接字
	if err != nil {
		fmt.Println("Accept err:", err)
		return
	}
	defer conn.Close()
	fmt.Println("--------------------3")

	// 创建缓冲区，保存服务器读到的数据
	buf := make([]byte, 4096)

	for {
		// 与客户端数据通信： 读、写。
		n, err := conn.Read(buf)
		if err != nil {
			fmt.Println("Read err:", err)
			return
		}
		fmt.Println("服务器读到客户端发送：", string(buf[:n]))
	}
}
