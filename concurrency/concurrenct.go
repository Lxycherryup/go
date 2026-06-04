package main

import (
	"sync"
	"time"
)

//go代码中可能存在多个goroutine同时操作一个资源(临界区)，这种情况发生竞态问题

// var x int64
// var wg sync.WaitGroup

// func add() {
// 	for i := 0; i < 5000; i++ {
// 		x = x + 1
// 	}
// 	wg.Done()
// }

// func main() {
// 	wg.Add(2)
// 	go add()
// 	go add()
// 	wg.Wait()
// 	println(x) //结果可能不是10000 因为存在竞态问题
// 	//上面的代码开启了两个goroutine去累加变量x的值 这两个goroutine在访问和修改x变量的时候就会存在数据竞争，导致最后的结果与预期的不符合

// }

//互斥锁
/*
互斥锁是一种常用的控制共享资源访问的方法，它能够保证同时只有一个goroutine可以访问共享资源
go语言中使用sync包的Mutex类型来实现互斥锁
下面将使用互斥锁来解决上面代码中的竞态问题
*/

// var x int64
// var wg sync.WaitGroup
// var mu sync.Mutex //创建一个互斥锁
// func add() {
// 	for i := 0; i < 5000; i++ {
// 		/*
// 			使用互斥锁能够保证同一时间有且只有一个goroutine进入临界区，其他goroutine则在等待锁
// 			当互斥锁释放后
// 			等待的goroutine才可以获取锁进去临界区
// 			多个goroutine同时等待一个锁时
// 			唤醒的策略是随机的
// 		*/
// 		mu.Lock() //加锁
// 		x = x + 1
// 		mu.Unlock() //解锁
// 	}
// 	wg.Done()

// }

// func main() {

// 	wg.Add(2)
// 	go add()
// 	go add()

// 	wg.Wait()
// 	println(x) //结果是10000 因为使用了互斥锁保证了同时只有一个goroutine可以访问和修改x变量
// }

//读写互斥锁
/*
互斥锁是完全互斥的
但是有很多实际昌吉下是读多写少的，比如一个共享资源被多个goroutine频繁读取但很少修改的情况
在这种情况下使用互斥锁会导致性能下降 因为读操作也需要获取锁 但是读操作之间是没有冲突的
读写互斥锁允许多个goroutine同时读取共享资源 但是在写操作时需要独占访问
go语言中使用sync包的RWMutex类型来实现读写互斥锁
*/

/*
读写锁分为两种 读锁和写锁。
当一个goroutine获取读锁之后。其他的goroutine如果是获取读锁会继续获得锁，但是如果是获取写锁就会被阻塞
当一个goroutine获取写锁之后。其他的goroutine无论是获取读锁还是写锁都会被阻塞
*/

var (
	x      int64
	wg     sync.WaitGroup
	lock   sync.Mutex
	RWlock sync.RWMutex
)

func write() {

	RWlock.Lock() //加写锁
	x = x + 1
	time.Sleep(time.Millisecond * 10) //假设读操作耗时10毫秒
	RWlock.Unlock()                   //释放读写锁
	wg.Done()
}

func read() {
	RWlock.RLock()                   //加读锁
	time.Sleep(time.Millisecond * 1) //假设读操作耗时10毫秒
	RWlock.RUnlock()                 //释放读锁
	wg.Done()
}

func main() {
	start := time.Now()
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go write()
	}
	for i := 0; i < 100; i++ {
	}

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go read()
	}
	wg.Wait()
	end := time.Now()
	println("Time taken:", end.Sub(start).Milliseconds(), "ms")
}
