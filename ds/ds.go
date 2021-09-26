// 数据结构

// 数组
// 链表
// 栈
// 队列
// Set
// Map
// 堆
// 优先级队列

package main

import (
  "container/list"
  "fmt"
)

func DsArray() {
	nums := make([]int, 3)
	nums[0], nums[1], nums[2] = 10, 11, 12

	strs := []string{}
	strs[0], strs[1], strs[2] = "10", "11", "12"

	for i := 0; i < 5; i++ {
		fmt.Print(i, " ")
	}
	fmt.Println()

	idx := 0
	for idx < 5 {
		fmt.Print(idx, " ")
		idx++
	}
	fmt.Println()

	for i, v := range strs {
		fmt.Println("idx:", i, "val:", v)
	}
	for i, v := range nums {
		fmt.Println("idx:", i, "val:", v)
	}
}

func DsList() {
	students := list.New()
	students.PushBack(1)
	students.PushFront(2)
	for item := students.Front(); item != nil; item = item.Next() {
		fmt.Println(item.Value)
	}
}

/**
 * Golang保护性零值, 判断map元素是否存在
 */
func DsMap() {
	dict := map[string]int{}
	dict["1"] = 11
	dict["2"] = 22
	fmt.Println(dict["1"])
	v2, isExist := dict["2"]
	fmt.Println(v2, isExist)
	v3, isExist := dict["3"]
	fmt.Println(v3, isExist)

	dict0 := make(map[string]string)
	dict0["0"] = "0"

	delete(dict0, "0")

	v0, isExist := dict0["0"]
	fmt.Println(v0, isExist)
}


func main() {
	// DsArray()
	//DsMap()
	DsList()
}
