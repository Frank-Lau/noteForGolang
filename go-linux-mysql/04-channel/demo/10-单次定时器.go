package main

import (
	"time"
	"fmt"
)

func main()  {

	// 创建定时器、指定定时时长。
	myTimer := time.NewTimer(time.Second * 3)  // 3秒计时到达后系统向当前timer 的 C 写 当前时间。

	fmt.Println("1:", time.Now())
	for {
		// 从定时器读取数据。 无数据写入阻塞。
		data := <-myTimer.C
		fmt.Println("2:", data)
	}
}

/*
func main()  {
	fmt.Println("now:", time.Now())
	data := <-time.After(time.Second * 3)			// sleep()
	fmt.Println("data:", data)
}*/
// 定时器停止、重置
/*func main()  {
	// 创建定时器、指定定时时长。
	myTimer := time.NewTimer(time.Second * 3)  // 3秒计时到达后系统向当前timer 的 C 写 当前时间。
	fmt.Println("1:", time.Now())
	fmt.Println("----- 定时完成 ----")

	myTimer.Reset(time.Second)

	//myTimer.Stop()  // 关闭 myTimer 的 C

	// 从定时器读取数据。 无数据写入阻塞。
	data := <-myTimer.C

	fmt.Println("2:", data)
}*/
