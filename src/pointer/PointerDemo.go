package main

import "fmt"

func main() {
	var i int = 10
	fmt.Printf("value of i : %d , address of i %v\n", i, &i) //&i表示的i的地址

	var ptr *int = &i
	fmt.Printf("value of ptr : %v , address of ptr %v\n", ptr, &ptr) //指针存的是一个地址，存地址的空间也有地址
	fmt.Printf("point to value : %d\n", *ptr)

	*ptr = 50
	fmt.Printf("point to value : %d, value of i : %d\n", *ptr, i) //通过指针来修改i变量内存中的值，i的值会发生变化
}
