package main

import "fmt"

func f1() int {
	x := 5
	defer func() {
		x++
	}()
	return x
} //5  返回值是int 类型，就是x

func f2() (x int) {
	defer func() {
		x++
	}()
	return 5
} //6  函数返回值为x，这个函数返回了5，x++，所以为6

func f3() (y int) {
	x := 5
	defer func() {
		x++
	}()
	return x
} //5
func f4() (x int) {
	defer func(x int) {
		x++
	}(x)
	return 5
} //5

func calc(index string, a, b int) int {
	ret := a + b
	fmt.Println(index, a, b, ret)
	return ret
}

func funcA() {
	fmt.Println("func A")
}
func funcB() {
	defer func() {
		err := recover()
		if err != nil { //err不是空的，也就是说有错误且把recover赋给了err，则输出下面的
			fmt.Println("recover in B")
		}
	}()
	panic("panic in B")
}
func funcC() {
	fmt.Println("func C")
}

func main() {

	// fmt.Println(1)
	// defer fmt.Println("defer")
	// fmt.Println(2)
	// fmt.Println(3)

	// x := 1
	// y := 2
	// defer calc("AA", x, calc("A", x, y))
	// x = 10
	// defer calc("BB", x, calc("B", x, y))
	// y = 20

	funcA()
	funcB()
	funcC()
}
