package main

import "fmt"

func producer(send chan<- int)  {
	// 生产者模拟生产 10 次数据，写入 缓冲区中
	for i:=0; i<10; i++ {
		send <- i*i
	}
	close(send)
}

func consumer(recv <-chan int)  {
	// 消费者模拟从 公共区 取走数据。
	for data := range recv {
		fmt.Println("消费者读走：", data)
	}
}

func main() {
	// 创建 缓冲区, 带有容量的 channel
	ch := make(chan int, 6)
	// 促使生产者产生数据
	go producer(ch)
	// 消费者处理数据
	consumer(ch)
}
