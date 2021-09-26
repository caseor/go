package main

import (
	"fmt"
	"math"
	"strings"
	"time"
)

const (
	A = iota
	B
	C
)

func format() {
	flag := true
	fmt.Printf("%v", flag)
}
func str() {
	str := "hhh hello world"
	print(strings.Replace(str, " ", "", -1))
}
func ptr() {
	var num = 4
	var ptr *int
	ptr = &num
	print(ptr)
}

func circle() {
	key := 0
	strs := []string{"a", "b", "c"}
	nums := []int{100, 200, 300}
	for i := 0; i < 5; i++ {
		print(i, " ")
	}
	println()
	for key < 5 {
		print(key, " ")
		key++
	}
	println()
	for i, v := range strs {
		println("idx:", i, "val:", v)
	}
	for i, v := range nums {
		println("idx:", i, "val:", v)
	}
}

func swap(x, y string) (string, string) {
	return y, x
}

func printSwap() {
	a, b := "Google", "Runoob"
	println(&a, &b)
	swap(a, b)
	println(&a, &b)
	fmt.Println(a, b)
}

type Circle struct {
	radius float64
}

func (circle Circle) getArea() float64 {
	return math.Pi * circle.radius * circle.radius
}

/**
 * 传入引用
 */
func (circle *Circle) changeRadius(radius float64) {
	circle.radius = radius
}

/**
 * 传入引用
 */
func changeRadius(circle *Circle, radius float64) {
	circle.radius = radius
}

func changeReferRadius() {
	var circle Circle
	circle.radius = 1
	println(circle.radius)
	circle.changeRadius(2)
	println(circle.radius)
	changeRadius(&circle, 3)
	println(circle.radius)
}

func array() {
	arr := []int{1: 1, 3: 2, 5: 3}
	for i, v := range arr {
		print(i, v, " ")
	}
}
func sliceArr() {
	arr := []int{1, 2, 3, 4, 5}
	// slice范围, 前闭后开[begin, end)
	printArr(arr[2:4])
	printArr(arr[:2])
	printArr(arr[2:])
	println("len(arr)", len(arr))
	println("cap(arr)", cap(arr))

	// arr[2:] ptr指向第2个元素, 从ptr开始后面还有3个, cap=3
	println("len(arr[2:])", len(arr[2:]))
	println("cap(arr[2:])", cap(arr[2:]))

	// arr[:2] ptr指向第0个元素, 从ptr开始后面还有5个, cap=5
	println("len(arr[:2])", len(arr[:2]))
	println("cap(arr[:2])", cap(arr[:2]))

	// arr[2:3] ptr指向第2个元素, 从ptr开始后面还有3个, cap=3
	println("len(arr[2:3])", len(arr[2:3]))
	println("cap(arr[2:3])", cap(arr[2:3]))

}
func deliverArray(arr []int) {
	arr[1] = 0
	//print(arr)
}

func printArr(arr []int) {
	fmt.Printf("%v\n", arr)
}
func prt2ptr() {
	var val = 1
	var ptr *int
	var pptr **int
	ptr = &val
	pptr = &ptr
	println(val)
	println(*ptr)
	println(*pptr)
	println(**pptr)
}

func printMap(dict map[string]string) {
	for key, val := range dict {
		println(key, val)
	}
}
func mapDemo() {
	dict := map[string]string{}
	dict["a"] = "1"
	dict["b"] = "2"
	dict["c"] = "3"
	printMap(dict)
	delete(dict, "b")
	printMap(dict)
}

type Phone interface {
	call()
}

type IPhone struct {
}

type Huawei struct {
}

func (iphone IPhone) call() {
	println("this is iPhone")
}

func (huawei Huawei) call() {
	println("this is Huawei")
}

func interfaces() {
	iphonex := new(IPhone)
	iphonex.call()

	p40pro := new(Huawei)
	p40pro.call()
}

func concurrence(val string) {
	for i := 0; i < 5; i++ {
		time.Sleep(time.Second)
		println(val)
	}
}

func chanSum(arr []int, channel chan int) {
	sum := 0
	for _, v := range arr {
		sum += v
	}
	channel <- sum
}
func chanType() {
	arr := []int{1, 2, 3, 4, 5, 6}
	channel := make(chan int)
	go chanSum(arr[:len(arr)/2], channel)
	go chanSum(arr[len(arr)/2:], channel)
	x, y := <-channel, <-channel
	println(x)
	println(y)
}

func chanWithCache() {
	ch := make(chan int, 3)
	ch <- 1
	ch <- 2
	println(<-ch)
	//println(<-ch)
	//println(<-ch)
}

func fibonacci(n int, ch chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		ch <- x
		x, y = y, x+y
	}
	close(ch)
}

func chanFibonacci() {
	ch := make(chan int, 10)
	go fibonacci(10, ch)
	for v := range ch {
		println(v)
	}
}

func main() {
	//x := 1
	//fmt.Println(x) //prints 1
	//{
	//	fmt.Println(x) //prints 1
	//	x = 3
	//	//x := 2         // 不会影响到外部x变量的值
	//	fmt.Println(x) //prints 2
	//	//x = 5        // 不会影响到外部x变量值
	//}
  x:=1
	fmt.Println(x) //prints 3
	//getSqrt := func(val float64) float64 {
	//  return math.Sqrt(val)
	//}
	//chanFibonacci()
}
