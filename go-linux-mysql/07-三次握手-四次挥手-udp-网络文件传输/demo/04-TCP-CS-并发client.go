package main

import (
	"net"
	"fmt"
	"os"
)

func main()  {
	// 发起连接请求
	conn, err := net.Dial("tcp", "127.0.0.1:8007")
	if err != nil {
		fmt.Println("Dial err:", err)
		return
	}
	defer conn.Close()

	// 创建一个缓冲区，存储用户输入的 数据
	str := make([]byte, 4096)

	// 循环， 获取用户键盘输入
	go func() {
		for {
			//fmt.Scan()		不能使用，遇到空格、回车终止提取数据。
			n, err := os.Stdin.Read(str)
			if err != nil {
				fmt.Println("Stdin.Read err:", err)
				return
			}
			// 将读到的数据内容，原封不动的写给 服务器
			conn.Write(str[:n])
		}
	}()

	buf := make([]byte, 4096)
	//buf := make([]byte, 4)
	// 另一个go程（主go程）读取服务器回发数据内容，显示
	for {
		n, err := conn.Read(buf)
		if n == 0 {
			fmt.Println("服务器关闭！本端也关闭")
			return
		}
		if err != nil {
			fmt.Println("conn.Read err:", err)
			return
		}
		// 显示到 stdout
		fmt.Println("服务器发送数据：", string(buf[:n]))
	}
}
