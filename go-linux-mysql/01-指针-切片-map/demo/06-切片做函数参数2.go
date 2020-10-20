package main

import "fmt"

func test2(s []int)  {
	s = append(s, 777)
	fmt.Println("in test: s=", s, "cap=", cap(s))
}

func main()  {
	s1 := []int{1, 2, 3}

	fmt.Println("main before: s1=", s1, "cap=", cap(s1))
	test2(s1)
	fmt.Println("main after: s1=", s1, "cap=", cap(s1))

}
