package model

type teacher struct {
	name    string
	balance float64
}

//TODO 直接返回teacher好像也可以，和返回指针有什么不同呢，为什么直接返回teacher实例会有警告
func GetInstance(name string, balance float64) *teacher {
	return &teacher{
		name:    name,
		balance: balance,
	}
}

func (teacher *teacher) GetName() string {
	return teacher.name
}

func (teacher *teacher) GetBalance() float64 {
	return teacher.balance
}
