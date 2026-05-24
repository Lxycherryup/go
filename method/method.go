package main

import "fmt"

//当接受者是指针式。即使使用值类型调用那么函数内部也是对指针的操作

type Data struct {
	x int
}

func (self Data) ValueTest() {
	fmt.Printf("value : %p\n", &self)
}

func (self *Data) PointerTest() {
	fmt.Printf("Pointer : %p\n", self)
}

func main() {
	d := Data{}
	p := &d
	fmt.Printf("Data :%p\n", p)

	d.ValueTest()
	d.PointerTest()

	p.ValueTest()
	p.PointerTest()
}
