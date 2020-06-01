package main

import (
	"fmt"
	"strconv"
	"time"
)

var num = 0

func test() {
	for i := 1; i <= 10; i++ {
		fmt.Println("test() hello test..." + strconv.Itoa(i))
		time.Sleep(time.Second)
	}
}

func add() {
	for i := 1; i <= 1000; i++ {
		num++
	}
}

//Go中协程特点
//有独立的栈空间
//共享程序堆空间
//调度由用户控制
//协程是轻量级的线程
func main() {
	//go test() //开启一个协程
	//for i := 1; i <= 10; i++ {
	//	fmt.Println("main() hello golang..." + strconv.Itoa(i))
	//	time.Sleep(time.Second)
	//}

	//并发修改共享资源会引发线程安全问题，结果不确定
	for i := 0; i < 100; i++ {
		go add()
	}

	time.Sleep(time.Second * 10)
	fmt.Println(num)
}
