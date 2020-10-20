package main

import (
	"os"
	"fmt"
)

func main()  {
	fdir, err := os.OpenFile("./", os.O_RDONLY, os.ModeDir)
	if err != nil {
		fmt.Println("OpenFile err:", err)
		return
	}
	defer fdir.Close()

	// 读取目录，获取目录项
	fileInfo, err := fdir.Readdir(-1)
	if err != nil {
		fmt.Println("Readdir err:", err)
		return
	}
	for _, data := range fileInfo {
		if data.IsDir() {
			fmt.Println(data.Name(), "is a Dir")
		} else {
			fmt.Println(data.Name(), "is not a Dir")
		}
	}
}
