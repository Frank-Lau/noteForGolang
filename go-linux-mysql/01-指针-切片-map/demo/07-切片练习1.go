package main

import "fmt"

// 传值 --- append
func noEmpty(str []string) []string {
	// 定义一个空切片
	//out := make([]string, 0)

	// 从源切片上截取一个空切片
	out := str[:0]
	// 遍历，取出每一个 元素
	for _, word := range str {
		if word != "" {
			out = append(out, word)
		}
	}
	return out
}
// 传值 --- 不使用append
func noEmpty2(str []string) []string {
	i := 0
	// 遍历，取出每一个 元素
	for _, word := range str {
		if word != "" {
			str[i] = word
			i++
		}
	}
	return str[:i]
}

func main()  {
	str := []string{"red", "", "black", "", "", "pink", "blue"}
	//afterData := noEmpty(str)
	afterData := noEmpty2(str)
	fmt.Println("afterData=", afterData)
}
