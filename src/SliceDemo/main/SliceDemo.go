package main

import "fmt"

//切片(slice)是数组的一个引用，属于引用类型，与数组不同的是，切片的长度是可以动态变化的
//切片从底层来看，其实是一个结构体，包含着三个属性，一个数组的指针，一个长度len，一个容量cap
func main() {
	//创建切片的三种方法，后面两种方法创建的切片并没有提供对数组的直接引用，只能通过切片来操作数组
	var arr = [5]int{5, 2, 3, 9, 4}
	//1.在某个数组的基础上创建切片，含头不含尾
	s := arr[1:3]
	fmt.Println("slice =", s)
	fmt.Println("length of slice :", len(s))
	fmt.Println("capacity of slice :", cap(s))
	for index, value := range s {
		fmt.Println("index =", index, "value =", value)
	}
	fmt.Println()
	//2.使用make函数创建,传入参数分别为切片类型，切片长度，切片容量
	s = make([]int, 5, 10)
	s[1] = 30
	s[4] = 24
	fmt.Println("slice =", s)
	fmt.Println("length of slice :", len(s))
	fmt.Println("capacity of slice :", cap(s))
	fmt.Println()
	//3.用类似定义数组的方式来定义一个切片
	s = []int{5, 6, 2, 0}
	fmt.Println("slice =", s)
	fmt.Println("length of slice :", len(s))
	fmt.Println("capacity of slice :", cap(s))
}
