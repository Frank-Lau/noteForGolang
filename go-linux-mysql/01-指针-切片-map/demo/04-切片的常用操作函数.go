package main

import "fmt"

// append
/*func main()  {
	s1 := [4]int {1, 2, 5, 6}

	fmt.Println("s1", s1, "len=", len(s1), "cap=", cap(s1))

	s1 = append(s1, 888)

	fmt.Println("s1", s1, "len=", len(s1), "cap=", cap(s1))

}*/

// copy
func main()  {
	data := [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	s1 := data[8:]		// {8, 9}
	s2 := data[:5]		// {0, 1, 2, 3, 4}

	copy(s1, s2)

	fmt.Println("s1=", s1)
	fmt.Println("s2=", s2)

}


