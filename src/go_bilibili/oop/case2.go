package _oop

import (
	"fmt"
)

type Student struct {
	Name   string
	Age    uint8
	Gender Gender
}

// Run 成员方法
func (*Student) Run() {
	fmt.Println("Student run ")
}

func (*Student) Sleep() {
	fmt.Println("Student sleep ")
}

func StudentCase() {

	student1 := Student{
		Name:   "student1",
		Age:    18,
		Gender: MALE,
	}
	student1.Run()
	student1.Sleep()
	fmt.Println("student1", student1)
	fmt.Println("======================================")

	student2 := Student{}
	student2.Run()
	student2.Sleep()
	fmt.Println("student2", student2)
	fmt.Println("======================================")

}

/**

func (Student) Run()
func (*Student) Run()
值类型接收器方法：用于不需要修改对象的场景，方法内使用的是接收器的副本。
指针类型接收器方法：用于需要修改对象的场景，方法内使用的是接收器的地址。

*/
