package _oop

import (
	"fmt"
)

//	2.oop 继承

// Person 基类
type Person struct {
	Name string
	Age  int
}

func (per Person) PrintInfo() {
	fmt.Printf("name = %s, age = %d\n", per.Name, per.Age)
}

// Student 派生类
type Student struct {
	Person    // 嵌套匿名结构体
	StudentId int
}

func (stu Student) Study() {
	fmt.Printf("student %d is studying...\n", stu.StudentId)
}

// Teacher 派生类
type Teacher struct {
	*Person   // 嵌套匿名结构体指针
	TeacherId int
}

func (tch Teacher) Teach() {
	fmt.Printf("teacher %d is teaching...\n", tch.TeacherId)
}

func main() {
	var stu = Student{Person{"Alice", 12}, 100}
	stu.PrintInfo() // name = Alice, age = 12
	stu.Study()     // student 100 is studying...

	var tch = Teacher{&Person{"Bob", 22}, 200}
	tch.PrintInfo() // name = Bob, age = 22
	tch.Teach()     // teacher 200 is teaching...
}
