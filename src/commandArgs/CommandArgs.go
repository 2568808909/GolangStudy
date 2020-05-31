package main

import (
	"flag"
	"fmt"
	"os"
)

//这种方式虽然可以从命令行获取参数，但是参数的顺序一开始就要确定好，是不能改变的，这就使得程序不够灵活
func OsArgs() {
	//通过os.Args可以获取命令行参数 如：args.exe root mysql 8080获得的结果将会是：
	//args number = 4
	//arg[0] = args.exe
	//arg[1] = root
	//arg[2] = mysql
	//arg[3] = 8080
	args := os.Args
	fmt.Println("args number =", len(args))
	for index, arg := range args {
		fmt.Printf("arg[%d] = %s\n", index, arg)
	}
}

func flagArgs() {
	//通过flag包进行命令行参数的解析，实现类似 mysql -u root -p 123456的命令标志命令参数解析
	var user string
	var password string
	var host string
	var port int

	//StringVar用指定的名称、默认值、使用信息注册一个string类型flag，并将flag的值保存到p指向的变量
	flag.StringVar(&user, "u", "", "用户名")
	flag.StringVar(&password, "p", "", "密码")
	flag.StringVar(&host, "h", "localhost", "主机名")
	flag.IntVar(&port, "P", 3306, "端口")

	//从os.Args[1:](os.Args[0]是可执行文件名)中解析参数，不执行参数无法解析
	flag.Parse()

	//flagArgs.exe -u root -P 8080 -h 192.168.88.102 -p 654321   //这里的标志不能和参数贴在一起(-p123456)，不然会宝座
	//username = root ,password = 654321 ,host = 192.168.88.102 ,port = 8080
	fmt.Printf("username = %s ,password = %s ,host = %s ,port = %d", user, password, host, port)
}

func main() {

}
