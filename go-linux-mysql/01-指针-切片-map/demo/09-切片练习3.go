package main

import "fmt"

func myMove(data []int, idx int) []int {
	//copy({{8, 9}, 9}, {8, 9})
	copy(data[idx:], data[idx+1:])
	return data[:len(data)-1]
}

func main()  {
	data := []int{5, 6, 7, 8, 9}
	afterData := myMove(data, 2)
	fmt.Println(afterData)
}
