package main

import (
	"fmt"
)

// 嵌套结构体
type Address struct {
	city string
	mail string
}

type Person struct {
	name string
	age  int
	addr Address
}

// 嵌套结构体也可以使用匿名字段的方式
type Person2 struct {
	name string
	age  int
	Address
}

// 结构体继承
type Animal struct {
	name string
}

func (a *Animal) move() {
	fmt.Printf("%s会动\n", a.name)
}

type Dog struct {
	feet    int
	*Animal //狗继承了Animal结构体  采用了匿名字段的方式
}

func (b *Dog) wang() {
	fmt.Printf("%s会汪汪汪！\n", b.name)
}

func main() {
	var p1 = Person{
		name: "章子怡",
		age:  19,
		addr: Address{
			"北京", "li56",
		},
	}
	fmt.Printf("%#v\n", p1) //main.Person{name:"章子怡", age:19, addr:main.Address{city:"北京", mail:"li56"}}

	var p2 Person2
	p2.name = "李"
	p2.age = 15
	p2.Address.mail = "lklk@16.com"
	p2.city = "济南"
	fmt.Printf("%#v\n", p2) //main.Person2{name:"李", age:15, Address:main.Address{city:"济南", mail:"lklk@16.com"}}

	d1 := &Dog{
		feet: 4,
		Animal: &Animal{
			name: "乐乐",
		},
	}
	d1.move() //乐乐会动
	d1.wang() //乐乐会汪汪汪！
}
