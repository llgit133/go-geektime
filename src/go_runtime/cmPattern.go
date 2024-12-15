package main

import "fmt"

func producer(ch chan int) {
	for i := 1; i <= 5; i++ {
		ch <- i
	}
	close(ch)
}

func consumer(ch chan int) {
	for value := range ch {
		fmt.Println("Consumed:", value)
	}
}

func main() {
	ch := make(chan int)

	go producer(ch)
	consumer(ch)
}

/*

生产者-消费者模式
发布订阅模型

*/
