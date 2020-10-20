package main

import (
	"runtime"
	"fmt"
	"math/rand"
	"time"
	"sync"
)

var num int 		// 初值 0
var rwlock sync.RWMutex		// 创建 读写锁： 一把锁。两种模式：w、r

// 每个读go程，读取公共区数据，访问之前以读模式加锁，访问结束解锁。
func readGo(idx int)  {
	for {
		rwlock.RLock() // 访问公共区之前，加读属性锁
		fmt.Printf("第%d个 读go程，读到：%d\n", idx, num)
		rwlock.RUnlock()
		time.Sleep(time.Millisecond * 500)
	}
}
// 每个写go程，向公共区写数据，写之前以写模式加锁，写结束解锁。
func WriteGo(idx int)  {
	for {
		data := rand.Intn(500)		// 0-499
		rwlock.Lock()		// 访问公共区之前，加写属性锁
		num = data
		fmt.Printf("------第%d个 写go程，写入：%d\n", idx, data)
		rwlock.Unlock()		// 访问公共区结束，解锁

		time.Sleep(time.Millisecond * 200)
	}
}

func main() {
	// 播种随机数种子
	rand.Seed(time.Now().UnixNano())

	for i:=0; i<5; i++ {		// 一次性创建 5 个读go程
		go readGo(i+1)
	}

	for i:=0; i<5; i++ {		// 一次性创建 5 个写go程
		go WriteGo(i+1)
	}

	for {
		runtime.GC()
	}
}