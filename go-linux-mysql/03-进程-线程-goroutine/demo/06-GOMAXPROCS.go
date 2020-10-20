package main

import (
	"runtime"
	"fmt"
)

func main()  {
	n := runtime.GOMAXPROCS(3)
	fmt.Println("n = ", n)

	for {
		go fmt.Print("0")		// 子go程
		fmt.Print("1")			// 主go程
	}
}
