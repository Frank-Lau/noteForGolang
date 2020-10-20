package main

import (
	"regexp"
	"fmt"
)

func main()  {
	str := `<!DOCTYPE html>
<html lang="zh-CN">
<head>
	<title>Go语言标准库文档中文版 | Go语言中文网 | Golang中文社区 | Golang中国</title>
	<meta name="viewport" content="width=device-width, initial-scale=1, maximum-scale=1.0, user-scalable=no">
	<meta http-equiv="X-UA-Compatible" content="IE=edge, chrome=1">
	<meta charset="utf-8">
	<link rel="shortcut icon" href="/static/img/go.ico">
	<link rel="apple-touch-icon" type="image/png" href="/static/img/logo2.png">
	<meta name="author" content="polaris <polaris@studygolang.com>">
	<meta name="keywords" content="中文, 文档, 标准库, Go语言,Golang,Go社区,Go中文社区,Golang中文社区,Go语言社区,Go语言学习,学习Go语言,Go语言学习园地,Golang 中国,Golang中国,Golang China, Go语言论坛, Go语言中文网">
	<meta name="description" content="Go语言文档中文版，Go语言中文网，中国 Golang 社区，Go语言学习园地，致力于构建完善的 Golang 中文社区，Go语言爱好者的学习家园。分享 Go 语言知识，交流使用经验">
</head>

	<div>hel?lo</div>
	<div>hello regexp</div>
	<div>hello7</div>
	<div>hello666</div>
	<div>
		  2块钱啥时候还？
         过了年再说吧！
         刚买了车，没钱。。。
	</div>

<frameset cols="15,85">
	<frame src="/static/pkgdoc/i.html">
	<frame name="main" src="/static/pkgdoc/main.html" tppabs="main.html" >
	<noframes>
	</noframes>
</frameset>
</html>`

	// 1. 编译、解析正则表达式
	//rep := regexp.MustCompile(`<div>(.*)</div>`)
	//rep := regexp.MustCompile(`<div>(?s:(.*))</div>`)
	rep := regexp.MustCompile(`<div>(?s:(.*?))</div>`)

	// 2. 提取需要数据
	alls := rep.FindAllStringSubmatch(str, -1)
	//fmt.Println(alls)
	n := len(alls)
	for i:= 0; i<n; i++ {
		fmt.Printf("all[%d][0]=%s\n", i, alls[i][0])
		fmt.Printf("all[%d][1]=%s\n", i, alls[i][1])
	}
}
