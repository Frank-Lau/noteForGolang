package main

import (
	"fmt"
	"time"
)

func main()  {
	// 创建一个用于通信 channel
	ch1 := make(chan string)

	// 子go程，循环打印i。写数据给 主go程
	go func() {
		for i:=0; i<2; i++ {
			fmt.Println("子go程， i=", i)
			time.Sleep(time.Second)
		}
		ch1 <- "子go程 循环打印2次完成"
	}()

	// 主go程 从 channel中 读取数据
	str := <- ch1
	fmt.Println("主go程，读到：", str)
}
