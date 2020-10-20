package main

import (
	"fmt"
	"time"
)

// 用于同步的channel
var mych = make(chan int)

// 创建一个打印机函数
func printer(str string)  {
	for _, ch := range str {
		fmt.Printf("%c", ch)
		time.Sleep(time.Millisecond * 300)
	}
}
// 使用打印机的go程1	先： 写
func Person1()  {

	printer("hello")
	mych <- 789

}
// 使用打印机的go程2	后： 读
func Person2()  {
	// 读channel 。 如果写端没有写，会阻塞当前 person2 go程
	<- mych
	printer("world")
}
func main()  {
	// 创建 两个go程分别使用打印机
	go Person1()
	go Person2()
	for {		// 防止主go程先于子go程提前结束。
		;
	}
}
