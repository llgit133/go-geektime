package main

import (
	"fmt"
	_exception "go_book_01/exception"
	_generic "go_book_01/generic"
	_interface "go_book_01/interface"
	_oop "go_book_01/oop"
)

func main() {
	_oop.UserCase()
	_oop.StudentCase()

	//值传递和指针传递
	_oop.AbCase()

	_interface.StudentCase()
	_interface.TeacherCase()

	// 多态测试
	fmt.Println("多态测试")
	_interface.PersonCase(&_interface.Student{})
	_interface.PersonCase(&_interface.Teacher{})

	// 方法集问题 方法集：决定了一个类型可以调用哪些方法。
	fmt.Println("方法集问题")
	student := _interface.Student{}
	student.To()
	_interface.PersonCase(&student)

	teacher := &_interface.Teacher{}
	teacher.To()
	_interface.PersonCase(teacher)

	//泛型测试
	fmt.Println("泛型测试")
	_generic.PersonCase1(&student)
	_generic.PersonCase1(teacher)

	// 异常测试
	fmt.Println("异常测试")
	_exception.AACase()
	_exception.TryCatchCase()

	// 参数预计算
	_exception.DeferCase()
}
