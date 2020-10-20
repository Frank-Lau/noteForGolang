package main

import (
	"fmt"
)

type Person3 struct {
	name string
	age int
	flg bool
	interest []string
}

func initPerson(p *Person3)  {
	p.name = "AAA"
	p.age = 800
	p.flg = true
	p.interest = []string{"sleeping", "eating", "shopping"}
}

func initPerson2() *Person3 {
	p := new(Person3)
	p.name = "AAA"
	p.age = 800
	p.flg = true
	p.interest = []string{"sleeping", "eating", "shopping"}
	return p
}

func initPerson3(p *Person3) *Person3 {
	if p == nil {
		p = new(Person3)
	}
	p.name = "AAA"
	p.age = 800
	p.flg = true
	p.interest = []string{"sleeping", "eating", "shopping"}
	return p
}

func initPerson4(p **Person3) {
	if *p == nil {
		*p = new(Person3)
	}
	(*p).name = "AAA"
	(*p).age = 800
	(*p).flg = true
	(*p).interest = []string{"sleeping", "eating", "shopping"}
}

func main()  {
	var p1 Person3
	initPerson(&p1)
	fmt.Println("p1=", p1)

	p2 := initPerson2()
	fmt.Println("p2=", p2)

	var p3 *Person3
	p3 = initPerson3(p3)
	fmt.Println("p3=", p3)

	var p4 *Person3
	initPerson4(&p4)
	fmt.Println("p4=", p4)

}
