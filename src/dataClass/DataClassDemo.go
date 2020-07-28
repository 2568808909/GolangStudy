package main

import (
	"golang.org/x/tools/go/ssa/interp/testdata/src/fmt"
)

//Golang中数据类型分为两种，值类型和引用类型
//值类型：基本数据类型,int系列，float系列，bool，string，数组和struct（结构体）
//引用类型：指针，slice(切片)，map，channel(管道)，interface（接口）
//两者区别：
//值类型变量直接存储值，内存通常在栈中分配
//引用类型：变量存储的是一个地址，这个地址对应的空间才真正存放数据，内存通常在堆上进行分配，当没有任何变量	引用这个地址时，改地址所对应的空间就会变认定为垃圾，会被GC回收
//逃逸分析：有可能在编译时将一些值类型放入堆，一些引用类型放入栈区
func main() {
	fmt.Print("")
}
