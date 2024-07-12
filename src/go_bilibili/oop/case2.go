package _oop

import (
	"fmt"
)

type Worker struct {
	Name   string
	Age    uint8
	Gender Gender
}

// Run 成员方法
func (*Worker) Run() {
	fmt.Println("Worker run ")
}

func (*Worker) Sleep() {
	fmt.Println("Worker sleep ")
}

func StudentCase() {

	worker1 := Worker{
		Name:   "worker1",
		Age:    18,
		Gender: MALE,
	}
	worker1.Run()
	worker1.Sleep()
	fmt.Println("worker1", worker1)
	fmt.Println("======================================")

	worker2 := Worker{}
	worker2.Run()
	worker2.Sleep()
	fmt.Println("worker2", worker2)
	fmt.Println("======================================")

}

/**

func (Student) Run()
func (*Student) Run()
值类型接收器方法：用于不需要修改对象的场景，方法内使用的是接收器的副本。
指针类型接收器方法：用于需要修改对象的场景，方法内使用的是接收器的地址。

*/
