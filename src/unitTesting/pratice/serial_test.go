package main

import (
	"testing"
)

func TestSerial(t *testing.T) {
	count := Serial()
	t.Log("序列化成功，写入文件字节数:", count)
}

func TestUnSerial(t *testing.T) {
	defer func() {
		if err := recover(); err != nil { //recover可以获取到现有的错误信息
			t.Fatalf("执行出错，错误信息：%v", err)
		}
	}()
	monster := UnSerial(60)
	t.Log("反序列化成功 :", monster)
}
