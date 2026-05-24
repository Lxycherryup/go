package main

import (
	"fmt"
	"sort"
)

//map是一种无序的key value的数据结构 go中的map是一种引用数据类型。必须初始化才能使用

func testMap() {
	//声明map
	var map1 map[string]int
	//没有初始化的map读操作不会出发panic 但是写会触发panic

	fmt.Println(map1)
	// map1["age"] = 1

	//初始化map
	map1 = make(map[string]int)
	fmt.Println(map1)
	map1["age"] = 18
	fmt.Println(map1)
	map2 := map[string]int{}
	map2["age"] = 19
	fmt.Println(map2)

}

// go的包级别也就是函数外部 只允许声明变量和函数 不能执行语句
func testAnonymousFunction() {
	func() {
		fmt.Println("匿名函数")
	}()
}

// 使用map判断某个键是否存在
func testMapKeyExits() {
	map1 := make(map[string]int)
	map1["age"] = 18
	fmt.Println(map1)
	value, exists := map1["age"]
	fmt.Println(value, exists)

}

// 删除map中的键值对
func testMapDelete() {
	map1 := make(map[string]int)
	map1["age"] = 18
	fmt.Println(map1)
	delete(map1, "age")
	fmt.Println(map1)
}

// map的遍历
func testMapRange() {
	map1 := make(map[string]int)
	//赋值多个
	map1["age"] = 18
	map1["score"] = 199
	map1["name"] = 200
	map1["class"] = 201
	fmt.Println(map1)
	//map中的遍历是有序还是无序？无序
	//每次遍历的结果都不一样 为什么是无序的？因为map的底层实现是哈希表，哈希表的特性就是无序的
	//如何指定遍历顺序？可以先获取所有key，然后排序，再遍历
	keys := make([]string, 0, len(map1))
	for key := range map1 {
		keys = append(keys, key)
	}
	fmt.Println(keys)
	//对key进行排序
	sort.Strings(keys)
	fmt.Println(keys)

	for key, value := range map1 {
		fmt.Println(key, value)
	}
}

// 使用map实现对一个长度为1-n+1的n个数进行去重
func testMapDeduplication() {
	// TODO: 实现去重逻辑
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	map1 := make(map[int]bool, len(nums))

	for _, num := range nums {
		map1[num] = true
	}
	fmt.Println(map1)
	//将map转为去重后的slice
	result := make([]int, 0, len(map1))
	for key := range map1 {
		result = append(result, key)
	}
	//对result进行排序
	sort.Ints(result)
	//打印去重后的结果
	fmt.Println(result)
}
func main() {

	testMap()
	testAnonymousFunction()
	testMapKeyExits()
	// testMapDelete()
	testMapDelete()
	testMapRange()
	testMapDeduplication()
}
