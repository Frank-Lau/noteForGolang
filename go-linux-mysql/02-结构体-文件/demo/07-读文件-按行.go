package main

import (
	"os"
	"fmt"
	"bufio"
	"io"
)

func main()  {
	f, err := os.Open("C:/itcast/test2.txt")
	if err != nil {
		fmt.Println("Open err:", err)
		return
	}
	defer f.Close()

	// 1. 创建一个带有缓冲区的 reader
	reader := bufio.NewReader(f)

	// 创建一个 缓冲区，存储读到的数据。
	//buf := make([]byte, 4096)

	for {
		// 2. 按照 ‘\n’从 reader的缓冲区中 一行一行提取数据
		buf, err := reader.ReadBytes('\n')
		if err != nil && err == io.EOF{
			fmt.Println("read finish")
			break
		}
		fmt.Println(string(buf))
	}
}
