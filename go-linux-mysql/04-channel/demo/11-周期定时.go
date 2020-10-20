package main

import (
	"time"
	"fmt"
	"os"
)

func main()  {

	myTicker := time.NewTicker(time.Second*1)
	fmt.Println("1:", time.Now())

	i := 0
	for {
		data := <-myTicker.C
		fmt.Println("2:", data)
		i++
		if i == 3 {
			myTicker.Stop()  // 关闭 定时器
			//break
			//return
			//runtime.Goexit()
			os.Exit(0)
		}
	}
}
