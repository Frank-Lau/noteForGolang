package main

import "fmt"

func test(b int)  {
	fmt.Println("b = ", b)
}

func main()  {

	var a int = 10
	fmt.Println("a=", a)

	var p *string  //
	p = new(string)

	test(67)

	fmt.Printf("*p = %q\n", *p)
}
