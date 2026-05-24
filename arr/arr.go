package main

import "fmt"

func printArr(arr *[5]int) {

	arr[0] = 100
	for i, v := range arr {
		fmt.Printf("index: %d, value: %d\n", i, v)
	}
}

// 求所有数组元素之和
func sumArr(arr *[5]int) int {

	sum := 0
	for _, v := range arr {
		sum += v
	}
	return sum
}

// 找出数组中给定索引的两个数之和
func sumAtindex(arr *[10]int, index1 int, index2 int) int {
	return arr[index1] + arr[index2]
}
func main() {

	var arr1 [5]int
	printArr(&arr1)

	fmt.Println("arr1:", arr1)

	arr2 := [...]int{1, 2, 3, 4, 5}
	printArr(&arr2)

	arr3 := [5]int{10, 20, 30, 40, 50}
	sum := sumArr(&arr3)
	fmt.Println("arr3:", arr3)
	fmt.Println("sum of arr3:", sum)

	arr4 := [10]int{10, 20, 30, 40, 50, 60, 70, 80, 90, 100}
	result := sumAtindex(&arr4, 2, 5)
	fmt.Printf("arr4: %v\n", arr4)
	fmt.Printf("sum of arr4 at index 2 and 5: %d\n", result)
}
