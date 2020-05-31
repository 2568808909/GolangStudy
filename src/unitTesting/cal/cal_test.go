package cal

import (
	"testing"
)

//相应路径下的中断输入go test -v后，Go就会启用其单元测试框架，这个框架会导入所有名为xxx_test.go的文件，并执行这些文件内的Testxxx(t *testing.T)方法
//go test不加-v这个option，则只会打印错误日志
//值执行某个xxx_test.go文件下的单元测试方法		go test -v xxx_test.go xxx.go (要执行的单元测试源文件和被测试的函数所在的源文件,顺序没有关系)
//只执行某个单元测试方法		go test -v -test.run TestXXX
//方法必须要以Test开头，接下来的字幕不能小写(如,Testadd)
func TestAdd(t *testing.T) {
	res := addUpdater(10)
	if res != 55 {
		t.Fatalf("计算结果错误，预计值55，实际值%d", res)
	}
	t.Log("计算结果正确....")
}

func TestAnother(t *testing.T) {
	res := addUpdater(100)
	if res != 5050 {
		t.Fatalf("计算结果错误，预计值5050，实际值%d", res)
	}
	t.Log("计算结果正确....")
}
