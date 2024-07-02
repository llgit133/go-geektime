package main

import (
	"flag"
	"fmt"
)

var name string

func init() {
	flag.StringVar(&name, "name", "everyone", "The greeting object.")
	//flag.String("name", "everyone", "The greeting object.")
}

func main() {
	flag.Parse()
	fmt.Printf("Hello, %s!\n", name)
}

/**
flag 同于接收和解析命令参数，可以利用这个包解析命令行参数，执行程序
StringVar 四个参数：参数地址，参数名，默认值，注释
String    三个参数：参数名，默认值，注释


main 包中的所有 init 函数都会先于 main 函数执行，
main 包直接依赖的那些代码包中的 init 函数会先于 main 包中的 init 函数执行，以此类推（与依赖方向完全相反）。


*/
