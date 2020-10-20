package main

import (
	"fmt"
	"runtime"
)

func main()  {
	go func() {
		for {
			fmt.Println("this is a 匿名 go 程")
			//time.Sleep(time.Millisecond * 200)
		}
	}()

	for i := 0; i < 10; i++ {
		fmt.Println("main ...")
		runtime.Gosched()
	}
}
