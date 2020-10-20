package main

import (
	"fmt"
)

func main()  {
	ch := make(chan int, 0)		// ch := make(chan int)  创建无缓冲channel
	fmt.Println("len=", len(ch), "cap=", cap(ch))

	// 创建用于指定主、子go程使用 stdout (标准输出) 的 channel
	ch2 := make(chan bool)

	go func() {
		for i:=0; i<5; i++ {
			ch <- i
			fmt.Println("子go程写：", i)
			// 子go程写完，向 ch2 中写数据。
			ch2 <- true
		}
		fmt.Println("子go程 finish")
	}()

	//time.Sleep(time.Second * 1)		// 主go程睡眠 5 秒。
	// 有多少次的写入、对应多少次的读取。读、写成对出现。
	for i:=0; i<5; i++ {
		num := <-ch
		// 从ch2 中激活读阻塞，打印读到的数据
		<- ch2
		fmt.Println("主go程读到：", num)
	}
}
