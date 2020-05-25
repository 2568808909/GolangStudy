package main

import "fmt"

//该函数的返回值是一个函数，而返回的这个函数把外面函数的n加入到自己的运算中，函数和这个变量就形成了一个整体，也就是一个闭包返回到外部
//闭包：闭包就是一个函数与其现骨干的引用环境相结合形成的一个整体
func addUpdater() func(int) int {
	n := 10
	return func(x int) int {
		n += x
		return n
	}
}

func main() {
	f := addUpdater()
	//f其实就是一个闭包，f中包含着一个n和一个匿名函数，多次操作n的值会叠加
	//这其实有点类似于面型对象的类，调用addUpdater返回的就是一个对象，n是一个字段,而匿名函数就是其方法
	fmt.Println(f(1))
	fmt.Println(f(2))
	fmt.Println(f(3))
}
