package main

import "fmt"

//切片(slice)是数组的一个引用，属于引用类型，与数组不同的是，切片的长度是可以动态变化的
//切片从底层来看，其实是一个结构体，包含着三个属性，一个数组的指针，一个长度len，一个容量cap
func main() {
	//创建切片的三种方法，后面两种方法创建的切片并没有提供对数组的直接引用，但数组时存在的，只是对程序员不可见，只能通过切片来操作数组
	var arr = [5]int{5, 2, 3, 9, 4}
	fmt.Println(&arr)
	//1.在某个数组的基础上创建切片，含头不含尾
	s := arr[:3]
	//若从0开始0可以省略
	//s :=arr[:4]  //相当于arr[0:4]
	//同理如果是渠道最后一个也可以省略
	//s :=arr[1:]  //相当于arr[1:len(arr)]
	//s :=arr[:]   //相当于arr[0:len(arr)]
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
	//对于以下操作虽然切片的容量为10，但目前长度为5，这样还是会报下标越界异常
	//s[6] = 50
	fmt.Println("slice =", s)
	fmt.Println("length of slice :", len(s))
	fmt.Println("capacity of slice :", cap(s))
	fmt.Println()
	//3.用类似定义数组的方式来定义一个切片
	s = []int{5, 6, 2, 0}
	fmt.Println("slice =", s)
	fmt.Println("length of slice :", len(s))
	fmt.Println("capacity of slice :", cap(s))
	fmt.Println()

	//切片可以继续切片
	s = s[1:3]
	fmt.Println("slice =", s)
	fmt.Println("length of slice :", len(s))
	fmt.Println("capacity of slice :", cap(s))
	fmt.Println()

	//使用append内置函数对切片进行扩容，如果有足够的capacity，则会重新切片以容纳新的元素，否则，则分配一个新的数组。append会返回跟新后的切片
	slice := append(s, 5)
	slice[1] = 100
	fmt.Println("slice =", slice)
	fmt.Println("length of slice :", len(slice))
	//此次append因为容量充足，所以不需要创建新的数组，slice和s存放的地址都是一致的，也就是说指向的是同一个数组；但是slice和s不属于同一个切片，他们的长度和存放数据的地址都不同
	fmt.Printf("address of slice : %p %p\n", &slice, slice)
	fmt.Println("capacity of slice :", cap(slice))
	fmt.Println()
	//append后原来的切片并不会发生改变
	fmt.Println("slice =", s)
	fmt.Println("length of slice :", len(s))
	fmt.Printf("address of slice : %p %p\n", &s, s)
	fmt.Println("capacity of slice :", cap(s))
	fmt.Println()
	//append也可以使用其他的切片来追加(...是语法规则必须要带上)，这里只能用切片来进行追加，不能是数组
	slice = append(slice, s...)
	fmt.Println("slice =", slice)
	fmt.Println("length of slice :", len(slice))
	fmt.Printf("address of slice : %p %p\n", &slice, slice)
	fmt.Println("capacity of slice :", cap(slice))
	fmt.Println()

	//切片的拷贝
	slice = make([]int, 5, 10)
	copy(slice, s)
	fmt.Println("slice =", slice)
	fmt.Println("length of slice :", len(slice))
	fmt.Printf("address of slice : %p %p\n", &slice, slice)
	fmt.Println("capacity of slice :", cap(slice))
	fmt.Println()
}
