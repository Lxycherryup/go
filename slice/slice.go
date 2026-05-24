package main

import (
	"fmt"
)

//切片是数组的一个引用类型，切片本身不存储任何数据，它只是描述了底层数组的一段连续的元素。切片具有动态大小，可以根据需要增长或缩小。
//切片的遍历方式和数组一样。可以使用len获取切片商都
//cap可以获取silce的最大扩张容量 不能超出数组的长度 0 <=len(slice) <=len(array) array是slice的引用
//如果slice == nil 则len(slice) == 0 cap(slice) == 0

//切片的创建方式
//1 数组切片
//2 make函数创建 make([]type,len,cap) cap可以省略 默认为len
//3 切片字面量 slice := []type{value1,value2,...} 这种方式创建的切片没有底层数组 需要编译器自动创建一个数组来存储这些值

func main() {

	//1.声明
	var slice1 []int
	fmt.Println("slice1:", slice1, "len:", len(slice1), "cap:", cap(slice1))

	slice2 := []int{}
	slice3 := make([]int, 5)
	//从数组进行切片
	slice4 := [5]int{1, 2, 3, 4, 5}

	slice5 := slice4[:4]

	fmt.Println("slice2:", slice2, "len:", len(slice2), "cap:", cap(slice2))
	fmt.Println("slice3:", slice3, "len:", len(slice3), "cap:", cap(slice3))
	fmt.Println("slice4:", slice4, "len:", len(slice4), "cap:", cap(slice4))
	fmt.Println("slice5:", slice5, "len:", len(slice5), "cap:", cap(slice5))

	//切片的遍历
	for i, v := range slice5 {

		fmt.Printf("index:%d , value:%d ,", i, v)
	}

	//append向切片尾部插入元素
	//slice5底层数组容量不足时，会重新分配更大的底层数组，并将原数组的元素复制到新数组中 通常是引用数组的长度的两倍
	slice5 = append(slice5, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6)
	fmt.Println("\nslice5 after append:", slice5, "len:", len(slice5), "cap:", cap(slice5))

	slice6 := make([]int, 5)
	fmt.Println("slice6:", slice6, "len:", len(slice6), "cap:", cap(slice6))

	slice7 := [5]int{1, 2, 3, 3, 4}
	//使用语法糖...进行append
	//...的用法
	slice7_1 := make([]int, 5)
	copy(slice7_1, slice7[:])

	slice8 := append(slice7[:], 5, 6, 7)
	fmt.Println("slice8:", slice8, "len:", len(slice8), "cap:", cap(slice8))

	slice9 := append(slice8, slice7_1...)
	fmt.Println("slice9:", slice9, "len:", len(slice9), "cap:", cap(slice9))

}
