package main

import (
	"bufio"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
)

func writeFile01() {
	//写入文件时，默认会清空文件再写入，如果想追加写入，则在打开文件时需要选择os.O_APPEND
	//这种方式写入文件需要文件存在，否则会报错
	file, e := os.OpenFile("C:\\Users\\Administrator\\Desktop\\杂七杂八的文档\\test.txt", os.O_RDWR|os.O_APPEND, 0666)
	if e == nil {
		defer func() {
			e := file.Close()
			if e != nil {
				panic(e)
			}
			fmt.Println("file already close")
		}()
		writer := bufio.NewWriter(file)
		msg := "hello,Golang 任我行!\r\n"
		for i := 0; i < 10; i++ {
			_, e := writer.WriteString(msg)
			if e != nil {
				fmt.Println(e)
			}
		}
		e := writer.Flush()
		if e != nil {
			fmt.Println("error message :", e)
		}
		fmt.Println("completed")
	} else {
		fmt.Println("error message :", e)
	}
}

//这种方式可以实现文件的拷贝，ioutil写入文件时，不需要文件存在
func copyFile() {

	bytes, e := ioutil.ReadFile("C:\\Users\\Administrator\\Desktop\\杂七杂八的文档\\test.txt")
	if e == nil {
		e := ioutil.WriteFile("C:\\Users\\Administrator\\Desktop\\杂七杂八的文档\\abcd.txt", bytes, 0666)
		if e != nil {
			fmt.Println(e)
		}
	} else {
		fmt.Println(e)
	}
}

//判断文件是否存在
func FileExists(path string) (bool, error) {
	_, e := os.Stat(path)
	if e == nil {
		return true, nil
	} else if os.IsNotExist(e) { //文件路径不存在
		return false, nil
	}
	return false, errors.New("出现不确定错误")
}

func main() {
	b, e := FileExists("C:\\Users\\Administrator\\Desktop\\杂七杂八的文档\\ab.txt")
	fmt.Println(b)
	fmt.Println(e)
}
