package main

import (
	"fmt"
	_exception "go_bilibili/exception"
	_generic "go_bilibili/generic"
	_interface "go_bilibili/interface"
	_oop "go_bilibili/oop"
)

func main() {
	//------------------------对象实例化------------------------
	_oop.UserCase()
	_oop.StudentCase()

	//------------------------值传递和指针传递------------------------
	_oop.AbCase()

	_interface.StudentCase()
	_interface.TeacherCase()

	// ------------------------多态测试------------------------
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

	//------------------------泛型测试------------------------
	fmt.Println("泛型测试")
	_generic.PersonCase1(&student)
	_generic.PersonCase1(teacher)

	// ------------------------异常测试------------------------
	fmt.Println("异常测试")
	_exception.AACase()
	_exception.TryCatchCase()

	// 参数预计算
	_exception.DeferCase()

	// ------------------------oop -1 封装------------------------
	fmt.Println("-----oop -1 封装----")
	pet := &_oop.Pet{}
	pet.SetName("111")
	pet.SetAge(2)
	pet.SetGender("3")

	println(pet.GetName())
	println(pet.GetAge())
	println(pet.GetGender())

	pet.Print()

	//------------------------ oop -2 继承------------------------
	fmt.Println("-----oop -2 继承----")
	var stu = _oop.Student{_oop.Person{"Alice", 12}, 100}
	stu.PrintInfo() // name = Alice, age = 12
	stu.Study()     // student 100 is studying...

	var tch = _oop.Teacher{&_oop.Person{"Bob", 22}, 200}
	tch.PrintInfo() // name = Bob, age = 22
	tch.Teach()     // teacher 200 is teaching...

	//有无方法重写

	// ------------------------oop -3 多态------------------------
	fmt.Println("-----oop -3 多态----")

}
