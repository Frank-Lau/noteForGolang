package main

import (
	"net"
	"fmt"
	"os"
	"io"
)

func main()  {
	// 创建监听套接字  net.Listen()  --- listener
	listener, err := net.Listen("tcp", "127.0.0.1:8008")
	if err != nil {
		fmt.Println("Listen err:", err)
		return
	}
	defer listener.Close()

	//  阻塞监听 发送端连接请求  listener.Accept()  --- conn
	conn, err := listener.Accept()
	if err != nil {
		fmt.Println("Listen err:", err)
		return
	}
	defer conn.Close()

	buf := make([]byte, 4096)

	// 读取 发送端发送的文件名（不含路径）conn.read . 保存文件名
	n, err := conn.Read(buf)
	if err != nil {
		fmt.Println("Read err:", err)
		return
	}
	fileName := string(buf[:n])

	//  回发的 “ok”给发送端
	conn.Write([]byte("ok"))

	// 封装函数 recvFile（文件名，conn）
	recvFile(fileName, conn)
}

func recvFile(fileName string, conn net.Conn)  {
	// 创建新文件名  os.Create()
	filePath := "C:/itcast/test2/" + fileName
	//f, err := os.Create(fileName)  // 当前目录位置
	f, err := os.Create(filePath)  // 指定目录位置
	if err != nil {
		fmt.Println("Create err:", err)
		return
	}
	defer f.Close()

	buf := make([]byte, 4096)
	for {
		n, err := conn.Read(buf)
		if n == 0 {
			fmt.Println("下载文件完毕")
			return
		}
		if err != nil && err != io.EOF {
			fmt.Println("Read err:", err)
			break
		}
		// 写到 本地文件中
		f.Write(buf[:n])
	}
}
