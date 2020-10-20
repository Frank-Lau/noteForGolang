package main

import (
	"fmt"
	"unsafe"
)

type Person2 struct {
	age int				// 8
	name string			// 16
	sex byte			// 1
}

/*func main()  {
	var p1 *Person2 = &Person2{18, "abc", 'f'}
	fmt.Println("p1=", p1)

	p2 := &Person2{18, "abc", 'f'}
	fmt.Println("p2=", p2)

	p3 := &Person2{sex:'m', name:"xyz"}
	fmt.Println("p3=", p3)

	var p4 *Person2
	p4 = new(Person2)		// 给指针开辟内存空间
	p4.sex = 'f'
	p4.name = "opq"
	fmt.Println("p4=", p4)

	fmt.Printf("p4=%p\n", p4)
	fmt.Printf("&p4.age = %p\n", &p4.age)
}*/

func test2(p *Person2)  {
	p.age = 100
	p.name = "qqq"
	p.sex = 'f'
	fmt.Println(" ==== size=", unsafe.Sizeof(p))
}

func main()  {

	var p1 *int
	var p2 *byte
	var p3 *Person2  // nil 没有空间

	var vP Person2  // 有内存空间

	fmt.Println("size=", unsafe.Sizeof(p1))
	fmt.Println("size=", unsafe.Sizeof(p2))
	fmt.Println("size=", unsafe.Sizeof(p3))
	p3 = new(Person2)

	test2(p3)
	test2(&vP)
	fmt.Println("p3=", p3)
	fmt.Println("size=", unsafe.Sizeof(Person2{}))
}
