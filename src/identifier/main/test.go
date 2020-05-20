package main

import (
	"fmt"
	"identifier/model" //导入自己写的包，需要配置GOPATH
)

func main() {
	fmt.Println(model.StudentAge) //包外只能访问首字母大写的属性
	model.Show()                  //通过这个方法访问包内所有属性
}
