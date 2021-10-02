// 数据结构

// 数组
// 链表
// 栈
// 队列
// Set
// Map
// 堆(优先级队列)

package main

import (
	"container/list"
	"fmt"
	dsBinHeap "github.com/emirpasic/gods/trees/binaryheap"
	dsUtils "github.com/emirpasic/gods/utils"
)

func DsArray() {
	nums := make([]int, 3)
	nums[0], nums[1], nums[2] = 10, 11, 12

	strs := [3]string{}
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

	for _, v := range strs {
		fmt.Println("val:", v)
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

func DsHeap() {
	// Min-heap
	heap := dsBinHeap.NewWithIntComparator()  // empty (min-heap)
	heap.Push(2)                              // 2
	heap.Push(3)                              // 2, 3
	heap.Push(1)                              // 1, 3, 2
	fmt.Println(heap.Values())                // 1, 3, 2
	_, _ = heap.Peek()                        // 1,true
	_, _ = heap.Pop()                         // 1, true
	_, _ = heap.Pop()                         // 2, true
	_, _ = heap.Pop()                         // 3, true
	_, _ = heap.Pop()                         // nil, false (nothing to pop)
	heap.Push(1)                              // 1
	heap.Clear()                              // empty
	heap.Empty()                              // true
	fmt.Println("size: ", heap.Size())        // 0
	fmt.Println("not empty: ", !heap.Empty()) // 0

	// Max-heap
	inverseIntComparator := func(a, b interface{}) int {
		return -dsUtils.IntComparator(a, b)
	}
	heap = dsBinHeap.NewWith(inverseIntComparator) // empty (min-heap)
	heap.Push(2, 3, 1)                             // 3, 2, 1 (bulk optimized)
	fmt.Println(heap.Values())                     // 3, 2, 1

	// Min-heap
	arrayHeap := dsBinHeap.NewWith(func(a, b interface{}) int {
		// interface转array
		return dsUtils.IntComparator(a.([]int)[2], b.([]int)[2])
	})
	arr0, arr1, arr2 := []int{2, 2, 2}, []int{1, 1, 1}, []int{3, 3, 3}
	arrayHeap.Push(arr0)
	arrayHeap.Push(arr1)
	arrayHeap.Push(arr2)
	fmt.Println(arrayHeap.Values())

	arr, _ := arrayHeap.Pop()
	fmt.Println(arr.([]int))
}

func main() {
	//DsArray()
	//DsMap()
	//DsList()
	DsHeap()
	//height := [][]int{{1, 2, 3}, {1, 2, 3}}
	//test(height)
}

func test(height [][]int) {

	rLen, cLen := len(height), len(height[0])
	fmt.Println(rLen)
	fmt.Println(cLen)

	visited := make([][]bool, rLen)
	for i := range visited {
		visited[i] = make([]bool, cLen)
	}

	visited[1][1] = true
	fmt.Println(visited)

	result := 0
	result += 2
	fmt.Println(result)

}
