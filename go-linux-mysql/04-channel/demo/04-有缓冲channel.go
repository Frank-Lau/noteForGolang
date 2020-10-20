package main

import (
	"fmt"
)

func main()  {
	// 创建一个有 3 个元素缓冲区的 channel
	ch := make(chan int, 3)
	fmt.Println("len=", len(ch), "cap=", cap(ch))

	go func() {
		for i:=0; i<5; i++ {
			ch <- i
			fmt.Println("子go程写：", i, "len=", len(ch), "cap=", cap(ch))
		}
		fmt.Println("子go程 finish")
	}()

	//time.Sleep(time.Second * 2)

	for i:=0; i<5; i++ {
		num := <-ch
		fmt.Println("主读到：", num)
	}
}
