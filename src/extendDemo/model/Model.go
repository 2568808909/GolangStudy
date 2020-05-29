package model

import "fmt"

type person struct {
	Name string
	age  int
}

type teacher struct {
	person
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
