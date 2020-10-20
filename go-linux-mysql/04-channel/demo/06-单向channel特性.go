package main

import "fmt"

func main()  {

	ch := make(chan int)		// 双向channel

	var ch1 chan<- int = ch		// 单向写channel
	ch1 <- 770
	//fmt.Println("<-ch1", <-ch1)

	var ch2 <-chan int = ch 	// 单向读channel
	num := <-ch2
	ch2 <- 88888
	fmt.Println("num=", num)
}
