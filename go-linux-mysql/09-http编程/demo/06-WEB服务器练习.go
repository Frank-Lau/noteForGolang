package main

import (
	"net/http"
	"fmt"
	"os"
)

func sendFile(fileName string, w http.ResponseWriter)  {
	filePath := "C:/itcast/test" + fileName
	// 打开文件
	f, err := os.Open(filePath)
	errHtml := "<html><head><title></title></head> <body bgcolor=\"#cc99cc\"><h4>404 not found </h4> <hr>你所请求的文件不存在！！！不好意思！</body></html>"
	if err != nil {
		// 该文不存在！！！
		//w.Write([]byte("No such file or directroy"))
		w.Write([]byte(errHtml))
		return
	}
	defer f.Close()
	// 读取文件
	buf := make([]byte, 4096)
	for {
		n, _ := f.Read(buf)
		if n == 0 {
			fmt.Println("读取文件完成！")
			break
		}
		// 发送文件, 使用w，回写给 浏览器
		w.Write(buf[:n])
	}
}

func WEBHandle(w http.ResponseWriter, r*http.Request)  {
	fmt.Println("客户端请求：", r.URL)
	sendFile(r.URL.String(), w)
}

func main()  {
	// 注册回调函数
	http.HandleFunc("/", WEBHandle)

	// 绑定服务器地址结构
	http.ListenAndServe("127.0.0.1:8006", nil)
}
