package main

import "fmt"

type Currency int

const (
	USD Currency = iota // 美元
	EUR                 // 欧元
	GBP                 // 英镑
	RMB                 // 人民币
)

func main() {

	var a [3]int             // array of 3 integers
	fmt.Println(a[0])        // print the first element
	fmt.Println(a[len(a)-1]) // print the last element, a[2]

	// Print the indices and elements.
	for i, v := range a {
		fmt.Printf("%d %d\n", i, v)
	}

	// Print the elements only.
	for _, v := range a {
		fmt.Printf("%d\n", v)
	}

	var q [3]int = [3]int{1, 2, 3}
	var r [3]int = [3]int{1, 2}
	fmt.Println(q[2]) // "0"
	fmt.Println(r[2]) // "0"

	p := [...]int{1, 2, 3}
	fmt.Printf("%T\n", p) // "[3]int"

	b := [3]int{1, 2, 3}
	fmt.Println(b)
	//b = [4]int{1, 2, 3, 4} // compile error: cannot assign [4]int to [3]int

	symbol := [...]string{USD: "$", EUR: "€", GBP: "￡", RMB: "￥"}
	fmt.Println(RMB, symbol[RMB]) // "3 ￥"

}

/**

四种类型——数组、slice、map和结构体
*/
