package main

import (
	"fmt"
)

// func 函数名(参数)(返回值){
//     函数体
// }

func sum1(x int, y int) int {
	return x + y
}
func sayhello1() {
	fmt.Println("hello")
}

// 可变参数
// 可变参数是指函数的参数数量不固定。Go语言中的可变参数通过在参数名后加...来标识。
// 注意：可变参数通常要作为函数的最后一个参数。
func intsum2(x ...int) int {
	//fmt.Println(x) //x是一个切片
	sum := 0
	for _, v := range x {
		sum = sum + v
	}
	return sum
}

// 固定参数搭配可变参数
func intsum3(x int, y ...int) int {

	sum := x
	for _, v := range y {
		sum = sum + v
	}
	return sum
}

// 多返回值
func calc(x int, y int) (int, int) {
	sum := x + y
	sub := x - y
	return sum, sub
}

// 给返回值命名
func calc2(x int, y int) (sum int, sub int) {
	sum = x + y
	sub = x - y
	return
}
func main() {

	a := sum1(2, 3)
	fmt.Println(a)
	sayhello1()
	fmt.Println("可变参数")
	b := intsum2(10, 20, 30)
	fmt.Println(b)
	fmt.Println("固定+可变参数")
	c := intsum3(100, 20, 30, 520)
	fmt.Println(c)
	fmt.Println("多返回值")
	d, e := calc(20, 10)
	fmt.Println(d, e)
	fmt.Println("返回值命名")
	f, g := calc2(30, 20)
	fmt.Println(f, g)
}
