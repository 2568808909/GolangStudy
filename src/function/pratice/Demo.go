package main

import (
	"fmt"
	"strings"
)

//练习：
//编写一个函数makeString(suffix string)，可以接收一个文件后缀名(如.jpg)，然后返回一个闭包
//调用这个闭包时，会传入一个文件名，如果文件包含指定后缀，则直接返回，如果没有则为文件名拼接上该后缀
func mkString(suffix string) func(string) string {
	return func(name string) string {
		if strings.HasSuffix(name, suffix) {
			return name
		} else {
			return name + suffix
		}
	}

}

func main() {
	make := mkString(".jpg")

	fmt.Println(make("123.jpg"))
	fmt.Println(make("家庭教师"))
}
