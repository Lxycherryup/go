// package main

// import (
// 	"encoding/json"
// 	"fmt"
// )

// /*
// 数据格式
// 是系统中数据交互不可缺少的内容
// 例如 JSON XML MSGPack
// */

// //JSON
// /*
// JSON是完全独立于语言的文本格式。是k-v的形式
// 应用场景 前后端数据交互。 系统间数据交互

// JSON使用go语言内置的encoding/json标准库
// 编码json使用json.Marshal()函数对一组数据进行JSON格式的编码
// */

// /*
// func Marshal(v any) ([]byte, error) 返回的是byte类型。要想格式化输出就要转为string
// 要使用Marshal的结构体中的字段名必须为大写开头
// */

// func main() {

// 	type person struct {
// 		Name string
// 		Age  string
// 	}

// 	jsonStr := `{"Name":"lisi","Age":"19"}`

// 	p1 := person{
// 		Name: "lisi",
// 		Age:  "19"}

// 	res, _ := json.Marshal(p1)
// 	fmt.Printf("p1 =%s\n", res)
// 	fmt.Printf("p1 = %v\n", string(res))
// 	fmt.Println("p1 =", string(res))

// 	var p person
// 	json.Unmarshal([]byte(jsonStr), &p)
// 	//格式化输出时候 %v只输出值 %+v输出key+值 %#v 输出完整类型+key+值
// 	fmt.Printf("%v", p)
// 	fmt.Printf("%+v", p)
// 	fmt.Printf("%#v", p)
// }
