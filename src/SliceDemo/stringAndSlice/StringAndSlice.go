package main

import "fmt"

//Go的string底层其实就是一个byte数组，所以string也可以通过slice进行操作
func main() {

	s := "hello,world"
	fmt.Println("string =", s)
	fmt.Println("length of string :", len(s))
	fmt.Printf("address of string : %v\n", &s)
	fmt.Println()

	//直接通过以下方式是无法将string转换为切片的
	//slice :=s[1:5]
	//string转换成成切片有两种方式[]unit8(str)或者[]rune，unit8其实就是byte，rune就是int32
	//两者的不同就是byte是按照字节切片，所以无法处理中文（一个中文占三个字节），而rune的数据范围比较大，相当于按字符切片，可以用来处理中文
	slice := []uint8(s)
	fmt.Println("slice =", slice)
	fmt.Println("length of slice :", len(slice))
	fmt.Printf("type of slice : %T", slice)
	//对于string类型，切片所指向的数组并非原来string指向的，而是创建了一个新的数组，然后将string中的数据copy到新的数组当中
	fmt.Printf("address of slice : %p %p\n", &slice, slice)
}
