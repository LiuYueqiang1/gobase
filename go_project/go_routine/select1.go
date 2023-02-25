package main

import "fmt"

// 可处理一个或多个 channel 的发送/接收操作。
// 如果多个 case 同时满足，select 会随机选择一个执行。
// 对于没有 case 的 select 会一直阻塞，可用于阻塞 main 函数，防止退出。
func main() {
	ch := make(chan int, 1)
	for i := 0; i < 10; i++ {
		select {
		case x := <-ch:
			fmt.Println(x)
		case ch <- i:
		default:
			fmt.Println("空操作")
		}
	}
}

//给一个拿一个，有空了再给一个，再取一个
//0
//2
//4
//6
//8
