package main

import (
	"os"
	"fmt"
)

/*func main()  {
	list := os.Args
	n := len(list)
	for i:=0; i<n; i++ {
		fmt.Printf("list[%d]:%s\n", i, list[i])
	}
}*/
// 命令行参数，传递文件名
func main()  {
	list := os.Args
	if len(list) != 2 {
		fmt.Println("请按套路输入：格式为： go run xxx.go 文件名")
		return
	}
	filePath := list[1]			// 保存命令行参数传递进来的 文件绝对路径
	// 获取文件属性
	fileInfo, err :=  os.Stat(filePath)
	if err != nil {
		fmt.Println("Stat err:", err)
		return
	}
	fmt.Println("文件名为：", fileInfo.Name())	// 不带有 路径的文件名
	fmt.Println("文件大小：", fileInfo.Size())
}
