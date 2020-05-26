package main

import "fmt"

func cal() {
	//defer+recover配合可以获取到错误信息，并对该错误进行处理，另外，这样处理后，程序遇到错误也不会崩溃而结束，而是跳出有问题的方法，然后继续运行
	defer func() {
		if err := recover(); err != nil { //recover可以获取到现有的错误信息
			fmt.Println("err=", err)
		}
	}()
	a := 10
	b := 0
	res := a / b //这里不能直接用常量，直接用常量编译器在编译时就能发现这里有错误，编译不会通过
	fmt.Println("calculation completed... res=", res)
}

func main() {
	cal()
	fmt.Println("program completed...")
}
