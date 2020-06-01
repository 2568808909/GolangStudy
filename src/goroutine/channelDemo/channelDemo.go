package main

import (
	"fmt"
	"time"
)

//channel是值类型数据，传入的参数属于引用传递
func test(channel chan int) {
	time.Sleep(time.Second * 2)
	fmt.Println(<-channel)
}

//为什么需要使用channel作为同步机制
//1.主线程等待所有goroutine都完成的时间难以确定，时间短了可能还有协程未完成任务，时间长了可能所有协程早就完成任务，有造成时间的浪费
//2.通过加锁实现同步，并不利于协程对共享资源的读写操作，并发度（效率）会将降低
func main() {
	//channel本质就是一个队列（先进先出），并且是线程安全的，对其操作不需要加锁
	//channel是有类型的，下面定义的channel就只能存放int类型数据
	var intChan chan int
	intChan = make(chan int, 3) //channel需要分配空间后才可使用，并且，其容量是不可动态扩展的，这里定义了其容量为3，
	intChan <- 10               //将元素放入channel中
	intChan <- 5
	intChan <- 20
	fmt.Printf("length of channel : %d , capacity of channel : %d\n", len(intChan), cap(intChan))
	fmt.Println(<-intChan) //从channel中取出元素
	intChan <- 0
	//如果这里开启了另一个协程去消费channel中的数据后，就能使得主线程不在阻塞
	go test(intChan)
	fmt.Println("main put int channel")
	intChan <- 50 //fatal error: all goroutines are asleep - deadlock! 一开始给channel分配的容量为3，放入多于3个数据就会死锁(其实就是阻塞，但由于没有其他线程/协程消费channel的数据，实际上就是出于死锁的状态)
	fmt.Println("put success")
	time.Sleep(time.Second * 4)

	//遍历channel。值得注意的是，如果没有关闭channel，使用这种方式遍历，会一直向channel中取数据，没有数据时就会死锁(阻塞)
	//关闭channel后，就不会有以上问题
	//for v := range intChan {
	//	fmt.Println(v)
	//}

	//关闭channel。关闭后不能往channel内写数据，但是可以读数据
	close(intChan)
	//fmt.Println(<-intChan)
	//fmt.Println(<-intChan)
	//intChan<-9 //panic: send on closed channel
	for v := range intChan {
		fmt.Println(v)
	}

}
