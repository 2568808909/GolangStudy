package main

import "fmt"

//Golang中对于数据的转换都没有隐式转换，也就是说，无论是大数据转小数据，还是小数据转大数据，
//都要通过显示转换来完成，显示转换的语法为T(v),T为要转换的数据类型，v为要转换的变量
func main() {
	var i int32 = 999
	var f = float32(i) //var f float32 = float32(i)  前面的float32可以省略，因为后面赋值已经知道其类型了
	var b = byte(i)
	fmt.Printf("i=%d f=%f b=%d", i, f, b)

	var k int64 = 999
	k = k + int64(i) //在进行运算时，也要显示进行转换

}
