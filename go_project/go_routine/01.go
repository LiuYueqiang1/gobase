package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func hello(i int) {
	defer wg.Done() // 通知计数器-1
	fmt.Println("hello 纳扎", i)

}

func main() { //开启一个主 goroutine 去执行这个函数

	//wg.Add(1)  //计数+1
	for i := 0; i < 10; i++ {
		wg.Add(1)   //计数+1
		go hello(i) //开启一个 goroutine 执行hello这个函数
	}

	fmt.Println("hello main")
	//time.Sleep(time.Second)   程序等待一秒钟
	wg.Wait() //  等所有的小弟干完活后收兵
}
