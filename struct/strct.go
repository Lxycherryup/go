package main

import (
	"encoding/json"
	"fmt"
)

//go 语言中没有类的概念 也不支持继承等面向对象的概念 go语言中通过结构体的内嵌再配合接口支持代码的组合

func defineType() {
	//自定义数据类型
	type myInt int
	var num myInt = 1
	fmt.Println(num)
	fmt.Printf("%T\n", num)

	//类型别名
	type myInt2 = int
	var num2 myInt2 = 2
	fmt.Println(num2)
	fmt.Printf("%T\n", num2)
}

// 结构体的定义
type Person struct {
	//结构体中的字段名是唯一的
	//字段名首字母大写表示public
	Name string
	//字段名首字母小写表示private
	age int
}

// 结构体实例化
func structInstance() {
	//1.声明一个结构体变量
	var p Person
	p.Name = "张三"
	//私有字段可以在包内访问 不能在包外访问
	p.age = 18
	fmt.Println(p)
}

// 创建指针类型的结构体
func pointerStructInstance() {
	//2.使用var声明结构体变量
	var p2 *Person = new(Person)
	p2.Name = "李四"
	p2.age = 20
	fmt.Println(p2)
}

// 什么是方法和接受者
// 方法是作用于特定类型对象的函数
// 接受者是方法作用的对象
// 方法和接受者
func (p Person) print() {
	fmt.Println(p)
}

// 带指针接受者的方法
func (p *Person) setAge(age int) {
	p.age = age
}

// 结构体和json
func structToJson() {
	//结构体转json
	p := Person{
		Name: "张三",
		age:  18,
	}

	fmt.Println("结构体:", p)
	jsonBytes, _ := json.Marshal(p)
	fmt.Println("JSON:", string(jsonBytes))

	// json转结构体
	var p2 Person
	json.Unmarshal(jsonBytes, &p2)
	fmt.Println("结构体:", p2)
}

func main() {
	defineType()
	structInstance()
	pointerStructInstance()
	structToJson()
}
