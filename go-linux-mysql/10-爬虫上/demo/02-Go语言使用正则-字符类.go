package main

import (
	"regexp"
	"fmt"
)

func main()  {
	str := "abc a7c mfc cat 8ca azc cba"
	// 编译、解析正则表达式
	//rep := regexp.MustCompile(`a.c`)
	//rep := regexp.MustCompile(`a[0-9]c`)
	rep := regexp.MustCompile(`a\dc`)

	// 提取正确数据信息
	alls := rep.FindAllStringSubmatch(str, -1)
	fmt.Println(alls)
}
