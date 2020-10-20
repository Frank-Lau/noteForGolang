package main

import (
	"fmt"
	"os"
	"strings"
	"bufio"
	"io"
)

// 打开文件，统计该文件内有多少个 “love”
func countLove(name, path string) int {
	// 记录love出现次数的变量
	var counter int = 0

	// 拼接完整的 .txt 文件路径
	fileName := path + "/" + name

	// 只读打开 .txt 文件
	f, err := os.Open(fileName)
	if err != nil {
		fmt.Println("Open err:", err)
		return 0
	}
	defer f.Close()

	// 按行读文件。 1. 创建一个带有缓冲区的 reader
	reader := bufio.NewReader(f)

	// 2. 循环 从缓冲区中，提取一行数据
	for {
		buf, err := reader.ReadBytes('\n')
		if err != nil && err == io.EOF {
			break
		}
		// 拆分 一行数据成一个个单词
		words := strings.Fields(string(buf))		// [ "this" "is" a test for love ]

		// 变量切片，找 love， 统计
		for _, word := range words {
			if word == "love" {
				counter++
			}
		}
	}
	return counter
}

func main()  {
	// 提示用户，输入待查询的目录
	var path string
	fmt.Print("请输入目录:")
	fmt.Scan(&path)

	// 打开目录
	dirFp, err := os.OpenFile(path, os.O_RDONLY, os.ModeDir)
	if err != nil {
		fmt.Println("OpenFile err:", err)
		return
	}
	defer dirFp.Close()

	// 统计所有love 个数
	allLove := 0

	// 提取所有目录项
	dirInfoS, err := dirFp.Readdir(-1)
	if err != nil {
		fmt.Println("Readdir err:", err)
		return
	}
	// 遍历所有目录项， 找出 .txt 文件 —— .Name() 获取文件名
	for _, file := range dirInfoS {
		if strings.HasSuffix(file.Name(), ".txt") {		//找到 .txt 文件
			// 封装函数。打开文件，统计该文件内有多少个 “love”
			allLove += countLove(file.Name(), path)
		}
	}

	// 打印目录下所有文件中 love 出现次数。
	fmt.Println("allLove = ", allLove)
}
