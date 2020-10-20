package main

import (
	"fmt"
	"strconv"
	"os"
	"net/http"
	"io"
)

// 保存一个网页数据成一个 html 文件。
func save2File(result string, idx int)  {
	// 组织文件名
	fileName := "第 " + strconv.Itoa(idx) + " 页.html"
	// 创建新文件。
	f, err := os.Create(fileName)
	if err != nil {
		fmt.Println("Create err:", err)
		return
	}
	defer f.Close()
	f.WriteString(result)
}

// 爬取一个页面。
func spiderPage(i int, quit chan<- int)  {
	// 组织每一页的 URL
	url:="https://tieba.baidu.com/f?kw=%E7%BB%9D%E5%9C%B0%E6%B1%82%E7%94%9F&ie=utf-8&pn=" + strconv.Itoa((i-1)*50)
	// 获取一个网页的数据
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Get err:", err)
		return
	}
	defer resp.Body.Close()
	// 创建缓冲区，保存读到的网页数据。
	buf := make([]byte, 4096)
	var result string
	for {
		n, err := resp.Body.Read(buf)
		if n == 0 {
			break
		}
		if err != nil && err != io.EOF {
			fmt.Println("Body.read err:", err)
			break
		}
		// 拼接每一个buf中的数据到 result 字符串中
		result += string(buf[:n])
	}
	// 封装函数，存储一个网页数据。
	save2File(result, i)
	// 一个网页保存结束，通知主go程
	quit <- i
}

func working(start, end int) {
	// 创建一个 channel 协调主、子go程执行先后顺序。
	quit := make(chan int)

	// 爬取 每一页数据。循环一次，一页。
	for i:=start; i<=end; i++ {
		go spiderPage(i, quit)  // 双 -- 单
	}
	for i:=start; i<=end; i++ {
		fmt.Printf("第%d个页面爬取完毕...\n", <-quit)
	}
/*	for idx := range quit {
		fmt.Printf("第%d个页面爬取完毕...\n", idx)
	}*/
}

func main()  {
	// 提示用户 输入爬取的起始、终止页面
	var start, end int
	fmt.Print("请输入爬取的起始页面(>=1):")
	fmt.Scan(&start)
	fmt.Print("请输入爬取的终止页面(>=start):")
	fmt.Scan(&end)

	// 封装函数，爬取贴吧数据。
	working(start, end)
}
