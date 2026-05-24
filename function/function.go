package main

//闭包 = 函数 + 外部引用环境
//闭包可以访问外部函数的变量
//闭包可以修改外部函数的变量
//闭包可以作为返回值

func makeCounter() func() int {
	counter := 0 //外部引用环境
	return func() int {
		counter++
		return counter
	}
}

func main() {
	counter := makeCounter()
	println(counter())
	println(counter())

	println(counter())
}
