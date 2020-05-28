package main

import (
	"encoding/json"
	"fmt"
)

//使用tag的目的主要是定义在序列化时字段的名称，因为在Go中，要使属性在包外也可以访问就需要首字母大写，这对于其他程序员其实并不习惯，就可以使用tag改变序列化时字段的名称
type Teacher struct {
	Name   string  `json:"name"`
	Age    int     `json:"age"`
	Salary float64 `json:"salary"`
}

func main() {
	teacher := Teacher{Name: "金虎", Age: 56, Salary: 6531.12}
	jsonByte, e := json.Marshal(teacher)
	if e == nil {
		//使用tag前 {"Name":"金虎","Age":56,"Salary":6531.12}
		//使用tag后 {"name":"金虎","age":56,"salary":6531.12}
		jsonStr := string(jsonByte)
		fmt.Println(jsonStr)
	} else {
		panic(e)
	}
}
