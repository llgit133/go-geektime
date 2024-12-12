// Package tempconv performs Celsius and Fahrenheit temperature computations.
package main

import "fmt"

type Celsius float64    // 摄氏温度
type Fahrenheit float64 // 华氏温度

const (
	AbsoluteZeroC Celsius = -273.15 // 绝对零度
	FreezingC     Celsius = 0       // 结冰点温度
	BoilingC      Celsius = 100     // 沸水温度
)

// 输入celsius转换为fahrenheit   float64
func CToF(c Celsius) Fahrenheit { return Fahrenheit(c*9/5 + 32) }

func FToC(f Fahrenheit) Celsius { return Celsius((f - 32) * 5 / 9) }

// 该方法返回该类型对象c带着°C温度单位的字符串
// 许多类型都会定义一个String方法，因为当使用fmt包的打印方法时，将会优先使用该类型对应的String方法返回的结果打印--类似与方法重写
func (c Celsius) String() string { return fmt.Sprintf("%g°C", c) }

func (c Celsius) printString() string { return fmt.Sprintf("%g°C", c) }

func main() {

	fmt.Println(CToF(AbsoluteZeroC))
	fmt.Println(CToF(FreezingC))
	fmt.Println(CToF(BoilingC))

	fmt.Printf("%g\n", BoilingC-FreezingC) // "100" °C
	boilingF := CToF(BoilingC)
	fmt.Printf("%g\n", boilingF-CToF(FreezingC)) // "180" °F
	//fmt.Printf("%g\n", boilingF-FreezingC)       // compile error: type mismatch

	var c Celsius
	var f Fahrenheit
	fmt.Println(c == 0) // "true"
	fmt.Println(f >= 0) // "true"
	//fmt.Println(c == f)          // compile error: type mismatch
	fmt.Println(c == Celsius(f)) // "true"!  但是Celsius(f)是类型转换操作，它并不会改变值，仅仅是改变值的类型而已。测试为真的原因是因为c和f都是零值

	println("==============================================")
	d := FToC(212.0) //100
	fmt.Println(d)
	fmt.Println(d.String())      // "100°C"
	fmt.Println(d.printString()) // "100"
	fmt.Printf("%v\n", d)        // "100°C"; no need to call String explicitly
	fmt.Printf("%s\n", d)        // "100°C"
	fmt.Println(d)               // "100°C"
	fmt.Printf("%g\n", d)        // "100"; does not call String
	fmt.Println(float64(d))      // "100"; does not call String

}

/**
Celsius和Fahrenheit分别对应不同的温度单位。
它们虽然有着相同的底层类型float64，但是它们是不同的数据类型，因此它们不可以被相互比较或混在一个表达式运算







*/
