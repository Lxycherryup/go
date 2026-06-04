package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

/*
代码中加锁操作因为设计内核态的上下文切换回比较耗时。代价比较高
针对基本数据类型可以使用原子操作来保证并发安全，因为原子操作是go语言提供的方法在用户态就可以完成
因此性能比加锁操作更好
go语言中原子操作由内至的标准库sync/atomic提供

一个操作员吗全部完成 要么完全不执行 中间不会被其他goroutine打断
*/

// var x int64
// var wg sync.WaitGroup

// func add(x int64) int64 {

// 	return x + 1
// }

// func main() {
// 	wg.Add(5)
// 	go func() {
// 		defer wg.Done()
// 		for i := 0; i < 10000; i++ {
// 			x = add(x)
// 		}

// 	}()
// 	go func() {
// 		defer wg.Done()
// 		for i := 0; i < 10000; i++ {
// 			x = add(x)
// 		}

// 	}()
// 	go func() {
// 		defer wg.Done()
// 		for i := 0; i < 10000; i++ {
// 			x = add(x)
// 		}

// 	}()
// 	go func() {
// 		defer wg.Done()
// 		for i := 0; i < 10000; i++ {
// 			x = add(x)
// 		}

// 	}()
// 	go func() {
// 		defer wg.Done()
// 		for i := 0; i < 10000; i++ {
// 			x = add(x)
// 		}

// 	}()

// 	wg.Wait()
// 	fmt.Println(x)
// }

// 使用atomic
/*
常用的atomic操作
加 atomic.AddInt64(&count,1)
减 atomic.AddInt64(&count,-1)

读取Load
v := atomic.LoadInt64(&count)

对变量进行赋值

v := atomic.StoreInt64(&count,100)
相当于
count = 100

交换
swap
old :=atomic.SwapInt64(&count,200)
count = 100
执行后
old_count = 100
count = 200

比较并交换
CAS Compare And Swap


atomic.CompareAndSwapInt64(&count,100,200)
如果count == 100
则改成200
否则不改
最终返回true
表示修改成功

*/
var x int64
var count int64
var wg sync.WaitGroup

func main() {

	count = 100
	ok := atomic.CompareAndSwapInt64(&count, 100, 200)
	fmt.Println(ok)
	fmt.Println(count)
	// wg.Add(5)
	// for i := 0; i < 5; i++ {
	// 	go func() {
	// 		defer wg.Done()
	// 		for i := 0; i < 10000; i++ {
	// 			atomic.AddInt64(&x, 1)
	// 		}
	// 	}()
	// }
	// wg.Wait()
	// fmt.Println(x)

}
