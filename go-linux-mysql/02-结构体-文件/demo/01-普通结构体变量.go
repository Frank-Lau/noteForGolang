package main

import (
	"fmt"
	"unsafe"
)

/*func main()  {
	var p1 Person = Person{"andy", 'f', 100}
	fmt.Println("p1=", p1)
	p2 := Person{"andy", 'f', 100}
	fmt.Println("p2=", p2)

	p3 := Person{name:"lucy", age:30}
	fmt.Println("p3=", p3)

	p3.name = "zhangsan"
	p3.sex = 'm'
	fmt.Println("p3=", p3)

	fmt.Printf("&p3=%p\n", &p3)
	fmt.Printf("&p3.name=%p\n", &p3.age)

	p4 := Person{100, "aaa", 'm'}
	p5 := Person{100, "aaa", 'f'}
	p6 := Person{100, "aaa", 'f'}

	fmt.Println("p4 == p5 ?", p4 == p5)
	fmt.Println("p4 == p6 ?", p4 != p6)
}*/

type Person struct {
	age int				// 8
	name string			// 16
	sex byte			// 1
}

func test(p Person)  {
	p.name = "xyz"
	p.age = 678

	fmt.Println("in test p=", p)
}

func main()  {
	var p Person  // var a int

	p.name = "bbb"
	p.age = 30
	fmt.Println("1 p=", p)
	test(p)
	fmt.Println("2 p=", p)

	fmt.Println(unsafe.Sizeof(p))
	fmt.Println(unsafe.Sizeof(p.name))
	fmt.Println(unsafe.Sizeof(p.age))
	fmt.Println(unsafe.Sizeof(p.sex))
}
