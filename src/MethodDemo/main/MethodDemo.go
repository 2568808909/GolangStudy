package main

import (
	"fmt"
	"strconv"
)

type Person struct {
	Name string
	Age  int
}

type Point struct {
	X int
	Y int
}

//方法的定义方式
//func (person Person) test() int {
//	return 0
//}

//给Person绑定一个方法，但是在Go中，因为结构体属于值类型数据，所以调用时使用的值传递，在此方法中对Name字段的修改，不会影响到main函数中的person
func (person Person) test01() {
	person.Name = "John"
	fmt.Println("test01 ", person)
}

//除了在方法名前定义要绑定的对象类型外，其他和函数的定义是一致的
func (point Point) calculation(add int) (sum int, sub int) {
	sum = point.X + point.Y + add
	sub = point.X - point.Y - add
	return
}

//通过为某个类型绑定String方法可以规定这个类型的输出格式，但是要注意的是，这里要注意对象传参时用的是指针还是实例，指针的话&person才会格式化输出
func (person Person) String() string {
	return "Name=" + person.Name + " Age=" + strconv.Itoa(person.Age)
}

func (person *Person) test02() {
	person.Name = "John"
	fmt.Println("test02 ", *person)
}

type myInt int

//要影响到原来的值同样要使用指针
func (i *myInt) add(num int) {
	*i += myInt(num)
}

func main() {
	person := Person{"Tom", 10}
	person.test01()
	fmt.Println("main", person)
	point := Point{5, 6}
	sum, sub := point.calculation(4)
	fmt.Println("sum =", sum, "sub =", sub)

	//以引用传递的方式调用方法
	//(&person).test02()
	//不过Go中对调用方式有所优化，直接调用也可以
	person.test02()
	//其实test01也可以用引用传递的方式调用，但是决定方法中的修改是否会影响到此处的person还是方法定义时，还是使用的是否为指针
	person.test01()
	fmt.Println("main", person)

	//同样对于int这样的基本类型其实也是可以为其绑定方法的，不过，不能直接为int绑定，而是要用type为其定期另一个类型才可
	var i myInt = 10
	i.add(5)
	fmt.Println(i)
}
