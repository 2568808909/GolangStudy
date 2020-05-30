package main

import (
	"extendDemo/model"
	"fmt"
)

type Student struct {
	StuNo string
	Name  string
	Age   int
	//Pupil  //可以进行套娃，所以Go中的继承并没有严格的父子概念
}

//Go中继承只要如下面一般，写入一个匿名结构体即可
type Pupil struct {
	Student
}

type Graduate struct {
	Student
}

//这些"父类"的方法，“子类”都可以调用（之所以用引号，是Go中并没有类的概念，继承也没有严格的父子概念）
func (stu *Student) GetAge() int {
	return stu.Age
}

func (stu *Student) GetStuNo() string {
	return stu.StuNo
}

//这是"子类"特有的方法
func (pupil *Pupil) study() {
	fmt.Println("小学生正在学习.......")
}

func (graduate *Graduate) study() {
	fmt.Println("大学生正在学习.......")
}

func main() {
	pupil := Pupil{Student{"11111", "阚忠良", 5}}
	//调用所继承的结构体的属性，通过以下两种方式调用
	fmt.Println("graduate student number =", pupil.Student.Age)
	fmt.Println("graduate student number =", pupil.StuNo)
	fmt.Println(pupil)

	graduate := Graduate{Student{"20164698", "PKB", 21}}
	fmt.Println("graduate student number =", graduate.StuNo)
	fmt.Println(graduate)

	teacher := model.GetTeacher("杨丰", 60)
	fmt.Println(teacher)
}
