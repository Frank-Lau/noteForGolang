package main

import (
	"net/http"
	"fmt"
)

func myHandle(w http.ResponseWriter, r *http.Request)  {
	// w 写会给浏览器， r，从浏览器读
	fmt.Println("Method =", r.Method)
	fmt.Println("URL=", r.URL)
	fmt.Println("Proto=", r.Proto)
	fmt.Println("RemoteAddr=", r.RemoteAddr)
	fmt.Println("Body=", r.Body)

	w.Write([]byte("this is 服务器"))
}

func main()  {
	// 注册回调函数
	http.HandleFunc("/itcast.html", myHandle)

	// 绑定地址结构
	http.ListenAndServe("127.0.0.1:8000", nil)
}
