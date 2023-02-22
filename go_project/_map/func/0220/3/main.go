package main

import "fmt"

//匿名函数
// func(参数)(返回值){
//     函数体
// }

func main() {
	// 将匿名函数保存到变量
	sayHello := func() {
		fmt.Println("匿名函数")
	}
	sayHello()
	//自执行函数：匿名函数定义完加()直接执行
	func() {
		fmt.Println("直接调用匿名函数")
	}()
	// 将匿名函数保存到变量
	add := func(x, y int) {
		fmt.Println(x + y)
	}
	add(2, 3) //5
	//自执行函数：匿名函数定义完加()直接执行
	func(x, y int) {
		fmt.Println(x - y)
	}(5, 3) //2
	fmt.Println("闭包：")
}
