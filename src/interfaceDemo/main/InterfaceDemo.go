package main

import (
	"fmt"
	"math/rand"
	"sort"
	"strconv"
)

//这个接口包含三个方法，实现这个接口的对象要包含着三个方法(少一个也不行)即可，而不需要像Java一样显示的写出implements xxx
//另外一个自定义类型可以实现多个接口，一个接口也可以继承多个接口，如果某个类型要实现这个接口，要把这个接口所继承的接口的方法都实现
//type Interface interface {
//	// Len方法返回集合中的元素个数
//	Len() int
//	// Less方法报告索引i的元素是否比索引j的元素小
//	Less(i, j int) bool
//	// Swap方法交换索引i和j的两个元素
//	Swap(i, j int)
//}

type A interface {
	test01()
}

type B interface {
	//被同一个接口继承的接口不能有同样方法，不然实现时就不知道应该实现哪一个。
	//但是如果两个接口没有同时被一个接口继承，而是被同一个类型实现，有重复的方法时没有问题的，因为Go只关心有没有实现接口的素有方法
	//test01()
	test02()
}

type C interface {
	A
	B
	test03()
}

type Hero struct {
	Name string
	Age  int
}

func (hero Hero) test01() {
	fmt.Printf("%p %p\n", &hero, hero)
}

func (hero Hero) test02() {
	fmt.Printf("%p %p\n", &hero, hero)
}

func (hero Hero) test03() {

}

//这里想实现对Hero切片的排序，因为方法不定绑定在[]Hero上，所以创建一个自定义类型HeroSlice替代[]Hero，并为其绑定方法
type HeroSlice []Hero

func (hs HeroSlice) Len() int {
	return len(hs)
}

func (hs HeroSlice) Less(i, j int) bool {
	return hs[i].Age < hs[j].Age
}

func (hs HeroSlice) Swap(i, j int) {
	//temp := hs[i]
	//hs[i] = hs[j]
	//hs[j] = temp
	//以上三条指令相当于以下一条指令，Go中可以以这种形式进行数据的交换
	hs[i], hs[j] = hs[j], hs[i]
}

func main() {
	var hs HeroSlice
	for i := 0; i < 10; i++ {
		hero := Hero{
			Name: "英雄" + strconv.Itoa(rand.Intn(100)),
			Age:  rand.Intn(100),
		}
		hs = append(hs, hero)
	}
	for _, v := range hs {
		fmt.Println(v)
	}
	//func Sort(data Interface) 传入的参数要实现Interface这个接口
	sort.Sort(hs)
	fmt.Println()
	for _, v := range hs {
		fmt.Println(v)
	}

	//C接口继承自A和B
	hero := Hero{}
	fmt.Printf("%p\n", &hero)
	//接口时引用类型数据，将对象赋值给接口时，如果直接赋值，会拷贝一份数据，并指向这份新的数据，如果是将对象引用给接口，则指向这个对象的地址
	var c C = hero
	hero.Age = 10
	fmt.Println(c)
	fmt.Println(hero)
	var a A = &Hero{}
	var b B = &Hero{}
	fmt.Println(c)
	fmt.Println(a)
	fmt.Println(b)
}
