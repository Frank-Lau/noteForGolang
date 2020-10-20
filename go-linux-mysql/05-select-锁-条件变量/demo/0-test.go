package main

import "time"

func main()  {

	time1 := time.NewTimer(3)
	time1.Reset(2)
	time2 := time.NewTicker(2)
	time2.Stop()
}
