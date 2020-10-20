package main

import "fmt"

/*// map 定义和赋值
func main()  {
	//var m1 map[string]int = map[string]int{"aaa":100, "bbb":200, "ccc":300}
	m1 := map[string]int{"aaa": 100, "bb": 200, "ccc": 300}
	fmt.Println("m1=", m1)

	m2 := make(map[int]string)		// 开一个没有元素的 map ，有空间
	m2[1] = "hello"					// 自动扩容。
	fmt.Println("m2=", m2, "len=", len(m2))

	m3 := make(map[int]string, 5)	// 5 表初始容量。 —— cap
	m3[100] = "go"
	m3[100] = "C++"
	fmt.Println("m2=", m3, "len=", len(m3), ) //"cap=", cap(m3) 不允许对map使用

	var m4 map[string]int		// 创建一个map 类型变量，没有空间
	m4 = make(map[string]int)
	m4["hello1"] = 100			// 如果没有 make 赋值--- 报错
	m4["hello2"] = 200
	m4["hello3"] = 100
	fmt.Println("m4=", m4, "len=", len(m4))
}*/

// map的遍历和删除
func main()  {
	m1 := map[string]int{"a": 100, "b": 200, "c": 300, "f":896, "k":343}
	for key, value := range m1 {
		fmt.Println(key, ":", value)
	}
	for data := range m1 {
		fmt.Println(data)
	}

	// 判断map中的key 是否存在
	if value, isTrue := m1["m"]; isTrue == true {
		fmt.Printf("1, value=%d, isTrue=%v\n", value, isTrue)
	} else {
		fmt.Printf("2, value=%d, isTrue=%v\n", value, isTrue)
	}

	retMap := myDelete(m1)
	fmt.Println("main:", m1)

	fmt.Println(retMap)
}

func myDelete(m1 map[string]int) map[int]string {

	delete(m1, "c")
	fmt.Println("myDelete:", m1)

	var m2 map[int]string = map[int]string{}

	return m2
}

