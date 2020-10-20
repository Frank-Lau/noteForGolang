package main

import (
	"fmt"
	"strings"
)

func main()  {
	// Contains
	fmt.Println(strings.Contains("hello", "lllo"))
	// Join
	str := []string{"hello", "go", "haha", "xixi"}
	retStr := strings.Join(str, "%")
	fmt.Println(retStr)
	// Trim
	retStr = strings.Trim(" hello ", " ")
	fmt.Println(retStr)
	// Replace
	myStr := "this is a test for a Replace"
	retStr = strings.Replace(myStr, "a ", "some ", -1)
	fmt.Println(retStr)
	// Split
	str = strings.Split(myStr, "a ")
	for _, word := range str {
		fmt.Println(word)
	}
	// Fields
	str = strings.Fields(myStr)
	for _, word := range str {
		fmt.Println(word)
	}
	// HasSuffix
	fmt.Println(strings.HasSuffix("word cup.jpg", ".mp3"))
	// HasPrefix
	fmt.Println(strings.HasPrefix("hello world", "hel"))
}
