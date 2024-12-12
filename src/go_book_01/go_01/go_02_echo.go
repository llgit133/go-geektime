// Echo1 prints its command-line arguments.
package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {

	// 1. for循环
	var s, sep string
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println("for循环" + s)

	// 2.range
	s1, sep1 := "", ""
	for _, arg := range os.Args[1:] {
		s1 += sep1 + arg
		sep1 = " "
	}
	fmt.Println("for range:" + s1)

	// += 改为 Join
	fmt.Println(strings.Join(os.Args[1:], " "))

	// 不格式化
	fmt.Println(os.Args[1:])

	// 即被执行命令本身的名字
	fmt.Println(os.Args[0])

	for i := 1; i < len(os.Args); i++ {
		fmt.Println("{}:{}", i, os.Args[i])
	}

}

/**
m

1.var s, sep string
var 声明定义了两个 string 类型的变量 s 和 sep。变量会在声明时直接初始化。如果变量没有显式初始化，则被隐式地赋予其类型的 零值（zero value），数值类型是 0，字符串类型是空字符串 ""。
这个例子里，声明把 s 和 sep 隐式地初始化成空字符串


等价的变量声明方式
s := ""
var s string
var s = ""
var s string = ""



2.for循环
Go 语言只有 for 循环这一种循环语句  无while do while。for 循环有多种形式，

for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}


for _, arg := range os.Args[1:] {
        s += sep + arg
        sep = " "
    }
_ ：忽略索引，只使用值。一种思路是把索引赋值给一个临时变量（如 temp）然后忽略它的值，但 Go 语言不允许使用无用的局部变量（local variables），因为这会导致编译错误。
Go 语言中这种情况的解决方法是用 空标识符（blank identifier），即 _

空标识符可用于在任何语法需要变量名但程序逻辑不需要的时候（如：在循环里）丢弃不需要的循环索引，并保留元素值
**/
