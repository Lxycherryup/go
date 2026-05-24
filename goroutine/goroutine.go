package main

import (
	"fmt"
	"sync"
)

func sayHello() {
	fmt.Println("Hello World")
}

// 启动多个goroutine
var wg sync.WaitGroup

func sayHelloMulti(n int) {
	defer wg.Done() //defer 在函数结束时执行
	fmt.Println("Hello World", n) //goroutine不是顺序执行的 原因是：goroutine是并发执行的，不是顺序执行的
}

func main() {
	// go sayHello()
	// fmt.Println("main goroutine")
	// time.Sleep(time.Second)
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go sayHelloMulti(i)
	}
	wg.Wait()
}
