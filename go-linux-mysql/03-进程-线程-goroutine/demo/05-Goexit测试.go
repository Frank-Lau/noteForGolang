package main

import (
	"fmt"
	"os"
)

func test()  {
	//runtime.Goexit()			// 退出当前 go程。
	//return
	os.Exit(0)			// 结束当前 进程。
	fmt.Println("dddddddddddddd")

}
func myGo()  {
	defer fmt.Println("cccccccccccccc")
	test()
	fmt.Println("kkkkkkkkkkkkkk")
}
func main()  {
	fmt.Println("aaaaaaaaaaaa")
	go myGo()		// 创建子go程
	fmt.Println("bbbbbbbbbbbb")
	for {
		;
	}
}
