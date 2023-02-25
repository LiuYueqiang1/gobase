package main

import "fmt"

// 使用值类型的接收者和指针类型的接收者的区别
type person struct {
	Name string
	Age  int
}

// 接口的嵌套
type animal interface {
	sayr
	mover
}

type sayr interface {
	say()
}
type mover interface {
	mo()
}
type cat2 struct {
	Name string
}

func (c cat2) say() {
	fmt.Printf("%s会喵喵喵~~~\n", c.Name)
}

// 使用指针类型的接收者，只有类型指针能够保存到接口变量中
func (p *person) say() {
	fmt.Printf("%d的%s会啊啊啊\n", p.Age, p.Name)
}
func (p *person) mo() { //一个struct类型可以实现多个接口
	fmt.Printf("%d的%s会动\n", p.Age, p.Name)
}

func main() {
	//p1 := person{ //p1是person类型的值，error
	//	Name: "纳扎",
	//	Age:  18,
	//}

	p2 := &person{ //p2是person类型的指针
		Name: "纳扎",
		Age:  16,
	}
	var s sayr = p2 //定义一个sayr类型的变量=p2
	s.say()
	var m mover = p2 //定义一个mover类型的变量=p2
	m.mo()

	c3 := cat2{
		Name: "米粒",
	}
	s = c3
	s.say()

	//s = p1 //p1是person类型的值没有实现sayr接口
	//s.say()

	var a animal //接口的嵌套
	a = p2
	a.say()
	a.mo()
	println(a)
}
