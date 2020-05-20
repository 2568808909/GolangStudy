package main

import "fmt"

func main() { //main方法要放在main包下才可运行，整个包目前看来与真实的目录结构并没有关系
	// Go中没有特定的字符类型，所以表示字符类型数据要用整数类型
	var c1, c2 byte = 'a', '0'
	fmt.Println("通常输出 c=", c1, "c2=", c2)
	fmt.Printf("通常输出 c1=%c c2=%c\n", c1, c2)

	//对于汉字，因为Go中默认使用的编码为UTF-8，一个汉字占用三个字节，所以要使用int类型进行存储
	c3 := '陈'
	fmt.Printf("c3=%c 编码为：%d 类型为：%T", c3, c3, c3)
}
