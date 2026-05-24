package main

//接口是一种类型
//一个接口可以有多种方法
//一个类型可以实现多个接口
//一个类型实现了接口中所有的方法，那么就叫做实现了这个接口
//值类型可以调用指针类型的方法，但是指针类型不能调用值类型的方法
import "fmt"

type People interface {
	Speak(string) string
}

type Student struct{}

func (stu *Student) Speak(think string) (talk string) {
	if think == "sb" {
		talk = "你是个大帅比"
	} else {
		talk = "您好"
	}
	return
}

func main() {
	var peo People = &Student{}
	think := "bitch"
	fmt.Println(peo.Speak(think))
}
