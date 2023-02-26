package main

import (
	"fmt"
	"sync"
	"time"
)

//读写互斥锁 程序中读的次数远远大于写，且读的时候不改变内容

var (
	x1     int64
	wg4    sync.WaitGroup
	lock   sync.Mutex   //此锁在读的时候也不允许其它进程读写
	rwLock sync.RWMutex //读写锁
)

func read() {
	//lock.Lock()
	rwLock.RLock()
	time.Sleep(time.Millisecond)
	//lock.Unlock()
	rwLock.RUnlock()
	wg4.Done()
}
func write() {
	//lock.Lock()
	rwLock.Lock() //写的时候也不允许读
	x1 += 1
	time.Sleep(time.Millisecond * 10)
	//lock.Unlock()
	rwLock.Unlock()
	wg4.Done()
}

func main() {
	start := time.Now()
	for i := 0; i < 1000; i++ {
		wg4.Add(1)
		go read()
	}
	for i := 0; i < 10; i++ {
		wg4.Add(1)
		go write()
	}

	wg4.Wait()
	fmt.Println(time.Now().Sub(start))
}
