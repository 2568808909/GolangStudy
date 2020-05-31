package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"jsonDemo/serial"
	"os"
)

func Serial() int {
	monkey := serial.Monster{
		Name:   "齐天大圣孙悟空",
		Age:    640,
		Salary: 15643.14,
	}
	fmt.Println(monkey)
	bytes, err := json.Marshal(monkey)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(bytes))
	file, err := os.OpenFile("C:\\Users\\Administrator\\Desktop\\杂七杂八的文档\\Monster.txt", os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0666)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	writer := bufio.NewWriter(file)
	nn, err := writer.Write(bytes)
	if err != nil {
		panic(err)
	}
	err = writer.Flush()
	if err != nil {
		panic(err)
	}
	return nn
}

func UnSerial(nn int) serial.Monster {
	readFile, _ := os.OpenFile("C:\\Users\\Administrator\\Desktop\\杂七杂八的文档\\Monster.txt", os.O_RDONLY, 0666)
	defer readFile.Close()
	reader := bufio.NewReader(readFile)
	data := make([]byte, nn) //这里data的切片要大于等于文件中的字节数，不然无法存储所有的数据
	//data := make([]byte, 100)     //如果给data的空间比文件所含的字节数多，进行反序列化时需要注意，要对此切片进行再切片，否则无法成功反序列化
	n, err := reader.Read(data)
	if err != nil {
		panic(err)
	} else {
		fmt.Println(n)
		fmt.Println("read data :", string(data[:nn]))
		var monster serial.Monster
		err := json.Unmarshal(data, &monster)
		if err != nil {
			panic(err)
		}
		//json.Unmarshal(data[:nn], &monster)  //data的空间比文件所含的字节数，需要切片到有数据的部分
		fmt.Println(monster)
		return monster
	}
	//err = errors.New("手欠报个错")
	//panic(err)
}

func main() {
	defer func() {
		if err := recover(); err != nil { //recover可以获取到现有的错误信息(只能抓去那些被panic调用的错误)
			fmt.Println("error message :", err)
		}
	}()
	n := Serial()
	UnSerial(n)
}
