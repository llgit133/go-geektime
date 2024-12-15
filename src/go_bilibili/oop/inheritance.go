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

//func main() {
//	var stu = Student{Person{"Alice", 12}, 100}
//	stu.PrintInfo() // name = Alice, age = 12
//	stu.Study()     // student 100 is studying...
//
//	var tch = Teacher{&Person{"Bob", 22}, 200}
//	tch.PrintInfo() // name = Bob, age = 22
//	tch.Teach()     // teacher 200 is teaching...
//}

/*

使用值类型嵌套：Person 是直接嵌套在 Student 中（值类型），所有字段和方法都会被拷贝到 Student。
使用指针类型嵌套：Person 是以指针的形式嵌套在 Teacher 中，字段和方法不会被拷贝，而是通过指针引用。


综合建议
通用场景：优先使用值类型嵌套（如 Student），因为它简单且不易引发数据一致性问题。
高性能和共享需求场景：在需要共享基类数据或动态分配内存时，使用指针类型嵌套（如 Teacher）。

最佳实践：
如果基类较大（占用内存多），建议使用指针类型嵌套。
如果继承层级较深或涉及并发操作，尽量避免值类型嵌套，以减少数据拷贝带来的复杂性。

*/
