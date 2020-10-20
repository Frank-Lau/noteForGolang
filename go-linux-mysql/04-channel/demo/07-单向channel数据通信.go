package main

import "fmt"

// 发送数据 —— 写
func send(out chan<- int)  {
	out <- 890
	close(out)
}

// 接收数据 —— 读
func recv(in <-chan int)  {
	num := <-in
	fmt.Println("num = ", num)
}

func main()  {
	ch := make(chan int) 	// 双向channel
	go send(ch)			// 实参赋值给形参：   双向channel 赋值给 单向 写channel
	recv(ch)
}
