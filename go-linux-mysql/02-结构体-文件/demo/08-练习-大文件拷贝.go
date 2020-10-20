package main

import (
	"os"
	"fmt"
	"io"
)

func main()  {
	// 打开源文件
	f_r, err := os.Open("C:/itcast/10-open打开文件.avi")
	if err != nil {
		fmt.Println("open err:", err)
		return
	}
	defer f_r.Close()
	// 创建目标文件
	f_w, err := os.Create("./test.avi")
	if err != nil {
		fmt.Println("Create err:", err)
		return
	}
	defer f_w.Close()
	// 创建缓冲区
	buf := make([]byte, 4096)
	// 循环读取，写入
	for {
		n, err := f_r.Read(buf)
		if err != nil && err == io.EOF {
			break
		}
		f_w.Write(buf[:n])		// 读多少、写多少
	}
	fmt.Println("successful !!!")
}
