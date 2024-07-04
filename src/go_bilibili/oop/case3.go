package _oop

import "fmt"

type AA struct {
	Name string
}

type BB struct {
	Name string
}

func (u AA) PrintName() {
	fmt.Println("User Name:", u.Name)
}

func (u AA) SetName(name string) {
	u.Name = name
}

func (s *BB) PrintName() {
	fmt.Println("BB Name:", s.Name)
}

func (s *BB) SetName(name string) {
	s.Name = name
}

func AbCase() {
	aa := AA{Name: "Alice"}
	aa.PrintName() // 输出: User Name: Alice
	aa.SetName("Bob")
	aa.PrintName() // 输出: User Name: Alice, 修改无效

	bb := BB{Name: "Charlie"}
	bbPtr := &bb
	bbPtr.PrintName() // 输出: Student Name: Charlie
	bbPtr.SetName("Dave")
	bbPtr.PrintName() // 输出: Student Name: Dave, 修改有效
}

/**
1.方法集问题：

2.值传递和指针传递：
User Name: Alice
User Name: Alice
BB Name: Charlie
BB Name: Dave

*/
