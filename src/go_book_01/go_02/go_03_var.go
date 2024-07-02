package main

import (
	"fmt"
)

func main() {
	varTest()
	pointTest()
}

func varTest() {
	var a string
	fmt.Println(a) // ""

	var i, j, k int                 // int, int, int
	var b, c, d = true, 2.3, "four" // bool, float64, string
	println(i, j, k, b, c, d)

	//var f, err = os.Open(name) // os.Open returns a file and an error

	q := 100                  // an int
	var boiling float64 = 100 // a float64
	var names []string
	var err error
	//var p Point
	println(q, boiling, names, err)
}

func pointTest() {
	x := 1
	p := &x         // p, of type *int, points to x
	fmt.Println(*p) // "1"
	*p = 2          // equivalent to x = 2
	fmt.Println(x)  // "2"

}

/**
2.3. 变量
var 变量名字 类型 = 表达式


2.3.1. 简短变量声明

i, j = j, i // 交换 i 和 j 的值


2.3.2. 指针

*/
