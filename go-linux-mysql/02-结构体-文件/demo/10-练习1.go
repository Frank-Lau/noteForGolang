package main

import (
	"os"
	"fmt"
	"strings"
)

func main()  {
	fdir, err := os.OpenFile("C:/itcast/test", os.O_RDONLY, os.ModeDir)
	if err != nil {
		fmt.Println("OpenFile err:", err)
		return
	}
	defer fdir.Close()

	fileInfos, err := fdir.Readdir(-1)
	if err != nil {
		fmt.Println("Readdir err:", err)
		return
	}
	for _, fileInfo := range fileInfos {
		if strings.HasSuffix(fileInfo.Name(), ".jpg") {
			fmt.Println(fileInfo.Name(), "是一个 jpg 文件")
		}
	}
}
