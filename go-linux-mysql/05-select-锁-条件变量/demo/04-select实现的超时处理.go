package main

import (
	"fmt"
	"time"
)

func main()  {
	ch := make(chan int)		// 通信
	quit := make(chan bool) 	// 协调主、子go程退出

	go func() {
		for {
			select {
			case num := <-ch:			// 监听ch 上数据流动，一旦有数据，此case满足，会重置下面定时器
				fmt.Println("num = ", num)
			case <-time.After(time.Second * 3):		// 监听是否计时满 3 秒
				fmt.Println("子go程超时，退出")
				quit <- true			// 子go程运行结束，通知主go程退出
				return
			}
		}
	}()

	for i:=0; i<2; i++ {
		ch <- i			// 写数据到 ch 上，会导致 定时器重置。
		time.Sleep(time.Second * 4)
	}

	<-quit				// 阻塞！直到读到子go程写入的数据才解除阻塞
}
