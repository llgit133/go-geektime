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

// Run 非成员方法
func Run() {
	fmt.Println(" run ")
}

func (User) Sleep() {
	fmt.Println("user sleep ")
}

// UserCase 对象实例化
func UserCase() {

	fmt.Println("================1. 直接初始化字段================")
	user0 := User{"1111", 9, FEMALE}
	fmt.Println("user0:", user0)

	fmt.Println("===============值类型的实例=======================")
	user3 := User{}
	user3.Name = "wangwu"
	user3.Age = 30
	user3.Gender = THIRD
	fmt.Println("user3", user3)

	Run()

	fmt.Println("================2. 基于字段名的初始化======================")
	user1 := User{
		Name:   "2222",
		Age:    10,
		Gender: MALE,
	}
	fmt.Println("user1:", user1)

	fmt.Println("================3. 使用 new 函数======================")

	user2 := new(User)
	user2.Name = "3333"
	user2.Age = 20
	user2.Gender = FEMALE
	fmt.Println("user2", user2)

	// 以下为常用的方式
	fmt.Println("==================指针类型初始化1====================")
	user5 := &User{"5555", 40, THIRD}
	fmt.Println("user5", user5)

	fmt.Println("==================指针类型2====================")
	user4 := &User{}
	user4.Name = "wangwu"
	user4.Age = 40
	user4.Gender = THIRD
	fmt.Println("user3", user3)

}

/**

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
