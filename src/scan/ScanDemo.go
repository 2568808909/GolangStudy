package main

import "fmt"

func main() {

	var age int
	var name string
	var balance float64
	var pass bool

	fmt.Scanln(&age) //若输入不符合数字格式的字符，会从头开始解析，取第一个不符合数字格式之前的进行数字解析(若没有，返回0)。这个字符之后的字符被放入缓冲区，作为下次输入
	fmt.Scanln(&name)
	fmt.Scanln(&balance)
	fmt.Scanln(&pass)

	fmt.Println(age)
	fmt.Println(name)
	fmt.Println(balance)
	fmt.Println(pass)

	fmt.Scanf("%d %s %f %v", &age, &name, &balance, &pass) //一旦前面某个变量的输入格式不正确，后面的变量的值都会用默认值代替，这个格式不正确的变量的取值可以参考上面

	fmt.Printf("%d %s %f %v", age, name, balance, pass)
}
