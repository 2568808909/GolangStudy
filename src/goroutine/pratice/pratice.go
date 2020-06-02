package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

type tuple struct {
	key   int
	value int
	no    int
}

var (
	intChan        = make(chan int, 20000)
	resChan        = make(chan tuple, 20000)
	countDownLatch = sync.WaitGroup{} //实现类似Java中CountDownLatch的效果
)

//将1-2000放入intChan
func putDataIntoChannel() {
	for i := 1; i <= 20000; i++ {
		intChan <- i
	}
	fmt.Println("-----------", len(intChan))
	close(intChan)
}

func calculation(n int) (res int) {
	for i := 1; i <= n; i++ {
		res += i
	}
	return
}

//从intChan中取出数据，并写入到resChan中
func process(no int) {
	//在协程执行的方法中就要使用recover对错误进行处理，避免错误导致程序终止
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	for n := range intChan {
		res := calculation(n)
		resChan <- tuple{n, res, no}
	}
	countDownLatch.Done() //countDown
}

func main() {
	countDownLatch.Add(8)                //相当于 new CountDownLatch(8)
	cpu := runtime.NumCPU()              //获取系统的cpu核数
	runtime.GOMAXPROCS(cpu)              //设置计算使用的最大核数
	start := time.Now().UnixNano() / 1e6 //获取时间戳(纳秒->毫秒)
	go putDataIntoChannel()
	for i := 0; i < 8; i++ {
		go process(i)
	}
	//go process(0)
	countDownLatch.Wait() //countDownLatch.await()
	close(resChan)        //关闭resChan防止死锁
	count := 0
	for v := range resChan {
		count++
		fmt.Printf("%d : %v\n", count, v)
	}
	end := time.Now().UnixNano() / 1e6
	fmt.Println("cost time : ", end-start, "ms")
}
