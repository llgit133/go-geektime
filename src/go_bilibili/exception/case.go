package _exception

import "fmt"

type AA struct {
	Name string
}

type BB struct {
	Name string
}

func (u *AA) PrintName() {
	fmt.Println("User Name:", u.Name)
}

func (u *AA) SetName(name string) {
	u.Name = name
}

func (s *BB) PrintName() {
	fmt.Println("BB Name:", s.Name)
}

func (s *BB) SetName(name string) {
	s.Name = name
}

func AACase() {

	defer func() {
		fmt.Println("AA defer")
	}()
	return
}

func TryCatchCase() {

	// 捕获异常,程序不中断，捕获异常信息
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println("err:", err)
		}
	}()
	PanicCase()
}

func PanicCase() {

	panic("程序出现异常")
}

func DeferCase() {

	fmt.Println("参数预计算")

	// 1.传参
	i := 1
	defer func(j int) {
		fmt.Println("defer j", j)
	}(i + 1)

	//2.闭包
	defer func() {
		i++
		fmt.Println("defer i", i)
	}()
	i = 99
	return
	
	// 栈的顺序
	//defer i 100
	//defer j 2
}
