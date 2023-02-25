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
	fmt.Printf("%s会喵喵喵~~~\n", c.Name)
}

type sayer interface {
	say()
}

func da(arg sayer) {
	arg.say()
}
func main() {
	//c1 := cat{
	//	Name: "米粒",
	//}
	//da(c1)
	//d1 := dog{}
	//da(d1)
	var s sayer
	c2 := cat{
		Name: "米粒",
	}
	s = c2
	s.say()        //米粒会喵喵喵~~~
	da(s)          //米粒会喵喵喵~~~
	fmt.Println(s) //{米粒}

	d2 := dog{}
	s = d2
	da(s) //汪汪汪

	//米粒会喵喵喵~~~
	//{米粒}
	//汪汪汪
}
