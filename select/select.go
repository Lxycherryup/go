package main

import (
	"fmt"
	"time"
)

//在某些场景下需从多个channel中接接收数据。通道在接收数据时，如果没有数据可以接受 将会发生阻塞
//为了解决这个问题 go内置了select语句 可以同时响应多个通道的操作

//select语句类似switch 语句 但是select语句中的case必须是一个通道操作
//select语句会随机执行一个case 如果多个case都满足条件 则会随机选择一个执行
//如果没有case满足条件 则会执行default语句 如果default语句也没有 则select语句会一直阻塞直到有case满足条件

//例如
// select {

// 	case msg1 := <-ch1:
// 		fmt.Println("Received", msg1) //如果从ch1 接收到数据 则执行这个case

// 	case ch2<-msg2:
// 		fmt.Println("Sent", msg2) //如果向ch2 发送数据成功 则执行这个case

// 	default:
// 		fmt.Println("No communication") //如果没有case满足条件 则执行这个case
// }

//select可以同时监听一个或多个channel 直到其中一个channel ready

// func main() {

// 	output1 := make(chan string)
// 	output2 := make(chan string)

// 	go func(ch chan string) {
// 		time.Sleep(2 * time.Second)
// 		ch <- "Hello from output1"
// 	}(output1)

// 	go func(ch chan string) {
// 		time.Sleep(1 * time.Second)
// 		ch <- "Hello from output2"
// 	}(output2)

// 	//使用select监听两个channel
// 	select {
// 	case msg1 := <-output1:
// 		println("Received:", msg1)
// 	case msg2 := <-output2:
// 		println("Received:", msg2)
// 	}
// }

// 如果多个channel同时ready select会随机选择一个执行 例如
// func main() {
// 	output1 := make(chan string)
// 	output2 := make(chan string)

// 	go func(ch chan string) {
// 		time.Sleep(1 * time.Second)
// 		ch <- "Hello from output1"
// 	}(output1)

// 	go func(ch chan string) {

// 		time.Sleep(1 * time.Second)
// 		ch <- "Hello from output2"
// 	}(output2)

// 	select {
// 	case msg1 := <-output1:
// 		println("Received:", msg1)
// 	case msg2 := <-output2:
// 		println("Received:", msg2)

// 	}
// }

// 可以使用select来判断管道是否存满
func main() {
	ch := make(chan string, 10)
	//开启子协程协数据
	go writer(ch)

	for s := range ch {
		fmt.Println("Received:", s)
		time.Sleep(time.Second * 1)
	}

}

func writer(ch chan string) {
	for {
		select {
		//写数据
		case ch <- "data":
			fmt.Println("Data written to channel")
			//如果管道满了 就执行default
		default:
			fmt.Println("Channel is full, cannot write data")

		}
		time.Sleep(time.Millisecond * 500)
	}
}
