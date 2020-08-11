package main

//同一个协程没有机会向chan连续发送2条消息,对方不处理完第一条消息一直阻塞
//不同的协程给同一个管道发消息只要管道的接收方一直不收消息就一直阻塞
import (
	"fmt"
	"time"
)

func main() {
	c := make(chan int)
	go func(chan int) {
		for {
			i := <-c
			fmt.Println("recive ", i)
			time.Sleep(time.Second * 5)
		}
	}(c)

	go func(chan int) {
		c <- 2
		fmt.Println("2 over")
	}(c)
	time.Sleep(time.Microsecond * 1)
	c <- 1
}
