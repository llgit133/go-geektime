package main

import "time"

type Employee struct {
	ID        int
	Name      string
	Address   string
	DoB       time.Time
	Position  string
	Salary    int
	ManagerID int
}

var dilbert Employee

func main() {

}

/*

4.4 struct 结构体


4.4.1. 结构体字面值


4.4.2. 结构体比较


4.4.3. 结构体嵌入和匿名成员
****/
