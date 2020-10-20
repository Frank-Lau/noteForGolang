package main

import (
	"net"
	"fmt"
	"os"
	"io"
)

func main()  {
	list := os.Args			// 提取命令行参数，存入 list 中
	if len(list) != 2 {		// 必须指定文件名
		fmt.Println("请按套路输入：格式为： go run xxx.go 文件名")
		return
	}
	filePath := list[1]		// 提取带路径的文件名

	conn, err := net.Dial("tcp", "127.0.0.1:8008")
	if err != nil {
		fmt.Println("Dial err:", err)
		return
	}
	defer conn.Close()

	// 使用 os.stat() 获取 文件名（不含路径）
	fileInfo, err := os.Stat(filePath)
	if err != nil {
		fmt.Println("Stat err:", err)
		return
	}
	fileName := fileInfo.Name()   // 文件名（不含路径）

	// 使用 conn 将文件名 write 给 server端（接收端）
	conn.Write([]byte(fileName))

	buf := make([]byte, 16)

	// 使用 conn 读取 server端（接收端）回发的 “ok”
	n, err := conn.Read(buf)
	if err != nil {
		fmt.Println("Read err:", err)
		return
	}
	if string(buf[:n]) == "ok" {
		// 封装函数 sendFile（带路径的文件名, conn）
		sendFile(filePath, conn)
	}
}

func sendFile(filePath string, conn net.Conn)  {
	// 只读打开文件（含路径）
	f, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Open err:", err)
		return
	}
	defer f.Close()
	buf := make([]byte, 4096)
	for {
		//  读取文件内容 ， 存到缓冲区中
		n, err := f.Read(buf)
		if err != nil {
			// 判断读到文件结尾 。结束文件读取 。 跳出循环。
			if err == io.EOF {
				fmt.Println("文件读取完毕")
				break
			}
			fmt.Println("Read err:", err)
			return
		}
		// 将缓冲区读到的数据 ，写给 server端（接收端） conn.write
		conn.Write(buf[:n])
	}
}
