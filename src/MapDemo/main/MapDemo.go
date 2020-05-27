package main

import (
	"fmt"
)

func main() {

	//map在声明后为nil，并不能立刻使用，要使用make函数进行空间的分配后才可以
	var m map[string]string
	m = make(map[string]string, 3)
	m["name"] = "tom"
	m["age"] = "123"
	m["balance"] = "1000.0"
	m["sex"] = "male"
	fmt.Println("length of map : ", len(m))
	fmt.Println("map :", m)
	fmt.Println()

	//可以在使用make给map分配空间时，给map定义键值对
	//m = map[string]string{"name":"jack","age":"22"}

	//删除一个元素
	delete(m, "sex")
	fmt.Println("length of map : ", len(m))
	fmt.Println("map :", m)
	fmt.Println()

	//Go中的map没有提供一次性全部清空的功能，要清空map有两种方式
	//第一种，通过for-range遍历map，逐一删除键值对
	//for key, value := range m {
	//	fmt.Println("key =", key, "value =", value)
	//	delete(m, key)
	//}
	//fmt.Println("length of map : ", len(m))
	//fmt.Println("map :", m)
	//fmt.Println()
	//第二种，通过make函数重新给map分配空间，原来的空间失去引用，之后会被GC回收
	m = make(map[string]string, 10)

}
