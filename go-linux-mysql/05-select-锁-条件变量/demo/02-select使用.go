package main

import (
	"fmt"
	"time"
	"runtime"
)

func main()  {
	// 创建一个用于通信的channel
	ch := make(chan int)
	// 创建channel 用于结束 主、子go程通信
	quit := make(chan bool)
	go func() {			// 子go程监听channel上的读事件

		for {			// 大多数 select 监听是放到 for 中
			select {
			case num := <-ch:
				fmt.Println("num = ", num)
			case <-quit:
				//return 			// 返回当前函数
				runtime.Goexit()	// 结束go程
				//os.Exit(0)	// 结束进程
				//break			// break跳出当前case分支，而不是 for 循环。
				goto AAA
			}
			fmt.Println("-----------over------------")
		}
	AAA:
		fmt.Println("---------this is AAAA")
	}()

	// 主go程向channel中写数据
	for i:=0; i<7; i++ {
		ch <- i				// 写数据到 channel
		time.Sleep(time.Millisecond * 300)
	}
	quit <- true			// 主go程写数据完成。关闭。
}
