package _oop

import "fmt"

// 枚举类型
const (
	FEMALE  Gender = 1
	MALE    Gender = 2
	THIRD   Gender = 3
	UNKNOWN Gender = 4
)

// Gender 起别名
type Gender int8

// User User类 成员变量
type User struct {
	Name   string
	Age    uint8
	Gender Gender
}

// Run 成员方法
// 需要 USer的具体实例，才能调用Run方法
func (User) Run() {
	fmt.Println("user run ")
}

func (User) Sleep() {
	fmt.Println("user sleep ")
}

// UserCase 对象实例化
func UserCase() {

	user1 := User{
		Name:   "zhangsan",
		Age:    10,
		Gender: MALE,
	}
	user1.Run()
	user1.Sleep()
	fmt.Println("user1:", user1)
	fmt.Println("======================================")

	user2 := new(User)
	user2.Name = "lisi"
	user2.Age = 20
	user2.Gender = FEMALE
	user1.Run()
	user1.Sleep()
	fmt.Println("user2", user2)
	fmt.Println("======================================")

	user3 := User{}
	user3.Run()
	user3.Sleep()
	user3.Name = "wangwu"
	user3.Age = 30
	user3.Gender = THIRD
	fmt.Println("user3", user3)
	fmt.Println("======================================")

	user4 := &User{}
	user4.Run()
	user4.Sleep()
	user4.Name = "wangwu"
	user4.Age = 40
	user4.Gender = THIRD
	fmt.Println("user3", user3)
	fmt.Println("======================================")

}

/**
值类型 (user1)：直接使用结构体字面值初始化，适用于需要值语义的场景。
指针类型 (user2 和 user3)：使用 new 或 & 创建指针，适用于需要引用语义或共享对象的场景。
访问方式：虽然 user2 和 user3 是指针类型，但 Go 语言自动处理指针解引用，所以访问字段和方法的方式与值类型一致。




自动解引用：
user := User{}
user.Run()  // OK
user.Sleep() // OK

userPtr := &User{}
userPtr.Run()  // OK, Go 自动解引用
userPtr.Sleep() // OK, Go 自动解引用


结构体属性
结构体方法

*/
