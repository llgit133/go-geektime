package main

import "fmt"

// 无缓冲通道
func noBufferedChannel() {
	ch := make(chan int) // 创建无缓冲通道

	// 启动 Goroutine 发送数据
	go func() {
		ch <- 42
	}()

	// 接收数据
	value := <-ch
	fmt.Println("Received:", value)
}

// 创建缓冲区大小为 2 的通道
// 缓存队列
// 出入队列
func hasBufferChannel() {
	ch := make(chan int, 2)

	ch <- 1
	ch <- 2

	fmt.Println(<-ch) // 输出 1
	fmt.Println(<-ch) // 输出 2
}

// 通道方向
func direction() {
	ch := make(chan int)

	go send(ch)
	receive(ch)
}

// 只发送
func send(ch chan<- int) {
	ch <- 42
}

// 只接收
func receive(ch <-chan int) {
	fmt.Println("Received:", <-ch)
}

func main() {
	noBufferedChannel()
	hasBufferChannel()
	direction()
}

/*


select 用于同时监听多个通道操作，它的行为类似于 switch，可以用于实现多路复用。

*/
