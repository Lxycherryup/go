package main

import (
	"context"
	"fmt"
	"time"
)

/*
context是go后端开发中最重要的标准库之一

核心作用是:在多个Groutine、多个函数调用之间传递控制信号和上下文信息
一个会沿着调用链不断向下传递的控制器

主要解决三个问题
超时控制(Timeout)
取消任务
请求级数据传递

为什么需要Context
假设有这样一个场景
func query DB(){
     time.Sleep(10*time.Second)
	 fmt.Println("查询完成")
}


http请求
func handler(w http.ResponseWriter, r *http.Request){
     queryDB()
}

用户执行http请求 等了一秒就关闭请求了。但是数据库还在查询。
这时候就会造成 客户端已经走了。服务端还在干活。浪费资源
context就是来告诉下游任务 别干了 请求已经取消了
*/

/*
Context源码是一个接口
type Context interface{
	Deadline() (deadline time.Time, ok bool)
	Done(). <-chan struct
	Err() error
	Value(key any) any
}

Done() <-chan 返回一个channel，当context被取消是close(done)
所有监听这个channel的groutine都会收到通知

Err() ctx.Err() 返回context.Canceled 或者 context.DeadlineEXceeded

Deadline() 查看超过时间 deadline,ok := ctx.Deadline()

Value()传递请求级数据
ctx.Value(key) 可以当成map使用
例如在上下文传递用户Id TraceId 请求链路信息
*/

// func worker(ctx context.Context) {
// 	for {
// 		select {
// 		case <-ctx.Done():
// 			fmt.Println("任务结束")
// 			return
// 		default:
// 			fmt.Println("working>>")
// 			time.Sleep(time.Second)
// 		}
// 	}
// }

// func main() {
// 	// 创建根context
// 	rootCtx := context.Background()

// 	//WithCancel  创建可以取消的context
// 	//返回func WithCancel(parent Context) (ctx Context, cancel CancelFunc)
// 	ctx, cancel := context.WithCancel(rootCtx)
// 	go worker(ctx)
// 	time.Sleep(3 * time.Second)
// 	cancel()
// 	time.Sleep(time.Second)
// }

//WithTimeout

func worker(ctx context.Context) {

	select {
	case <-time.After(5 * time.Second):
		fmt.Println("完成")
	case <-ctx.Done():
		fmt.Println(ctx.Err())
	}
}

func main() {
	ctx, cancel := context.WithTimeout(
		context.Background(),
		3*time.Second,
	)

	defer cancel()
	worker(ctx)
}
