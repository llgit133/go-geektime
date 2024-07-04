package _generic

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

func PersonCase(person Person) {
	person.Go()
	person.To()
	fmt.Println("===============================")
}

// PersonCase1  用interface{} 接口 替代泛型
func PersonCase1(person interface{}) {
	//断言
	p1, ok := person.(Person)
	if ok {
		p1.Go()
	} else {
		fmt.Println(" 类型不能识别 not Person")
	}
}
