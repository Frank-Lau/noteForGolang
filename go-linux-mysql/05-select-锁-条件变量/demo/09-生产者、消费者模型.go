package main

import (
	"math/rand"
	"time"
	"fmt"
	"runtime"
	"sync"
)
// 创建条件变量
var cond sync.Cond

func producer(send chan<- int, i int)  {
	for {
		// 给条件变量加锁
		cond.L.Lock()
		// 判断条件变量是否满足
		//if len(send) == 5 {
		for len(send) == 5 {
			cond.Wait()  // 3 个作用
		}
		// 产生数据
		data := rand.Intn(1000)
		// 访问公共区
		send <- data
		fmt.Printf("--%d--写出：%d\n", i, data)
		// 唤醒对端（消费者）
		cond.Signal()
		// 解锁
		cond.L.Unlock()

		time.Sleep(time.Millisecond * 500)
	}
}
func consumer(recv <-chan int, i int)  {
	for {
		cond.L.Lock()
		//if len(recv) == 0 {
		for len(recv) == 0 {
			cond.Wait()
		}
		fmt.Printf("%d读到：%d\n", i,  <-recv)
		cond.Signal()
		cond.L.Unlock()
		time.Sleep(time.Second)
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())
	ch := make(chan int, 5)

	// 创建绑定给条件变量的 锁
	cond.L = new(sync.Mutex)

	for i:=0; i<5; i++ {
		go producer(ch, i+1)
	}

	for i:=0; i<3; i++ {
		go consumer(ch, i+1)
	}

	for {
		runtime.GC()
	}
}
