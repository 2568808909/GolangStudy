package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func readFile01() {
	//os.Open(name string) (file *File ,err error) 这个方法会以只读的方式打开文件，并返回一个文件指针和错误对象
	//file, err := os.Open("C:\\Users\\Administrator\\Desktop\\我要干什么来着.txt")

	//func OpenFile(name string, flag int, perm FileMode) (file *File, err error) 是一个更一般性的文件打开函数。
	//flag决定用什么方式打开文件，如只读，只写，读写，追加写等
	//perm决定文件的权限，在windows下不起作用，只有在linux或者unix下才起作用
	file, err := os.OpenFile("C:\\Users\\Administrator\\Desktop\\我要干什么来着.txt", os.O_RDONLY, 0666)
	if err == nil {
		//使用defer延迟关闭文件
		defer func() {
			e := file.Close()
			if e != nil {
				panic(e)
			}
		}()
		//以缓存形式读取文件
		reader := bufio.NewReader(file)
		var line string
		var err error
		for {
			line, err = reader.ReadString('\n')
			if err == io.EOF {
				break
			}
			fmt.Println(line)
		}
	}
}

//通过ioutil下的函数，可以方便的一次性读取一个文件内的所有内容，打开文件和关闭文件的操作都封装在函数内了，用户并不需要关心
//不过这种方式只适合读取比较小的文件，因为过大的文件不能全部放入内存中
func readFile02() {
	bytes, e := ioutil.ReadFile("C:\\Users\\Administrator\\Desktop\\我要干什么来着.txt")
	if e == nil {
		fmt.Printf("%s", bytes)
	}
}

func main() {
	readFile01()
}
