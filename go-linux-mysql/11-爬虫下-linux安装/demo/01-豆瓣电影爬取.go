package main

import (
	"fmt"
	"strconv"
	"net/http"
	"io"
	"regexp"
	"os"
)

func httpGet(url string) (result string, err error) {
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

// 爬取一个页面所有数据内容
func spiderPage(idx int, quit chan<- int)  {
	// 根据网页分页器规律，组织每个页面url
	url := "https://movie.douban.com/top250?start=" + strconv.Itoa((idx-1)*25) + "&filter="
	// 爬取一个页面
	result, err:= httpGet(url)
	if err != nil {
		fmt.Println("httpGet err:", err)
		return
	}
	// 提取电影名, 编译解析正则表达式
	reg := regexp.MustCompile(`<img width="100" alt="(?s:(.*?))" src="`)

	alls := reg.FindAllStringSubmatch(result, -1)
	// 创建一个切片，存储一个网页中所有电影名称
	fileName := make([]string, 0)
	for _, data := range alls {
		//fmt.Println(data[1])
		fileName = append(fileName, data[1])
	}

	// 提取分数, 编译解析正则表达式
	pattern := `<span class="rating_num" property="v:average">(?s:(.*?))</span>`

	reg2 := regexp.MustCompile(pattern)

	alls2 := reg2.FindAllStringSubmatch(result, -1)
	// 创建一个切片，存储一个网页中所有电影名称
	fileScore := make([]string, 0)
	for _, data := range alls2 {
		fileScore = append(fileScore, data[1])
	}

	// 提取评价人数, 编译解析正则表达式
	reg3 := regexp.MustCompile(`<span>(?s:(\d*?))人评价</span>`)

	alls3 := reg3.FindAllStringSubmatch(result, -1)
	// 创建一个切片，存储一个网页中所有电影名称
	peopleNum := make([]string, 0)
	for _, data := range alls3 {
		//fmt.Println(data[1])
		peopleNum = append(peopleNum, data[1])
	}
	// 封装函数，将 电影名、评分、评价人数，存储至一个文件。
	save2File(idx, fileName, fileScore, peopleNum)

	quit<-idx
}

func save2File(idx int, fileName, fileScore, peopleNum []string)  {
	// 创建文件
	f, err := os.Create("C:/itcast/" + "第 " + strconv.Itoa(idx) + "页.txt")
	if err != nil {
		fmt.Println("Create err:", err)
		return
	}
	defer f.Close()
	n := len(fileName)
	// 拼接抬头
	f.WriteString("电影名称" + "\t" + "分数" + "\t" + "评价人数" + "\n")
	// 将 电影名、分数、评价人数 一一对应写入文件
	for i:=0; i<n; i++ {
		fileInfo := fileName[i] + "\t" + fileScore[i] + "\t" + peopleNum[i] + "\n"
		f.WriteString(fileInfo)
	}
}

func working(start, end int)  {
	quit := make(chan int)
	for i:=start; i<=end; i++ {
		go spiderPage(i, quit)
	}
	for i:=start; i<=end; i++ {
		fmt.Printf("第%d个页面爬取完毕。。。\n", <-quit)
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
	working(start, end)
}
