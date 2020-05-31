package main

import (
	"encoding/json"
	"fmt"
)

type Monster struct {
	Name   string  `json:"name"`
	Age    int     `json:"age"`
	Salary float64 `json:"salary"`
}

func SerialAndUnSerial() {
	monster := Monster{"牛魔王", 1000, 5212.65}
	//序列化
	bytes, e := json.Marshal(monster)
	if e == nil {
		fmt.Printf("json str = %s\n", bytes)
	} else {
		fmt.Println(e)
	}
	var another = Monster{}
	//反序列化
	e = json.Unmarshal(bytes, &another)
	fmt.Println(another)

	//反序列化成map(注意这样不需要使用make给map分配空间，因为在Unmarshal已经封装了分配空间的操作)
	var monsterMap map[string]interface{}
	e = json.Unmarshal(bytes, &monsterMap)
	fmt.Println(monsterMap)
}

//切片的序列化及反序列化，方法都是类似的
func SerialSlice() {
	var monsters []Monster
	monsters = append(monsters, Monster{"牛魔王", 1000, 5212.65})
	monsters = append(monsters, Monster{"黑熊精", 654, 5258.65})
	bytes, e := json.Marshal(monsters)
	if e == nil {
		fmt.Printf("json str = %s\n", bytes)
	} else {
		fmt.Println(e)
	}
	var slice []Monster
	//反序列化
	e = json.Unmarshal(bytes, &slice)
	fmt.Println(slice)
}

func main() {
	SerialAndUnSerial()
}
