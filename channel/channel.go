package main

import "fmt"

// import "fmt"

// //channel是一种引用类型
// //声明各位为 var 变量名 chan 数据类型
// //channel的零值是nil nil channel不能发送和接收数据 否则会触发panic
// //channel必须初始化才能使用 可以使用make函数来初始化

// func receive(ch chan int) {
// 	value := <-ch
// 	println("Received:", value)
// }

// func main() {
// 	// var ch chan int
// 	// fmt.Println(ch) //nil
// 	//向未初始化的channel发送数据会触发panic
// 	// ch <- 1
// 	// ch = make(chan int)
// 	// fmt.Println(ch) //0xc0000a2000
// 	// ch <- 1        //向channel发送数据
// 	// value := <-ch  //从channel接收数据
// 	// fmt.Println(value)
// 	// ch1 := make(chan int, 1) //带缓冲的channel 1表示缓冲区大小为1
// 	// ch2 := make(chan int)
// 	// fmt.Println(ch1, ch2) //0xc0000a2000 0xc0000a4000
// 	// ch1 <- 1
// 	// fmt.Println(<-ch1)
// 	//channel 有发送send 接受recevie 关闭close三种操作
// 	//发送和接受都使用<-符号 发送在通道变量在左边 接收通道变量在右边
// 	//关闭channel close函数只能关闭channel 不能关闭其他类型的变量 关闭一个已经关闭的channel会触发panic
// 	// ch := make(chan int, 1)
// 	// ch <- 1
// 	// x := <-ch
// 	// fmt.Println(x)
// 	// close(ch)
// 	// //对一个已经关闭的通道再进行关闭会触发panic
// 	// // close(ch)
// 	// //关闭channel后不能再发送数据 但是可以接收数据 接收不到数据会返回channel类型的零值
// 	// // ch <- 2 //panic: send on closed channel
// 	// y := <-ch
// 	// fmt.Println(y) //0

// 	//无缓冲通道又称为阻塞通道 发送的同时必须有协程接收数据。否则会触发deadlock
// 	// ch := make(chan int)
// 	// // ch <- 1 //deadlock
// 	// go receive(ch)
// 	// ch <- 1

// 	//close
// 	ch := make(chan int)
// 	go func() {

// 		for i := 0; i < 5; i++ {

// 			ch <- i

// 		}
// 		close(ch)
// 	}()

// 	for {
// 		if value, ok := <-ch; ok {
// 			fmt.Println("Received:", value)
// 		} else {
// 			fmt.Println("Channel closed")
// 			break
// 		}
// 	}

// 	fmt.Println("Main goroutine exiting")

// }
func counter(out chan<- int) {

	for i := 0; i < 100; i++ {
		out <- i
	}
	close(out)

}

func squarer(out chan<- int, in <-chan int) {
	for v := range in {
		out <- v * v
	}
	close(out)
}

func printer(in <-chan int) {
	for v := range in {
		fmt.Println(v)
	}
}

func main() {
	//单项通道 有的时候会将通道作为函数参数传递，但是函数只需要发送或者接收数据 这时候可以使用单向通道 来限制通道的使用方向
	//单向通道只能发送或者接收数据 不能同时发送和接收数据
	//单向通道的声明方式 var 变量名 chan<- 数据类型 或者 var 变量名 <-chan 数据类型
	//chan<-表示只能发送数据 <-chan 表示只能接收数据
	in := make(chan int)
	out := make(chan int)

	go counter(in)
	go squarer(out, in)
	printer(out)

}
