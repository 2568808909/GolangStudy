package util

import "fmt"

var Name string
var Age int

func init() {
	Name = "ccb"
	Age = 22
	fmt.Println("调用util包下的init方法")
}
