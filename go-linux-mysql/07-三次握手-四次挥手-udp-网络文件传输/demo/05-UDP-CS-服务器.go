package main

import (
	"net"
	"fmt"
)

func main()  {
	// 创建 UDP地址结构
	udpAddr, err := net.ResolveUDPAddr("udp", "127.0.0.1:8005")
	if err != nil {
		fmt.Println("ResolveUDPAddr err:", err)
		return
	}
	// 创建用于通信的 socket
	udpConn, err := net.ListenUDP("udp", udpAddr)
	if err != nil {
		fmt.Println("ListenUDP err:", err)
		return
	}
	defer udpConn.Close()

	buf := make([]byte, 4096)

	for {
		// 读 对端发送的数据内容
		n, cltAddr, err := udpConn.ReadFromUDP(buf)
		if err != nil {
			fmt.Println("ReadFromUDP err:", err)
			return
		}
		go func() {

			fmt.Println("fwq 读到 客户端发送： ", string(buf[:n]))
			// 写 回数据给对端
			udpConn.WriteToUDP([]byte("hello I am UDP server\n"), cltAddr)
		}()
	}
}
