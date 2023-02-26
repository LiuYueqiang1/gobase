package main

import (
	"fmt"
	"sync"
)

var (
	x   int64
	wg3 sync.WaitGroup
	lo  sync.Mutex //互斥锁  它能够保证同一时间只有一个 goroutine 可以访问共享资源
)

func add() {
	for i := 0; i < 5000; i++ {
		lo.Lock()
		x = x + 1
		lo.Unlock()
	}
	wg3.Done()
}

func main() {
	wg3.Add(2)
	go add()
	go add()
	wg3.Wait()
	fmt.Println(x)
}
