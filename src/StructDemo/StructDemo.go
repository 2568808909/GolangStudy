package main

import (
	"fmt"
	"unsafe"
)

type Person struct {
	Name   string
	Age    *int
	Number *int
}

func main() {
	p1 := Person{}
	p1.Age = new(int)
	p1.Number = new(int)
	p2 := p1
	fmt.Printf("address of p1 : %p p2 : %p", &p1.Name, &p2.Name)
	fmt.Println(p1.Age)
	p2.Name = "Jack"
	fmt.Println("p1`s name =", p1.Name, "p2`s name =", p2.Name)
	fmt.Printf("name : %p age : %p number : %p\n", &p1.Name, &p1.Age, &p1.Number)
	fmt.Printf("name : %s age : %d number : %d\n", p1.Name, p1.Age, p1.Number)
	p1.Name = "sn5Diphone6sn5Diphone6sn5Diphone6sn5Diphone6sn5Diphone6sn5Diphone6sn5Diphone6sn5Diphone6"

	fmt.Println(unsafe.Sizeof(p1.Name))
	fmt.Printf("name : %p age : %p number : %p\n", &p1.Name, &p1.Age, &p1.Number)
	fmt.Printf("name : %s age : %d number : %d\n", p1.Name, p1.Age, p1.Number)

}
