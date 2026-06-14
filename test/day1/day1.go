package main

/*------------------------------------------------------------------------------------------*/
/*
go1.21以前 range 循环里的va变量会被复用。所以每次&val取到的都是同一个地址 最后都指向最后一次的值

从go 1.22以后 for range中用:=声明的循环变量。每次迭代都会创建新的变量
所以每次&val都是不同的地址

会输出
1 -> 1
2 -> 2
3 -> 3
0 -> 0

顺序不固定是因为map遍历本来就是无序的
*/

// func main() {
// 	slice := []int{0, 1, 2, 3}
// 	m := make(map[int]*int)

// 	for key, val := range slice {
// 		m[key] = &val
// 	}

// 	for k, v := range m {
// 		fmt.Println(k, "->", *v)
// 	}
// }
/*------------------------------------------------------------------------------------------*/

// func main() {

// 	s1 := make([]int, 5)
// 	//初始化切片时 元素为当前类型的零值
// 	s1 = append(s1, 1, 2, 3)
// 	fmt.Println(s1)

// 	s2 := make([]int, 0)
// 	s2 = append(s2, 1, 2, 3)
// 	fmt.Println(s2)

// }
/*------------------------------------------------------------------------------------------*/

/*

$ go test .
# demo/test/day1
./day1.go:48:38: syntax error: missing parameter type
FAIL    demo/test/day1 [build failed]
FAIL
在go中函数有多个返回值时候，只要有一个返回值有命名 其他的也必须命名
如果有多个返回值必须加上括号()
如果只有一个返回值且命名也必须加上括号
*/
// func muliti(x, y int) (sum int ,error){

// 	return x+y ,nil
// }

// func main() {

// 	muliti()
// }

/*------------------------------------------------------------------------------------------*/

/*
new()和make()的区别
new(T)和make(T,args)是gp语言内建函数 用来分配内存。但适用的类型不同

new(T)会为T类型的新值分配已经置0的内存空间。并返回地址，即类型为*T的值
换句话说就是返回一个指针 该指针指向新分配的  类型为T的零值。适用于值类型。如数组。结构体等


make(T,args)返回初始化之后的T类型的值。这个值并不是T类型的零值。也不是指针*T，是经过初始化之后T的引用。
make只适用于slice，map，channel

*/

/*------------------------------------------------------------------------------------------*/

// func main() {

// 	list := new([]int)
// 	//invalid append: argument must be a slice; have list (variable of type *[]int)
// 	list = append(list, 1)
// 	//append 的参数必须是slice，new返回的是一个类型指针
// 	fmt.Println(list)

// }

// func main() {
// 	/*
// 		数组写法 数组和slice的区别是  数组长度固定 slice长度不固定 slice := []int{1,2,3,4,5} 长度不固定
// 		数组的长度也是类型的一部分 [3]int [5]int是两种数据类型 数组长度固定 值类型。赋值会整体拷贝(深拷贝)
// 	*/

// 	/*
// 		深拷贝和浅拷贝的区别
// 		a := [3]int{1,2,3}
// 		b := a
// 		b[0] = 100
// 		fmt.Println(a) //[1 2 3]

// 		a := []int{1,2,3}
// 		b := a
// 		b[0] = 100
// 		fmt.Println(a) //[100,2,3] 因为切片时引用类型，底层共享共一块内存 也就造就了浅拷贝
// 	*/
// 	s1 := [...]int{1, 2, 3, 4, 5}
// 	s2 := [...]int{6, 7, 8, 9, 10}
// 	result := append(s1[:], s2[:]...)
// 	fmt.Println(result)

// }

/*------------------------------------------------------------------------------------------*/

// 结构体比较
// func main() {
// 	sn1 := struct {
// 		age  int
// 		name string
// 	}{age: 11, name: "qq"}

// 	sn2 := struct {
// 		age  int
// 		name string
// 	}{age: 11, name: "qq"}

// 	if sn1 == sn2 {
// 		fmt.Println("sn1 == sn2")
// 	}

// 	sm1 := struct {
// 		age int
// 		m   map[string]string
// 	}{age: 11, m: map[string]string{"a": "1"}}
// 	sm2 := struct {
// 		age int
// 		m   map[string]string
// 	}{age: 11, m: map[string]string{"a": "1"}}

// 	if sm1 == sm2 {
// 		fmt.Println("sm1 == sm2")
// 		/*
// 		./day1.go:149:5: invalid operation: sm1 == sm2
// 		(struct containing map[string]string cannot be compared)
// 		*/

// 		/*
// 		可比较的类型（Comparable）
// 		整数、浮点数、复数、布尔值
// 		字符串
// 		指针
// 		数组（如果元素可比较）
// 		结构体（如果所有字段可比较）
// 		接口

// 		不可比较的类型（Uncomparable）
// 		map - 引用类型，内容可变
// 		slice - 动态数组，无法定义比较语义
// 		function - 函数无法比较
// 		*/
// 	}

// }

/*------------------------------------------------------------------------------------------*/
