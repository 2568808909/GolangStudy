package main

import (
	"fmt"
	"strconv"
)

//基本数据转为String
func data2String() {
	var num int32 = 128
	var num2 float64 = 0.999
	var b bool = true
	var char byte = 'c'
	var str string

	//占位符相关可以查看文档 https://studygolang.com/pkgdoc   fmt处
	str = fmt.Sprintf("%d", num)
	fmt.Printf("str type %T str=%q\n", str, str)

	//1.基本数据类型转字符串方法1 通过fmt.Sprintf()
	str = fmt.Sprintf("%f", num2)
	fmt.Printf("str type %T str=%q\n", str, str)

	str = fmt.Sprintf("%v", b)
	fmt.Printf("str type %T str=%q\n", str, str)

	str = fmt.Sprintf("%c", char)
	fmt.Printf("str type %T str=%q\n", str, str)
	fmt.Println()
	/**结果
	str type string str="1"
	str type string str="0.999000"
	str type string str="true"
	str type string str="c"
	*/
	//2.通过strconv包实现基本数据转字符串
	str = strconv.FormatBool(b)
	fmt.Printf("str type %T str=%q\n", str, str)

	str = strconv.FormatInt(int64(num), 10) //第二个参数表示转成几进制在2-36之间，第一个参数要转成int64才可
	fmt.Printf("str type %T str=%q\n", str, str)
}

//String转为各种基本数据类型
func string2Data() {
	var str string
	str = "false"
	var b bool
	//这个函数是这样的 func ParseBool(str string) (value bool, err error) 也就是说有两个返回值，第二个为错误信息
	b, _ = strconv.ParseBool(str) //这里_表示忽略第二个返回值
	fmt.Printf("str type %T str=%v\n", b, b)

	str = "123"
	var i int64
	i, _ = strconv.ParseInt(str, 10, 0) //不管是转成整数还是浮点数，最终都是转成int64或float64，如果要使用更小范围的数据，既要显示转换
	fmt.Printf("str type %T str=%d\n", i, i)

	//对于转换不成功的数值返回该类型的默认值，int->0 bool->false
	str = "hello"
	var e error
	i, _ = strconv.ParseInt(str, 10, 0) //hello并不可以转换为int64类型，这里会默认返回一个0
	fmt.Printf("str type %T str=%d\n", i, i)
	fmt.Println(e)
}

func main() {
	string2Data()
}
