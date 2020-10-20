package main

import (
	"fmt"
	"net/http"
)



// 传值
func swap(a, b int)  {
	a, b = b, a
	fmt.Println("swap : a=", a, " b=", b)
}
// 传地址（传引用）--- 传值
func swap2(a, b *int)  {
	*a, *b = *b, *a
	fmt.Println("swap : a=", *a, " b=", *b)
}

func main()  {
	var a int = 10
	var b int = 207

	fmt.Println("before main : a=", a, " b=", b)
	//swap(a, b)
	swap2(&a, &b)
	fmt.Println("after main : a=", a, " b=", b)

	http.Get()
}
