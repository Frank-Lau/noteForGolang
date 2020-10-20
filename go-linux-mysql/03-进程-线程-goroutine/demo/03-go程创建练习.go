package main

import (
	"fmt"
	"time"
)

func myGoroutine()  {
	for i:=0; i<3; i++ {
		fmt.Println(" I am 实名go程")
		time.Sleep(time.Millisecond * 200)
	}
}

func main()  {
	// 创建实名go程
	go myGoroutine()

	// 创建一个匿名go程
	go func() {
		for i:=0; i<3; i++ {
			fmt.Println(" ----------------I am 匿名go程")
			time.Sleep(time.Millisecond * 200)
		}
	}()

	// 主go程
	for i:=0; i<3; i++ {
		fmt.Println("======我是主go程 ")
		time.Sleep(time.Millisecond * 200)
	}

	for {
		;
	}
}
