package main

import (
	"fmt"
	"time"
)

func main()  {
	ch := make(chan int)

	// 创建匿名子go程
	go func() {
		for i:=0; i<6; i++ {
			ch <-i
			time.Sleep(time.Millisecond * 300)
		}
		close(ch)	// 所有数据全部写完，关闭channel
		//ch <- 888
	}()
/*	for {
		if data, has := <-ch; has == true {
			fmt.Println("主go程读到：", data)
		} else {
			fmt.Println("data=", data)
			break		// 停止读取
		}
	}*/

	for data := range ch {
		fmt.Println("data=", data)
	}
	fmt.Println("主go程将所有子go程数据读取完毕！！")
}
