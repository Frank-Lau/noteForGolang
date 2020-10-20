package main

import "fmt"

func noSame(str []string) []string {
	out := str[:1]		// 创建只有一个“red"元素的 切片

	for _, word := range str {
		i:=0
		for ; i<len(out); i++ {
			if out[i] == word {
				break			// out 切片中已经有。不追加
			}
		}
		// 从 for 表达式 2 中跳出循环， 应该追加。
		if i == len(out) {
			out = append(out, word)
		}
	}
	return out
}

func main()  {
	str := []string{"red", "black", "red", "pink", "Yello", "blue", "red", "pink", "Yello", "blue"}

	afterData := noSame(str)

	fmt.Println("afterData=", afterData)
}
