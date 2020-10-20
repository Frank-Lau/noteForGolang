package main

import (
	"fmt"
	"strconv"
	"net/http"
	"io"
	"regexp"
	"strings"
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

// 爬取一个笑话的页面,得到一个title 和一个对应的 content
func spiderOneJokePage(url string) (title, content string)  {
	// 调用 函数httpGetJoke 获取整个网页数据
	result, err := httpGetJoke(url)
	if err != nil {
		fmt.Println("httgGetJoke err:", err)
		return
	}
	// 编译解析正则表达式，提取 title
	reg := regexp.MustCompile(`<h1>(.*?)</h1>`)
	alls := reg.FindAllStringSubmatch(result, 1)			// 提取第一个匹配项。
	for _, data := range alls {
		title = data[1]
		title = strings.Replace(title, " ", "", -1)
		title = strings.Replace(title, "\t", "", -1)
/*		title = strings.Replace(title, "\n", "", -1)
		title = strings.Replace(title, "\r", "", -1)*/
		break;
	}
	// 编译解析正则表达式，提取 content
	reg2 := regexp.MustCompile(`<div class="content-txt pt10">(?s:(.*?))<a id="prev"`)
	alls2 := reg2.FindAllStringSubmatch(result, -1)			// 提取第一个匹配项。
	for _, data := range alls2 {
		content = data[1]
		content = strings.Replace(content, " ", "", -1)
		content = strings.Replace(content, "\t", "", -1)
		content = strings.Replace(content, "&nbsp;", "", -1)
		content = strings.Replace(content, "<br/>", "", -1)
		break;
	}
	return
}

// 爬取带有10个笑话页面
func spiderJokeS(idx int, quit chan<- int)  {
	// 根据网页分页器规律，组织每个页面url
	url := "https://www.pengfu.com/xiaohua_" + strconv.Itoa(idx) + ".html"

	// 调用 函数httpGetJoke 获取整个网页数据
	result, err := httpGetJoke(url)
	if err != nil {
		fmt.Println("httgGetJoke err:", err)
		return
	}
	// 筛选网页中的所有 笑话的 url
	reg := regexp.MustCompile(`<h1 class="dp-b"><a href="(?s:(.*?))"`)
	// 过滤，提取 url
	alls := reg.FindAllStringSubmatch(result, -1)

	titleS := make([]string, 0)			// 创建切片存储所有的标题
	contentS := make([]string, 0)		// 创建切片存储所有的段子内容

	for _, data := range alls {
		//fmt.Println(data[1])		// data[1] 指向一个joke的url
		title, content := spiderOneJokePage(data[1])
		titleS = append(titleS, title)
		contentS = append(contentS, content)
	}
	// 封装函数，保存title和content
	saveToFile(idx, titleS, contentS)
	quit <- idx
}

func saveToFile(idx int, titleS, contentS []string)  {
	f, err := os.Create(strconv.Itoa(idx) + ".txt")
	if err != nil {
		fmt.Println("Create err； ", err)
		return
	}
	defer f.Close()

	n := len(titleS)				// 求出[] 的元素个数
	for i:=0; i<n; i++ {
		f.WriteString(titleS[i]+"\n")			// 写入一个标题
		f.WriteString(contentS[i]+"\n")			// 对应标题写入内容
		f.WriteString("-------------------------------------------------------------------------\n")
	}
}

func toWork(start, end int)  {
	quit := make(chan int)
	for i:=start; i<=end; i++ {
		go spiderJokeS(i, quit)
	}

	for i:=start; i<=end; i++ {
		fmt.Printf("第%d个页面爬取完毕...\n", <-quit)
	}
}

func main()  {
	//  提示用户 输入起始、终止页面  start、end
	var start, end int
	fmt.Print("请输入爬取起始页（>=1）：")
	fmt.Scan(&start)
	fmt.Print("请输入爬取终止页（>=start）：")
	fmt.Scan(&end)

	//创建函数 working ，根据起始、终止页。组织每一页的url
	toWork(start, end)
}
