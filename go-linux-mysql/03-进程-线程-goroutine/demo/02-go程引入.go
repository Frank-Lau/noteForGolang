package main

import (
	"fmt"
	"time"
)

func singing()  {
	for i:=0; i<5; i++ {
		fmt.Println("-----正在唱歌：人猿泰山-----")
		time.Sleep(time.Millisecond * 200)
	}
}

func dancing()  {
	for i:=0; i<5; i++ {
		fmt.Println("-----正在跳舞：赵四街舞-----")
		time.Sleep(time.Millisecond * 200)
	}
}

func main()  {
	go singing()		// 子go程 1
	go dancing()		// 子go程 2

/*	for {			// 防止主go程提前结束
		;
	}*/
}
