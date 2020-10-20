package main

import (
	"net"
	"fmt"
)

func main()  {
	conn, err := net.Dial("tcp", "127.0.0.1:8000")
	if err != nil {
		fmt.Println("Dial err:", err)
		return
	}
	defer conn.Close()

	// 冒充浏览器组织一个最简单的 http请求协议。
	resquest := "GET /icast.txt HTTP/1.1\r\nHost:127.0.0.1:8000\r\n\r\n"
	// 将http请求协议。写给web服务器
	conn.Write([]byte(resquest))

	buf := make([]byte, 4096)
	for {
		n, err := conn.Read(buf)
		if err != nil {
			fmt.Println("Read err:", err)
			return
		}
		fmt.Printf("#\n%s#\n", string(buf[:n]))
	}
}
