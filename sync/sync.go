package main

import (
	"fmt"
	"strconv"
	"sync"
)

/*
在代码中生硬的使用time.Sleep来控制协程的执行顺序和时间是非常不可靠的，因为它可能会导致不可预测的行为和性能问题
go语言提供了sync包中的WaitGroup类型来实现协程之间的同步和等待
WaitGroup可以用来等待一组协程完成它们的工作
使用WaitGroup的基本步骤如下：

1.创建一个WaitGroup变量
2.在每个需要等待的协程中调用Add方法增加计数器的值
3.在每个协程完成工作后调用Done方法减少计数器的值
4.在主协程中调用Wait方法阻塞等待直到计数器的值为0
*/

// var wg sync.WaitGroup

// func hello() {
// 	defer wg.Done() //在函数结束时调用Done方法减少计数器的值
// 	fmt.Println("Hello, World!")

// }

// func main() {
// 	wg.Add(1)
// 	go hello() //开启一个协程执行hello函数
// 	fmt.Println("Waiting for goroutine to finish...")
// 	wg.Wait() //等待协程完成
// 	fmt.Println("Goroutine finished.")

// }

/*
sync.Once
在编程的很多场景下。我们需要确保某些操作在高并发场景下只执行一次
例如只夹在一次配置文件，只关闭一次通道等
*/

/*sync.Map*/
//go语言中内置的Map不是并发安全的 如下

var m = sync.Map{}

func main() {

	wg := sync.WaitGroup{}
	for i := 0; i < 20; i++ {
		wg.Add(1)
		go func(n int) {
			key := strconv.Itoa(n)
			m.Store(key, n)
			v, _ := m.Load(key)
			fmt.Printf("k=:%v,v:=%v\n", key, v)
			wg.Done()
		}(i)
	}
	wg.Wait()
}
