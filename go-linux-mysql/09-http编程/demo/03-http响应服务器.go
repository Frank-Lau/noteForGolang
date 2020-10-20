package main

import "net/http"

func handler(w http.ResponseWriter, r *http.Request)  {
	// 写数据给客户端（浏览器）
	w.Write([]byte("haha hehe xixi hoho !!!"))
}

func main()  {
	// 注册回调函数
	http.HandleFunc("/hello", handler)

	// 绑定服务器地址结构
	http.ListenAndServe("127.0.0.1:8000", nil)
}
