package main

import "fmt"

func main() {
	//ch := make(chan int, 1)
	var ch chan int
	//ch =make(chan int) //无缓冲区通道
	ch = make(chan int, 1) //带缓冲区通道
	ch <- 10
	fmt.Println(ch)
	fmt.Println("发送成功！")
	close(ch)
}
