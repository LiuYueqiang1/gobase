package main

import "fmt"

// 结构体初始化
type person struct {
	name, city string
	age        int
}

func main() {
	var p1 person
	fmt.Printf("%#v\n", p1) //main.person{name:"", city:"", age:0}

	p2 := new(person)
	p2.name = "li"
	p2.city = "bei"
	p2.age = 16
	fmt.Println(p2) //&{li bei 16}

	p3 := &person{
		"大卫",
		"洛杉矶",
		18,
	}
	fmt.Println(p3) //&{大卫 洛杉矶 18}
	p4 := person{
		"大卫",
		"洛杉矶",
		19,
	}
	fmt.Println(p4) //{大卫 洛杉矶 19}

	//结构体内存布局
	type test struct {
		a int8
		b int8
		c int8
		d int8
	}
	n := test{
		1,
		2,
		3,
		4,
	}
	fmt.Printf("%p\n", &n) //这个取的是第一位的地址
	fmt.Printf("%p\n", &n.a)
	fmt.Printf("%p\n", &n.b)
	fmt.Printf("%p\n", &n.c)
	fmt.Printf("%p\n", &n.d)
}
