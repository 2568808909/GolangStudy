package main

import "fmt"

//Go的string底层其实就是一个byte数组，所以string也可以通过slice进行操作
func main() {

	s := "hello,world"
	fmt.Println("string =", s)
	fmt.Println("length of string :", len(s))
	fmt.Println()

	//直接通过以下方式是无法将string转换为切片的
	//slice :=s[1:5]
	//string转换成成切片有两种方式[]unit8(str)或者[]rune，unit8其实就是byte，rune就是int32
	//两者的不同就是byte是按照字节切片，所以无法处理中文（一个中文占三个字节），而rune的数据范围比较大，相当于按字符切片，可以用来处理中文
	slice := []uint8(s)
	fmt.Println("slice =", slice)
	fmt.Println("length of slice :", len(slice))
	fmt.Printf("type of slice : %T", slice)
}
