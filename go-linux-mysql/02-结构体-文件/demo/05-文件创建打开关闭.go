package main

import (
	"os"
	"fmt"
)

func main()  {
	f, err := os.Create("C:/itcast/test2.txt")
	if err != nil {
		fmt.Println("Create err: ", err)
		return
	}
	defer f.Close()

	//f, err = os.Open("C:/itcast/test2.txt")
	f, err = os.OpenFile("C:/itcast/test2.txt",  os.O_RDWR,0666)
	if err != nil {
		fmt.Println("Open err: ", err)
		return
	}
	defer f.Close()

	_, err = f.WriteString("hello")
	if err != nil {
		fmt.Println("WriteString err: ", err)
		return
	}
	fmt.Println("open ok")
}
