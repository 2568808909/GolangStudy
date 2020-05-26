package main

import "fmt"

//数组时在内存中开辟一段连续的数据空间，用于存储类型相同的一组数据
//变量指向的是这个数据空间，而非一个地址，所以数组时值类型
//数组的长度是固定的，创建后就不可改变。
//并且，Go中的数组时值类型数据，也就是说传参时，会对数组的数据进行拷贝，放入该函数的栈中，在函数中的任何操作都不会影响原来函数的值。要想在函数中对传入的函数进行修改则需要使用指针
//值得注意的是，在Go中，数组的长度也是数组类型的一部分，不同长度的数组不能看作同一类型，这在传参时会有体现
func test(arr [5]int) {
	//for-range方式遍历是Go中独有的遍历方式，index代表下标，value是变量的值(这两个变量的名称不是固定的，可以自定义)
	for index, value := range arr {
		fmt.Println("index =", index, "value =", value)
	}
}
func main() {
	//1.指定长度创建数组
	arr := [5]int{} //创建数组时，要明确指明长度
	fmt.Println("array =", arr)
	fmt.Println("length of array :", len(arr))
	fmt.Println("capacity of array :", cap(arr))
	fmt.Println()
	//2. ...可以代替长度，它的值会根据数组元素个数而变化
	arr2 := [...]int{5, 3, 2, 1}
	fmt.Println("array =", arr2)
	fmt.Println("length of array :", len(arr2))
	fmt.Println("capacity of array :", cap(arr2))
	fmt.Println()
	//3.创建数组时可以不按照顺序存放数据
	arr3 := [5]int{1: 3, 0: 1, 4: 6}
	fmt.Println("array =", arr3)
	fmt.Println("length of array :", len(arr3))
	fmt.Println("capacity of array :", cap(arr3))
	fmt.Println()

	//长度也算是数组类型的一部分，arr2的长度为4，test函数的参数指定长度为5，这里调用是会报错的
	//test(arr2)

	test(arr3)
}
