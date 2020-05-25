package main

import "fmt"

//Go函数写法 func 函数名(参数列表) (返回值)，返回值可以有多个
func add(a int, b int) int {
	return a + b
}

//可以为返回值命名，这样返回的时候会按照返回值列表顺序返回数据，并不需要明确写出来
func cal(a int, b int) (sum int, sub int) {
	sum = a + b
	sub = a - b
	return //此处不写默认按照返回值列表顺序返回，当然也可以自己自定义返回顺序
}

//以上写法相当于以下写法
//func cal(a int, b int) (int, int) {
//	sum := a + b
//	sub := a - b
//	return sum, sub
//}

//Go中不允许重载，在同一个包中，不允许有两个及以上同名的函数
//func add(a float64,b float64,operation byte) float64{
//	return a-b
//}

func test(f myFunc, a int, b int) int {
	return f(a, b)
}

type myFunc func(int, int) int //给函数类型自定义类型
type myInt int

func main() {
	a, b := 50, 60
	fmt.Println("res=", add(a, b))

	var f myFunc

	f = add //在Go中函数也可作为变量来使用
	res := f(123, 654)
	fmt.Println("res=", res)

	//函数也可作为参数传递给另一个函数
	res = test(f, 95, 64)
	res = test(add, 95, 64)
	fmt.Println("res=", res)
}
