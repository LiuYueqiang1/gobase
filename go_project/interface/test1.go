package main

import (
	"fmt"
)

type dog struct {
}

func (d dog) say() {
	fmt.Println("汪汪汪")
}

type cat struct {
	Name string
}

func (c cat) say() {
	fmt.Println("喵喵喵~~~")
}

type sayer interface {
	say()
}

func da(arg sayer) {
	arg.say()
}
func main() {
	c1 := cat{
		Name: "米粒",
	}
	da(c1)
	d1 := dog{}
	da(d1)

}
