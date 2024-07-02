package main

func main() {

}

// 计算两个整数值的的最大公约数
func gcd(x, y int) int {
	for y != 0 {
		x, y = y, x%y
	}
	return x
}

// 计算斐波纳契数列（Fibonacci）的第N个数
func fib(n int) int {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		x, y = y, x+y
	}
	return x
}

/**

2.4. 赋值
	x = 1                       // 命名变量的赋值
	*p = true                   // 通过指针间接赋值
	person.name = "bob"         // 结构体字段赋值
	count[x] = count[x] * scale // 数组、slice或map的元素赋值

2.4.1. 元组赋值
x, y = y, x
a[i], a[j] = a[j], a[i]
i, j, k = 2, 3, 5


1.有些表达式会产生多个值，比如调用一个有多个返回值的函数。
f, err = os.Open("foo.txt") // function call returns two values

v, ok = m[key]             // map lookup
v, ok = x.(T)              // type assertion
v, ok = <-ch               // channel receive


2.map查找（§4.3）、类型断言（§7.10）或通道接收（§8.4.2）出现在赋值语句的右边时，并不一定是产生两个结果，也可能只产生一个结果
v = m[key]                // map查找，失败时返回零值
v = x.(T)                 // type断言，失败时panic异常
v = <-ch                  // 管道接收，失败时返回零值（阻塞不算是失败）

_, ok = m[key]            // map返回2个值
_, ok = mm[""], false     // map返回1个值
_ = mm[""]                // map返回1个值

和变量声明一样，我们可以用下划线空白标识符_来丢弃不需要的值。
_, err = io.Copy(dst, src) // 丢弃字节数
_, ok = x.(T)              // 只检测类型，忽略具体值




2.4.2. 可赋值性


赋值语句是显式的赋值形式，
但是程序中还有很多地方会发生隐式的赋值行为
medals := []string{"gold", "silver", "bronze"}

medals[0] = "gold"
medals[1] = "silver"
medals[2] = "bronze"
只有右边的值对于左边的变量是可赋值的，赋值语句才是允许的。




***/
