package _package

import "fmt"

// CToF converts a Celsius temperature to Fahrenheit.
func CToF(c Celsius) Fahrenheit { return Fahrenheit(c*9/5 + 32) }

// FToC converts a Fahrenheit temperature to Celsius.
func FToC(f Fahrenheit) Celsius { return Celsius((f - 32) * 5 / 9) }

func main() {
	fmt.Printf("Brrrr! %v\n", AbsoluteZeroC) // "Brrrr! -273.15°C"

}

/**
同package下
我们把变量的声明、对应的常量，还有方法都放到tempconv.go源文件中：
转换函数则放在另一个conv.go源文件中：


*/
