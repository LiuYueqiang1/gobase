package main

import "fmt"

//	func (接收者变量 接收者类型) 方法名(参数列表) (返回参数) {
//	   函数体
//	} *****

// 任意类型添加方法
type Myint int

func (mi Myint) sayHello() {
	fmt.Println("hello,my type is int")
}

// 结构体类型的匿名字段
//结构体允许其成员字段在声明时没有字段名而只有类型，这种没有名字的字段就称为匿名字段。

type Person struct {
	string
	int
}

func main() {
	var a Myint
	a.sayHello()
	a = 100
	fmt.Printf("%#v,%T\n", a, a)

	var b Myint = 20
	println(b)

	//调用匿名字段
	var p1 = Person{
		"dawei",
		18,
	}
	fmt.Println(p1.string, p1.int) //dawei 18
	fmt.Printf("%#v\n", p1)        //main.Person{string:"dawei", int:18}
}
