package main

import (
	"StructDemo/utils"
	"fmt"
)

type Person struct {
	Name   string
	Age    int
	Number *int
}

type A struct {
	Num int
}

type B struct {
	Num int
}

func test01(person Person) {
	person.Name = "Jack"
	fmt.Println("test ", person, *person.Number)
}

func test02(person *Person) {
	//在此处person这个变量实际上是一个Person类型的指针，要使用其调用person中的变量按照正常方式应该想下面一样通过(*person)获取到Person对象在调用
	//(*person).Name="Fuck"
	//但是Go中对此进行了优化，可以直接使用person这个指针来进行对象的各种调用，使得调用更加方便
	person.Name = "John"
	fmt.Println("test ", *person, *person.Number)
}

func main() {
	p1 := Person{"Tom", 22, new(int)}
	//结构体属于值类型，所以这里对p1中的数据进行拷贝，放入p2中，任何对p2的修改都不会影响p1,除非修改的是引用类型
	p2 := p1
	//Number是一个指针，所以这里对p2的修改也会影响到p1
	*p2.Number = 60
	fmt.Printf("address of p1 : %p p2 : %p", &p1.Name, &p2.Name)
	fmt.Println(p1.Age)
	p2.Name = "Jack"
	fmt.Println("p1`s name =", p1.Name, "p2`s name =", p2.Name)
	fmt.Printf("name : %p age : %p number : %p\n", &p1.Name, &p1.Age, &p1.Number)
	fmt.Printf("name : %s age : %d number : %d\n", p1.Name, p1.Age, p1.Number)

	//使用其他包的结构体也会遵循开头大写其他包能访问，首字母小写只能本包访问的原则
	student := utils.Student{Age: 10}
	fmt.Println(student.Age)

	//值传递，test01会对p1进行值拷贝，在test01中任何的修改都不会影响p1
	test01(p1)
	fmt.Println("main p1:", p1, *p1.Number)
	fmt.Println()

	//引用传递，给test02的参数传递的是p1的地址，在test02中的修改也会影响到p1
	test02(&p1)
	fmt.Println("main p1:", p1, *p1.Number)
	fmt.Println()

	//不同的类型的结构体实例可以互相转换，不过有个前提就是：两者之间的字段数量，字段名和字段类型都要一致
	a := A{5}
	b := B(a)
	fmt.Println(b)
}
