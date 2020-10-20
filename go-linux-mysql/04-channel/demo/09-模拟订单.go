package main

import (
	"fmt"
	"strconv"
)

type OrderInfo struct {
	id int
	userName string
}

// 产生订单模块 —— 生产者
func makeOrder(out chan<- OrderInfo)  {
	for i:=0; i<10; i++ {
		order := OrderInfo{id:i+1, userName:"user"+ strconv.Itoa(i+1)}
		out <- order
	}
	close(out)
}

// 处理订单模块 —— 消费者
func dealOrder(in <-chan OrderInfo)  {
	for data := range in {
		fmt.Printf("处理掉 %s 的订单：%d\n",data.userName, data.id)
	}
}

func main()  {
	ch := make(chan OrderInfo)
	// 产生订单
	go makeOrder(ch)
	// 处理订单
	dealOrder(ch)
}
