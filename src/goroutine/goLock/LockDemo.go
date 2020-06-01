package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	num   = 0
	count = 0
	lock  sync.Mutex
)

func add() {
	lock.Lock() //可以通过加锁的方式来保证线程安全问题
	for i := 1; i <= 1000; i++ {
		num++
	}
	count++
	fmt.Println(count)
	lock.Unlock() //记得要解锁
}

func main() {
	for i := 0; i < 100; i++ {
		go add()
	}
	time.Sleep(time.Second * 5)
	fmt.Println(num)
}
