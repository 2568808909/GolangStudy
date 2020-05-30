package model

import "fmt"

type person struct {
	Name string
	age  int
}

type teacher struct {
	person
}

type B struct {
}

type C struct {
}

//Go中支持多重继承，只需要写入多个匿名结构体即可，不过剁成继承容易造成混乱，所以不推荐使用
type D struct {
	B
	C
}

type A struct {
	person
	int //基本类型也可以进行继承
}

func GetTeacher(name string, age int) *teacher {
	return &teacher{
		person{
			Name: name,
			age:  age,
		},
	}
}

func ShowTeacher(teacher teacher) {
	//对于“父类”中首字母小写的字段和方法，“子类”都可以访问，这一点和私有也不同，不过这仅限于在包中可以
	fmt.Println(teacher.age)
	fmt.Println(teacher.Name)
}
