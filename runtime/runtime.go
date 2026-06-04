package main

import (
	"fmt"
	"runtime"
)

func main() {
	// go func(s string) {

	// 	for i := 0; i < 2; i++ {
	// 		fmt.Println(s)
	// 	}
	// }(", World!")

	// //主协程
	// for i := 0; i < 2; i++ {
	// 	//切一下 再次分配任务
	// 	runtime.Gosched()
	// 	fmt.Println("Hello")
	// }
	fmt.Println(runtime.NumCPU())       //获取CPU数量
	fmt.Println(runtime.NumGoroutine()) //获取当前正在运行的goroutine数量

}
