package main

import "fmt"

func main()  {
	arr := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	s1 := arr[1:5:7]  // 2, 3, 4, 5  []  []  stuct {  *p }

	s2 := s1[3:5:6]	// 5, 6		sturct{  *p  }

	fmt.Println("s1=", s1, " len=", len(s1), " cap=", cap(s1))
	fmt.Println("s2=", s2, " len=", len(s2), " cap=", cap(s2))
}
