package main

import (
	"errors"
	"fmt"
)

func readConf(name string) error {
	if name == "my.conf" {
		return nil
	}
	return errors.New("配置信息错误..") //通过errors.New可以自定义错误
}

func cal() {
	err := readConf("my.conf6")
	if err != nil {
		panic(err) //读取文件错误，输出这个错误并终止程序
	}
	fmt.Println("calculation complete")
}

func main() {
	cal()
}
