package main

import "fmt"

func modifyA(a int) {
	a = 10
}

func modifyAByPointer(a *int) {
	*a = 10
}

//new 和make
//new用于分配内存，返回指向类型的指针
//make用于分配内存，返回类型的值

func testNew() {
	a := new(int)
	fmt.Println(a)
	fmt.Println(*a)
}

func testMake() {
	a := make([]int, 0)
	fmt.Println(a)
}

// 在go中对于引用类型slice map channel 在使用的时候不仅要声明还要为其分配内存空间
func testMake2() {
	// 只声明了map，但是没有分配内存空间
	// var b map[string]int
	// b["key"] = 1
	// fmt.Println(b)

	//正确的做法
	b := make(map[string]int)
	b["key"] = 1
	fmt.Println(b)
}

//程序定义一个int变量num的地址并打印 将num的地址赋值给指针变量ptr 并修改num的值

func testPointer() {
	num := 1
	fmt.Println(num)
	fmt.Println(&num)

	ptr := &num
	*ptr = 20
	fmt.Println(num)
}

func main() {
	a := 1
	modifyA(a)
	fmt.Println(a)
	modifyAByPointer(&a)
	fmt.Println(a)
	testNew()
	testMake()
	testMake2()
	testPointer()
}
