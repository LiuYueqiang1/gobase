package main

import "fmt"

func f1(ch chan<- int) {
	for i := 0; i < 10; i++ {
		ch <- i
	}
	close(ch)
}
func f2(ch1 <-chan int, ch2 chan<- int) { //只能从通道接收值，只能向通道发送值
	//for {
	//	tmp, ok := <-ch1 //从ch中接收值
	//	if !ok {
	//		break
	//	}
	//	ch2 <- tmp * tmp //把值发送到ch中
	//}
	for temp := range ch1 {
		ch2 <- temp * temp
	}

	close(ch2)
}

func main() {
	ch1 := make(chan int, 10)
	ch2 := make(chan int, 20)
	go f1(ch1)
	go f2(ch1, ch2)
	for ret := range ch2 {
		fmt.Println(ret)
	}
}
