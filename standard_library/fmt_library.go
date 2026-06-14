// package main

// /*
// fmt包实现了类似C语言printf和scanf的格式化IO 主要分为向外输出内容和输入内容两大部分
// */

// /*
// Print系列函数会将内容输出到系统的标准输出 区别在
// Print函数直接输出内容
// Printf函数支持格式化输出字符串
// Println函数会在输出内容的结尾添加一个换行符

// func Print(a ...any) (n int, err error) //写入的字节数n和错误信息err

// any == interface{} 可以接受任意数量任意类型的参数
// */

// // func add(a, b int) (int, error) {
// // 	return a + b, nil
// // }
// // func main() {
// // 	fmt.Print("hello word\n")
// // 	fmt.Println("hello word")
// // 	fmt.Printf("hello word %d", 1)

// // }

// /*
// Fprint系列函数会将内容输出到一个io.Writer接口类型的变量W变量中 我们通常工行会用这个函数网文件中写入内容
// */

// // func main() {
// // 	// func Fprint(w io.Writer, a ...interface{}) (n int, err error)
// // 	// func Fprintf(w io.Writer, format string, a ...interface{}) (n int, err error)
// // 	// func Fprintln(w io.Writer, a ...interface{}) (n int, err error)

// // 	//    func Fprint(w io.Writer, a ...any) (n int, err error) {
// // 	// 	p := newPrinter()
// // 	// 	p.doPrint(a)
// // 	// 	n, err = w.Write(p.buf)
// // 	// 	p.free()
// // 	// 	return

// // 	fmt.Fprintln(os.Stdout, "向标准输出写入内容")
// // 	fileObj, err := os.OpenFile("./test.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
// // 	if err != nil {
// // 		fmt.Println("打开文件错误", err)
// // 		return
// // 	}
// // 	str := "枯藤老树昏鸦，小桥流水人家，古道西风瘦马，夕阳西下，断肠人在天涯。"
// // 	fmt.Fprintf(fileObj, "向文件中写入的内容为:%s", str)
// // }

// //Sprint系列函数会把传入的数据生成并返回一个字符串

// /*
// func Sprint(a ...interface{}) string
// func Sprintf(format string, a ...interface{}) string
// func Sprintln(a ...interface{}) string
// */

// // func main() {
// // 	S1 := fmt.Sprint("aaaa")
// // 	fmt.Println(S1)
// // 	name := "luvias"
// // 	age := 12
// // 	S2 := fmt.Sprintf("name:%s,age:%d", name, age)
// // 	fmt.Println(S2)
// // }

// // Errorf。函数根据format参数生成格式化字符串并返回一个包含该字符串的错误
// // func Errorf(format string, a ...any) error {
// // func main() {
// // 	err := fmt.Errorf("this is a mistake")
// // 	fmt.Println(err)
// // 	fmt.Println(reflect.TypeOf(err)) //可以进一步获取字段、方法、Kind 等反射信息
// // 	fmt.Printf("%T\n", err)          //不用导包、代码短

// // }

// //获取输入
// /*
// fmt.Scan
// fmt.Scanf
// fmt.Scanln
// */

// // func main() {
// // 	var (
// // 		name    any
// // 		age     any
// // 		married any
// // 	)
// // 	fmt.Scan(&name, &age, &married)
// // 	fmt.Printf("输入结果，name:%v,age:%v,married:%v", name, age, married)
// // }
