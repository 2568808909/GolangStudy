package main

import "fmt"

type Usb interface {
	start()
	stop()
}

type Phone struct {
	Name string
}

type Camera struct {
	Name string
}

type Computer struct {
}

func (phone Phone) start() {
	fmt.Println(phone.Name, "正在开始传输文件.....")
}

func (phone Phone) stop() {
	fmt.Println(phone.Name, "数据传输完成.....")
}

func (phone Phone) call() {
	fmt.Println(phone.Name, "正在打电话.....")
}

func (camera Camera) start() {
	fmt.Println(camera.Name, "正在开始传输照片.....")
}

func (camera Camera) stop() {
	fmt.Println(camera.Name, "数据传输完成.....")
}

func (computer Computer) transform(usb Usb) {
	usb.start()
	//通过类型断言，判断该对象是否为Phone，如果是则将其转回phone对象，并调用器call方法
	if phone, ok := usb.(Phone); ok {
		phone.call()
		//由于结构体为值类型，所以转回该结构体对象会发生只拷贝，这里的修改并不会影响原来的对象
		//如果传入的是指针，则能为同一对象，这里的修改会修改原来的对象
		phone.Name = "三星"
		fmt.Printf("address %p ,name : %s\n", &phone, phone.Name)
	}
	usb.stop()
}

//通过接口可以实现多态，因为实现了接口的对象可以向上转型为接口对象
func main() {
	var usb []Usb
	usb = append(usb, Camera{"索尼"})
	usb = append(usb, Phone{"小米"})
	usb = append(usb, Phone{"华为"})

	computer := Computer{}

	for _, v := range usb {
		computer.transform(v)
		fmt.Println(v)
	}
}
