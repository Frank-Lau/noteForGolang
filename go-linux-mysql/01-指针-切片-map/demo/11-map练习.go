package main

import (
	"fmt"
	"strings"
)

func wordCountFunc(str string) map[string]int {
	// 拆分字符串到 []string
	words := strings.Fields(str)
	// 创建用于存储每一个单词的 map
	m := make(map[string]int)

	for _, word := range words {
		//m[word]++
		//if val, ok := m[word]; ok == true {
		if _, ok := m[word]; ok{
			m[word]++
		} else {
			m[word] = 1
		}
	}
	return m
}

func main()  {
	str := "I love my work and I love my family too too love too too too"
	m := wordCountFunc(str)
	for key, val := range m {
		fmt.Println(key, ":", val)
	}
}
