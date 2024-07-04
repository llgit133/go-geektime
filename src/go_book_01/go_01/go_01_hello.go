package main

import "fmt"

func main() {

	//输出
	fmt.Println("Hello, 世界")
	fmt.Printf("%v\n", "你好")          // 可以作为任何值的占位符输出
	fmt.Printf("%v %T\n", "枫枫", "枫枫") // 打印类型
	fmt.Printf("%d\n", 3)             // 整数
	fmt.Printf("%.2f\n", 1.25)        // 小数
	fmt.Printf("%s\n", "哈哈哈")         // 字符串
	fmt.Printf("%#v\n", "")           // 用go的语法格式输出，很适合打印空字符串

	// 输入
	fmt.Println("输入您的名字：")
	var name string
	fmt.Scan(&name) // 这里记住，要在变量的前面加个&, 后面讲指针会提到
	fmt.Println("你输入的名字是", name)
}

/***


PS D:\Code\go\geektime-golang\Golang_Puzzlers\src\go_book_01\go_01> go run .\go_01_hello.go
Hello, 世界


*/
