package main

import "fmt"

func fibonacci(ch <-chan int, quit <-chan string)  {
	for {
		select {
		case num := <-ch:
			fmt.Println("num = ", num)
		case str := <-quit:
			fmt.Println("子go程收到指令，结束！str=", str)
			return  // Goexit()
		}
	}
}
func main()  {
	ch := make(chan int)		// 通信
	quit := make(chan string)	// 退出
	go fibonacci(ch, quit)

	x, y := 1, 1		// 指定fibonacci前2个元素。
	for i:=0; i<80; i++ {		// 找寻fibonacci数列规律
		ch <- x
		x, y = y, x+y
	}
	// 通知子go程结束
	quit <- "stop"

	//lfkjsdlkfjsd
	//jfklsdjflsd
	//fjkdlsjfdsf
}
