package main

import (
	"fmt"
	"math/rand"
)

// func sayHello() {
// 	fmt.Println("Hello World")
// }

// // 启动多个goroutine
// var wg sync.WaitGroup

// func sayHelloMulti(n int) {
// 	defer wg.Done() //defer 在函数结束时执行
// 	fmt.Println("Hello World", n) //goroutine不是顺序执行的 原因是：goroutine是并发执行的，不是顺序执行的
// }

// func main() {
// 	// go sayHello()
// 	// fmt.Println("main goroutine")
// 	// time.Sleep(time.Second)
// 	for i := 0; i < 10; i++ {
// 		wg.Add(1)
// 		go sayHelloMulti(i)
// 	}
// 	wg.Wait()
// }

// 启动多个goroutine
// var wg sync.WaitGroup

// func hello(i int) {
// 	defer wg.Done() //goroutine结束时协程-1
// 	fmt.Println("hello goroutine!", i)

// }

// func main() {
// 	for i := 0; i < 10; i++ {
// 		wg.Add(1) //协程+1
// 		go hello(i)

// 	}
// 	wg.Wait() //等待所有协程结束
// }

//worker pool 模式
// worker pool 模式是一种常见的兵并发模式
//本质上是生产消费模式 可以有效控制goroutine的数量 避免过多的goroutine导致系统资源耗尽

// 需求 计算一个数字的各个位数之和，例如数组123 结果为1+2+3=6
type job struct {
	Id      int //任务id
	RandNum int //需要计算的数字
}

type Result struct {
	job *job //任务
	sum int  //计算结果
}

func createWorkerPool(numWorkers int, jobChan chan *job, resultChan chan *Result) {
	//根据开启的协程个数 也就是worker的个数 来创建对应数量的goroutine
	for i := 0; i < numWorkers; i++ {
		go func(jobChan chan *job, resultChan chan *Result) {
			//执行运算 遍历job管道中的数据。进行相加 并将结果发送到result管道中
			for job := range jobChan {
				sum := 0
				n := job.RandNum
				for n != 0 {
					sum += n % 10
					n /= 10
				}
				result := &Result{
					job: job,
					sum: sum,
				}
				resultChan <- result
			}

		}(jobChan, resultChan)
	}
}

func main() {

	//需要两个channel
	jobChan := make(chan *job, 100)       //任务channel
	resultChan := make(chan *Result, 100) //结果channel
	//创建工作池
	createWorkerPool(5, jobChan, resultChan)
	//发送任务
	for i := 0; i < 100; i++ {
		jobChan <- &job{
			Id:      i,
			RandNum: rand.Intn(1000), //生成一个随机数
		}
	}
	close(jobChan) //关闭任务channel

	for i := 0; i < 100; i++ {
		result := <-resultChan
		fmt.Printf("Job ID: %d, RandNum: %d, Sum: %d\n", result.job.Id, result.job.RandNum, result.sum)
	}

}
