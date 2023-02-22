package main

import "fmt"

func test1() {
	a := 10
	b := &a

	fmt.Printf("a:%d ptr:%p\n", a, &a)
	fmt.Printf("b:%p type:%T\n", b, b)
	fmt.Println(&b)
}

func modify1(x int) {
	x = 100
}
func modify2(x *int) {
	*x = 100
}

//new 函数  func new(Type) *Type

//make 函数
func main() {
	test1()

	a := 10
	modify1(a)
	fmt.Println(a) //10

	modify2(&a)
	fmt.Println(a) //100

	b := new(int)
	c := new(bool)
	fmt.Println(b)        //0xc00001a0d8
	fmt.Println(c)        //0xc00001a0e0
	fmt.Printf("%T\n", b) //*int
	fmt.Printf("%T\n", c) //*bool
	fmt.Println(*b)       //0
	fmt.Println(*c)       //false

	var e *int
	e = new(int)
	*e = 10
	fmt.Println(*e)

	var f map[string]int
	f = make(map[string]int, 10)
	f["大海"] = 5
	fmt.Println(f)
}
