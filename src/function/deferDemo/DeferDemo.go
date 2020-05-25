package main

import "fmt"

func demo(a int, b int) int {
	//defer会将其后面的语句压入一个栈中，并不会立刻执行，在函数结束后，才会从栈中取出语句执行，执行顺序为先入后出
	//另外，相关变量的值也会压入栈中，即使之后对相关变量进行修改，从栈取出而执行的语句的变量的值都会是原来的值
	//defer一般是用于释放连接/资源，类似于java的finally块的功能
	defer fmt.Println("a=", a)
	defer fmt.Println("b=", b)

	a++
	b++
	res := a + b
	fmt.Println("res=", res)
	return res
}

func main() {
	res := demo(20, 25)
	fmt.Println("res=", res)
}
