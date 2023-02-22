package main

import (
	"errors"
	"fmt"
)

//函数进阶

// 我们可以使用type关键字来定义一个函数类型，具体格式如下：
// type calculation func(int, int) int
// 上面语句定义了一个calculation类型，它是一种函数类型，这种函数接收两个int类型的参数并且返回一个int类型的返回值。
func add(x, y int) int {
	return x + y
}
func sub(x, y int) int {
	return x - y
}

// 高阶函数分为函数作为参数和函数作为返回值两部分。
// 函数作为参数传入另一个函数
func calc(x, y int, op func(z int, h int) int) int { //注意函数内部的写法
	return op(x, y)
}

func calc2(x, y int, op func(int, int) int) int { //注意函数内部的写法
	return op(x, y)
}

// 函数作为返回值
func do(s string) (func(x, y int) int, error) {
	switch s {
	case "+":
		return add, nil
	case "-":
		return sub, nil
	default:
		err := errors.New("无法识别")
		return nil, err
	}
}

func main() {
	//calculation类型
	c := add //直接将函数名add赋值给c
	fmt.Printf("type of c :%T\n", c)
	fmt.Println(c(1, 2))

	f := sub //直接将函数名add赋值给c
	fmt.Printf("type of c :%T\n", f)
	fmt.Println(f(3, 2))

	fmt.Println("函数作为参数传入另一个函数")
	g := calc(10, 20, add)
	fmt.Println(g)

	h := calc2(30, 20, sub)
	fmt.Println(h)

	fmt.Println("函数作为返回值")
	do("+")

}
