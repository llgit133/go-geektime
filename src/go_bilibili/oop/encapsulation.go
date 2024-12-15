package _oop

import "fmt"

// 1.oop 封装

type Pet struct {
	name   string
	age    int
	gender string
}

// GetName 访问name字段
func (pet Pet) GetName() string {
	return pet.name
}
func (pet *Pet) SetName(name string) {
	pet.name = name
}

// GetAge 访问age字段
func (pet Pet) GetAge() int {
	return pet.age
}
func (pet *Pet) SetAge(age int) {
	pet.age = age
}

// GetGender 访问gender字段
func (pet Pet) GetGender() string {
	return pet.gender
}
func (pet *Pet) SetGender(gender string) {
	pet.gender = gender
}

// Print Pet 的其他方法
func (pet Pet) Print() {
	fmt.Printf("Pet info: <name: %s, age: %d, gender: %s>\n", pet.name, pet.age, pet.gender)
}
