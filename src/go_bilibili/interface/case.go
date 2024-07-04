package _interface

import "fmt"

// Student  结构体
type Student struct {
	Name   string
	Age    uint8
	Gender uint8
}

type Teacher Student

// Person 接口定义
type Person interface {
	Go()
	To()
}

// Go 接口的实现 ,只要方法对应的接口方法，就可以实现接口
// 多态，Student 和 Teacher,都实现了Person中Go和To方法

func (stu *Student) Go() {
	fmt.Println("Student Go")
}

func (stu *Student) To() {
	fmt.Println("Student To")
}

func (tea *Teacher) Go() {
	fmt.Println("Teacher Go")
}

func (tea *Teacher) To() {
	fmt.Println("Teacher To")
}

func StudentCase() {
	stu := &Student{}
	stu.Go()
	stu.To()
	fmt.Println("===============================")
}

func TeacherCase() {
	tea := &Teacher{}
	tea.Go()
	tea.To()
	fmt.Println("===============================")
}

func PersonCase(person Person) {
	person.Go()
	person.To()
	fmt.Println("===============================")
}
