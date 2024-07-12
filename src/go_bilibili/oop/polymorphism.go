package _oop

// 3.多态

import "fmt"

type Usb interface {
	Start()
	Stop()
}

type Phone struct{}

// Phone实现Usb接口的所有方法
func (p Phone) Start() {
	fmt.Println("phone start working...")
}
func (p Phone) Stop() {
	fmt.Println("phone stop working...")
}

type Camera struct{}

// Camera只实现Usb接口的Start方法
func (c Camera) Start() {
	fmt.Println("camera start working...")
}

func InterfaceAssert(value interface{}) {
	if _, ok := value.(Usb); ok { // 接口断言
		fmt.Printf("%T implemented the Usb interface...\n", value)
	} else {
		fmt.Printf("%T does not implemented the Usb interface...\n", value)
	}
}

func main() {
	phone := Phone{}
	camera := Camera{}

	InterfaceAssert(phone)  // main.Phone implemented the Usb interface...
	InterfaceAssert(camera) // main.Camera does not implemented the Usb interface...
}
