package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg2 sync.WaitGroup

func A() {
	for i := 0; i < 10; i++ {
		fmt.Println("a", i)
	}
	wg2.Done()
}
func B() {
	for i := 0; i < 10; i++ {
		fmt.Println("b", i)
	}
	wg2.Done()
}

func main() {

	runtime.GOMAXPROCS(32) //32个OS线程来执行Go代码
	wg2.Add(2)
	go A()
	go B()
	wg2.Wait()

}
