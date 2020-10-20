package main

import "fmt"

// 定义函数指针类型. 指向一类函数：有两个参数（int,bool）有一个int返回值。
type FUNCP func(x int, y bool) int

func useCallback(x int, y bool, p FUNCP) int {
	return  p(x, y)
}

func addOne(x int, y bool) int {
	if y == true{
		x++
	}
	return x
}
func subTen(x int, y bool) int {
	if y == true {
		x -= 10
	}
	return x
}
func main()  {
/*	addOne(100, true)
	subTen(200, true)*/
	var p FUNCP
	p = addOne
	res := useCallback(200, true, p)
	fmt.Println("res:", res)

	res = useCallback(30, true, subTen)
	fmt.Println("res:", res)
}
