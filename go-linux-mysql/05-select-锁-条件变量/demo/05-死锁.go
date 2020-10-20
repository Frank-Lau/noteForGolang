package main

import "fmt"

// 死锁1
/*func main()  {
	ch := make(chan int)

	ch <- 888
	num := <-ch
	fmt.Println("num=", num)
}*/

// 死锁2
/*func main()  {
	ch := make(chan int)
	ch <- 888		// 阻塞

	go func() {
		num := <-ch
		fmt.Println("num=", num)
	}()
}*/

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	go func() {
		for {
			select {
			case num :=<-ch1:
				fmt.Println("num=", num)
				ch2 <- num
			}
		}
	}()

	for {
		select{
		case num := <-ch2:
			fmt.Println("num = ", num)
			ch1<-num
		}
	}
}