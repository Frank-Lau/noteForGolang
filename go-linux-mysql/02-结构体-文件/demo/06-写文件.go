package main

import (
	"os"
	"fmt"
	"io"
)

func main()  {
	f, err := os.OpenFile("test.txt", os.O_RDWR, 0666)
	if err != nil {
		fmt.Println("OpenFile err:", err)
		return
	}
	defer f.Close()

	// 按字符串写入
	n, err := f.WriteString("This is a ")
	if err != nil {
		fmt.Println("WriteString err:", err)
		return
	}
	fmt.Println("n = ", n)

	//off, err := f.Seek(5, io.SeekStart)
	off, err := f.Seek(-2, io.SeekEnd)
	if err != nil {
		fmt.Println("Seek err:", err)
		return
	}
	//fmt.Println("off=", off)

	n, err = f.WriteAt([]byte("8888"), off)
	if err != nil {
		fmt.Println("WriteAt err:", err)
		return
	}
}

















