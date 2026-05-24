package main

// go中的switch和其他语言有何不同？
// 1. go中的switch可以不写break，会自动break
// 2. go中的switch可以写多个条件，用逗号分隔
// 3. go中的switch可以写表达式，用:=分隔
// 4. go中的switch可以写类型，用type分隔
// example
func testSwitchType() {
	//interface{} 是万能类型 ，可以接受任何类型的值
	var i interface{}
	i = 1
	switch i.(type) {
	case int:
		println("int")
	case string:
		println("string")
	default:
		println("unknown")
	}

}

//5. go中的switch可以写nil，用nil分隔
//6. go中的switch可以写interface，用interface分隔
//7. go中的switch可以写结构体，用结构体分隔
//8. go中的switch可以写数组，用数组分隔
//9. go中的switch可以写切片，用切片分隔
//10. go中的switch可以写map，用map分隔
//11. go中的switch可以写channel，用channel分隔
//12. go中的switch可以写函数，用函数分隔
//13. go中的switch可以写指针，用指针分隔
//14. go中的switch可以写接口，用接口分隔
//15. go中的switch可以写类型断言，用类型断言分隔
//16. go中的switch可以写类型转换，用类型转换分隔
//17. go中的switch可以写类型推断，用类型推断分隔
//18. go中的switch可以写类型匹配，用类型匹配分隔
//19. go中的switch可以写类型判断，用类型判断分隔
//20. go中的switch可以写类型转换，用类型转换分隔

//介绍go中的select
//select是go中的关键字，用于选择多个channel中的一个进行通信
//select可以配合default使用，default用于处理没有channel可用的情况

func main() {
	testSwitchType()
}
