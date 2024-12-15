package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan string)
	ch2 := make(chan string)

	go func() {
		time.Sleep(2 * time.Second)
		ch1 <- "from channel 1"
	}()

	go func() {
		time.Sleep(1 * time.Second)
		ch2 <- "from channel 2"
	}()

	select {
	case msg1 := <-ch1:
		fmt.Println(msg1)
	case msg2 := <-ch2:
		fmt.Println(msg2)
	case <-time.After(3 * time.Second): // 超时处理
		fmt.Println("timeout")
	}
}

/*

select 用于同时监听多个通道操作，它的行为类似于 switch，可以用于实现多路复用。



*/
