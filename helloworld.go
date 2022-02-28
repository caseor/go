package main

import (
	"errors"
	"flag"
	"fmt"
	"math"
	"os"
	"strings"
	"time"
)

const (
	A = iota
	B
	C
)

// GlobalFun 1. invoke func to init var GlobalFun
var GlobalFun = globalFunPrintStr()

// init 2. invoke after global variable init completed
func init() {
	fmt.Printf("init() before main() but after golbal variable\n")
}

// main 3. do main
func main() {
	parseCommandLine()
}

func globalFunPrintStr() string {
	fmt.Printf("This is globalFunPrintStr")
	return "This is globalFunPrintStr"
}

func parseCommandLine() {
	// go run ./helloworld.go -u caseor -p 123456 -h 127.0.0.1 -P 10000
	//var user string
	//var password string
	//var host string
	//var port int
	//flag.StringVar(&user, "u", "root", "username, default 'root'")
	//flag.StringVar(&password, "p", "", "password, default ''")
	//flag.StringVar(&host, "h", "localhost", "host, default 'localhost'")
	//flag.IntVar(&port, "P", 3306, "port, default 3306")
	//flag.Parse()
	//fmt.Printf("[user=%v, password=%v, host=%v, port=%v]", user, password, host, port)

	// return type is pointer
	user1 := flag.String("u", "root", "username, default 'root'")
	password1 := flag.String("p", "", "password, default ''")
	host1 := flag.String("h", "localhost", "host, default 'localhost'")
	port1 := flag.Int("P", 3306, "port, default 3306")
	flag.Parse()
	fmt.Printf("[user=%v, password=%v, host=%v, port=%v]", *user1, *password1, *host1, *port1)
}

func argsTest() {
	// go run ./helloworld.go I am a boy
	// There is no way pass args by formal parameter
	args := os.Args
	argsLen := len(args)
	for i := 0; i < argsLen; i++ {
		fmt.Printf("Parameter: %v, %v\n", i, args[i])
	}
}

func executionTimeTest() {
	start := time.Now()
	time.Sleep(time.Duration(10) * time.Second)
	elapsed := time.Since(start)
	fmt.Printf("elapsed=%v\n", elapsed)
}

func recoverTest() {
	// recover is used with defer
	// just like try catch
	defer func() {
		fmt.Printf("defer start\n")
		if err := recover(); err != nil {
			fmt.Println(err)
		}
		fmt.Printf("defer end\n")
	}()

	// index out range precedes panic
	arr := []string{"a", "b"}
	fmt.Printf("%v\n", arr[2])
	panic("runtime panic")

	fmt.Println("recoverTest end")
}

func panicTest(val int) {
	// usually fatal error
	if val == 0 {
		panic("val can not be zero")
	}
	fmt.Printf("val=%v", val)
}

func div(dividend, divisor int) (int, error) {
	if divisor == 0 {
		return 0, errors.New("division by zero")
	}
	return dividend / divisor, nil
}

func deferTest(val, sum int) {
	defer fmt.Printf("%v, %v \n", val, sum)
	defer fmt.Printf("%v, %v \n", val+1, sum+1)
	val = val + 1
	sum = sum + 2
}

func closureTest() {
	f := closureAdd()
	fmt.Printf("%v\n", f(1))
	fmt.Printf("%v\n", f(2))
	fmt.Printf("%v\n", f(3))
	fmt.Printf("%v\n", f(4))
}
func closureAdd() func(int) int {
	i := 0
	return func(x int) int {
		i = i + x
		return i
	}
}

func callbackMain() {
	callbackTest(2, func(i int) {
		fmt.Printf("%v", i)
	})
}

func callbackTest(val int, callback func(int)) {
	fmt.Printf("execute biz logic\n")
	callback(val * 2)
}

func anonymousFun() {
	f := func(val int, num int) {
		fmt.Printf("%v, %v ", val, num)
	}
	f(1, 2)
}

func getStaticStr(val int) string {
	return "hello" + fmt.Sprintf("%v", val)
}

func dynamicInterfacePrams(args ...interface{}) {
	for k, v := range args {
		fmt.Printf("%v, %v, %T\n", k, v, v)
	}
}

func dynamicPrams(args ...int) {
	for k, v := range args {
		fmt.Printf("%v, %v\n", k, v)
	}
}

func returnVal() (isOk bool, val int) {
	return true, 10
}

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

func switchTest() {
	// every case has a break at the end
	i := 1
	switch i {
	case 1:
		fmt.Printf("This is 1")
		fallthrough
	case 2:
		fmt.Printf("This is 2")

	case 3:
		fmt.Printf("This is 3")
	default:
		fmt.Printf("This is default")
	}
}

func gotoTest() {
	for i := 0; i < 10; i++ {
		fmt.Printf("%d\n", i)
		if i == 5 {
			goto end
		}
	}
end:
	fmt.Printf("This is end")
}

func returnTest(val int) (bool, int) {
	resultBool := false
	if val > 0 {
		resultBool = true
	}
	fmt.Printf("bool=%v, int=%v", resultBool, val)
	return resultBool, val
}
