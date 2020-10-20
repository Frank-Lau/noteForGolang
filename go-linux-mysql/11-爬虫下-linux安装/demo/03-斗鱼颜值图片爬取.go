package main

import (
	"fmt"
	"strconv"
	"net/http"
	"io"
	"regexp"
	"os"
)

func httpGetJoke(url string) (result string, err error) {
	// 使用url发送请求给web服务器
	resp, err1 := http.Get(url)
	if err1 != nil {
		err = err1;
		return
	}
	defer resp.Body.Close()
	buf := make([]byte, 4096)
	// 循环获web服务器的应答数据包
	for {
		n, err2 := resp.Body.Read(buf)
		if n == 0 {
			break;
		}
		if err2 != nil && err2 != io.EOF {
			err = err2;
			return
		}
		result += string(buf[:n])
	}
	return
}

// 爬取一张图片， 保存成一个文件。
func spiderImg(url string, idx int, quit chan<- int) {
	// 创一个文件。
	f, err := os.Create("C:/itcast/img/" + strconv.Itoa(idx) + ".jpg")
	if err != nil {
		fmt.Println("Create err； ", err)
		return
	}
	defer f.Close()

	// 使用url发送请求给web服务器
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Get err； ", err)
		return
	}
	defer resp.Body.Close()

	buf := make([]byte, 4096)
	// 循环获web服务器的应答数据包
	for {
		n, err := resp.Body.Read(buf)		// 存储图片的二进制数据。
		if n == 0 {
			break;
		}
		if err != nil && err != io.EOF {
			fmt.Println("resp.Body.Read err； ", err)
			return
		}
		f.Write(buf[:n])
	}
	quit <- idx
}

func main()  {
	// 获取 url
	url := "https://www.douyu.com/g_yz"

	// 爬取 颜值模块整个页面
	result, _ := httpGetJoke(url)

	// 编译解析正则，获取每张图片的 url
	reg := regexp.MustCompile(`data-original="(?s:(.*?))"`)
	alls := reg.FindAllStringSubmatch(result, -1)

	quit := make(chan int)

	n := len(alls)
	for idx, data := range alls {
		//fmt.Println(data[1])		// data[1]一张图片的 url
		go spiderImg(data[1], idx, quit)
	}
	for i:=0; i<n; i++ {
		<-quit
	}
	fmt.Println("完毕！")
}
