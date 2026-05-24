package main

import (
	"fmt"
)

// 递归实现（简洁但性能差）
// func fibRecursive(n int) int {
// 	if n <= 1 {
// 		return n
// 	}
// 	return fibRecursive(n-1) + fibRecursive(n-2)
// }

func fib(n int) int {
	if n <= 1 {
		return n
	}
	return fib(n-1) + fib(n-2)
}

// 迭代实现（推荐）
func fibIterative(n int) int {
	if n <= 1 {
		return n
	}
	a, b := 0, 1
	for i := 2; i <= n; i++ {
		a, b = b, a+b
	}
	return b
}

// 生成前 n 项数列
func fibSequence(n int) []int {
	seq := make([]int, n)
	for i := range seq {
		seq[i] = fibIterative(i)
	}
	return seq
}

func main() {
	fmt.Println("前 10 项斐波那契数列:")
	fmt.Println(fibSequence(10))
	fmt.Println("buildString 结果:", buildString())
	// fmt.Println("\n第 10 项（递归）:", fibRecursive(10))
	// fmt.Println("第 10 项（迭代）:", fibIterative(10))
}
