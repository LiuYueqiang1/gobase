package main

import (
	"fmt"
)

// 结构体的继承和方法
//
//	func (接收者变量 接收者类型) 方法名(参数列表) (返回参数) {
//		函数体
//	}
type animal struct {
	Name string
	age  int
}
type dog struct {
	Feet   int
	Animal *animal
}

func (a *animal) move() {
	fmt.Printf("%s会动\n", a.Name)
}
func (d *dog) wang() {
	fmt.Printf("%s会汪汪汪\n", d.Animal.Name)
}

func main() {
	d1 := &dog{
		Feet: 4,
		Animal: &animal{
			Name: "乐乐",
			age:  2,
		},
	}
	d1.Animal.move()
	d1.wang()
}
