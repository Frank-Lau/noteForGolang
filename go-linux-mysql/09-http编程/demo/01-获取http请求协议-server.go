package main

import (
	"net"
	"fmt"
	"os"
)

func errFunc(err error, errInfo string)  {
	if err != nil {
		fmt.Println(errInfo + " err:", err)
		os.Exit(-1)
	}
}

func main()  {
	// 创建监听套接字
	listener, err := net.Listen("tcp", "127.0.0.1:8000")
	errFunc(err, "Listen")
	defer listener.Close()

	// 阻塞等待客户端连接
	conn, err := listener.Accept()
	errFunc(err, "Accept")
	defer conn.Close()

	buf := make([]byte, 4096)
	for {
		n, err := conn.Read(buf)
		errFunc(err, "Read")
		fmt.Printf("#\n%s#\n", string(buf[:n]))
	}
}
