package main

import (
	"fmt"
	"strings"
)

//闭包 =函数+引用环境

func a() func() {
	name := "基本的"
	return func() {
		fmt.Println("闭包", name)
	}
}

func d(name string) func() {
	return func() {
		fmt.Println(name)
	}
}

func add(x int) func(int) int {
	return func(y int) int {
		x = x + y
		return x
	}
}
func add1() func(int) int {
	var x int
	return func(y int) int {
		x = x + y
		return x
	}
}

func makeSuffixFunc(suffix string) func(string) string {
	return func(name string) string {
		if !strings.HasSuffix(name, suffix) {
			return name + suffix
		}
		return name //意思时如果有后缀的话直接返回name
	}
}

// 两个函数作为返回值
func calc(base int) (func(int) int, func(int) int) {

	add := func(a int) int {
		base += a
		return base
	}

	sub := func(b int) int {
		base -= b
		return base
	}
	return add, sub
}

func main() {

	b := a() //b此时就是一个闭包，调用函数a，将返回值函数赋给b
	b()      // b执行闭包里面的函数

	e := d("闭包传参") //闭包在调用时传参
	e()            //调用闭包里的函数

	f := add(10)       //在add()中传入了一个初始值10
	fmt.Println(f(20)) //30
	fmt.Println(f(30)) //60

	g := add1()        //这个对比于add()相当于什么都没传入 x是一个nil
	fmt.Println(g(20)) //20
	fmt.Println(g(50)) //70

	jpgFunc := makeSuffixFunc(".jpg")
	txtFunc := makeSuffixFunc(".txt")
	fmt.Println(jpgFunc("test")) //test.jpg
	fmt.Println(txtFunc("test")) //test.txt

	h1, h2 := calc(10)
	m1 := h1(2)
	m2 := h2(2)
	fmt.Println(m1) // 12
	fmt.Println(m2) // 10
}
