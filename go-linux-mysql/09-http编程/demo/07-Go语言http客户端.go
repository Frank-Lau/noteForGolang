package main

import (
	"fmt"
	"net/http"
	"io"
)

func main()  {
	// 获取 web服务器回发的 http应答协议包. http:// 不能省略
	//resp, err := http.Get("http://127.0.0.1:8000/itcast.html")
	resp, err := http.Get("http://www.baidu.com")
	if err != nil {
		fmt.Println("GET err：", err)
		return
	}
	defer resp.Body.Close()

	// 查看 http应答协议
	fmt.Println("StatusCode=", resp.StatusCode)
	fmt.Println("Proto=", resp.Proto)
	fmt.Println("Header=", resp.Header)
	fmt.Println("Body=", resp.Body)

	// 循环从 Body 中读取服务器发送的数据。
	buf := make([]byte, 4096)

	var result string
	for {
		n, err := resp.Body.Read(buf)
		if n == 0 {
			break
		}
		if err != nil && err != io.EOF {
			fmt.Println("resp.Body.Read err:", err)
			break
		}
		result += string(buf[:n])
	}
	fmt.Println("result:", result)
}
