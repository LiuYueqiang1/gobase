package main

import "fmt"

type Person struct {
	name, city string
	age        int8
}

// 构造函数
// Go语言的结构体没有构造函数，我们可以自己实现。 例如，下方的代码就实现了一个person的构造函数。
// 因为struct是值类型，如果结构体比较复杂的话，值拷贝性能开销会比较大，
// 所以该构造函数返回的是结构体指针类型Go语言的结构体没有构造函数，我们可以自己实现。
// 例如，下方的代码就实现了一个person的构造函数。 因为struct是值类型，如果结构体比较复杂的话，
// 值拷贝性能开销会比较大，所以该构造函数返回的是结构体指针类型
func newPerson(name, city string, age int8) *Person {
	return &Person{
		name: name,
		city: city,
		age:  age,
	}
}

//方法和接收者
//Go语言中的方法（Method）是一种作用于特定类型变量的函数。
//这种特定类型变量叫做接收者（Receiver）。接收者的概念就类似于其他语言中的this或者 self。

//	func (接收者变量 接收者类型) 方法名(参数列表) (返回参数) {
//	   函数体
//	} *****
func (p Person) Dream() {
	fmt.Printf("%s的梦想是学好go语言\n", p.name) //张三的梦想是学好go语言
}

// 指针类型的接收者（可以修改变量本身）
// 新建一个方法修改年龄
func (p *Person) setAge(newAge int8) {
	p.age = newAge
}

// 值类型的接收者（无法修改变量本身，修改操作只是针对副本）
func (p Person) setCiy(newCity string) {
	p.city = newCity
}
func main() {
	//调用构造函数
	p1 := newPerson("张三", "沙河", 90)
	fmt.Printf("%#v\n", p1) //&main.person{name:"张三", city:"沙河", age:90}

	//调用方法
	p1.Dream()
	p1.setAge(30)       //修改年龄
	fmt.Println(p1.age) //30
	p1.setCiy("北京")
	fmt.Println(p1.city) //沙河 ，仍然是沙河，没有变化
}
