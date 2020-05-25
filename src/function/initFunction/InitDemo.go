package main

import (
	"fmt"
	//虽然util包下已经导入过fmt包，但是这个包要使用fmt包中的函数依旧要重新导入，不能传递
	"function/util" //导入其他包时，会对其他包进行初始化操作(也就是调用他们的init函数)，这些初始化操作会放在在本包的一切操作之前
)

//以这种方式对字段(属性)进行初始化，其执行顺序在init方法之前
var field int = test()

func test() int {
	fmt.Println("调用test方法")
	return 0
}

//init函数时在此包初始化时执行，进行一些属性的初始化操作，类似于构造方法
func init() {
	fmt.Println("调用init方法")
}

//main函数在init函数执行后在执行
func main() {
	fmt.Println("调用main方法")
	field += 0
	fmt.Println(util.Age)
}
