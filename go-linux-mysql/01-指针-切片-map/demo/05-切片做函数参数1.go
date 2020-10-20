package main

import "fmt"

func test1(s []int)  {
	s[1] = 666
	fmt.Println("in test: s=", s)
}

func main()  {
	s1 := []int{1, 2, 3}

	fmt.Println("main before: s1=", s1)
	test1(s1)
	fmt.Println("main after: s1=", s1)
}
