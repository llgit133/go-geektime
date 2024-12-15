package _oop

import "fmt"

type MyType struct{}

func (t MyType) ValueReceiverMethod() {
	fmt.Println("Value receiver method called")
}

func (t *MyType) PointerReceiverMethod() {
	fmt.Println("Pointer receiver method called")
}

func main() {
	value := MyType{}
	pointer := &MyType{}

	// 值类型方法集调用
	value.ValueReceiverMethod()   // ✅
	pointer.ValueReceiverMethod() // ✅ 自动解引用调用值方法

	// 指针类型方法集调用
	// value.PointerReceiverMethod() // ❌ 编译错误，值不能调用指针方法
	pointer.PointerReceiverMethod() // ✅
}

// 方法集问题
