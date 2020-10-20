package main

import (
	"net"
	"fmt"
	"time"
)

func main()  {
	// 建立用于数据通信的套接字 conn
	//conn, err := net.Dial("tcp", "192.168.42.51:8004")
	conn, err := net.Dial("udp", "127.0.0.1:8005")
	if err != nil {
		fmt.Println("Dial err:", err)
		return
	}
	defer conn.Close()

	num := 0

	// 循环发送数据给 服务器
	for {
		// 主动发送数据给服务器
		n, err := conn.Write([]byte("hello I'm client"))
		if err != nil {
			fmt.Println("Write err:", err)
			return
		}
		fmt.Println("客户端成功写出 ", n, " 字节")
		time.Sleep(time.Second)
		num++

		if num == 5 {
			conn.Write([]byte("exit"))
		}
	}
}
