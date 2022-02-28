package main

import (
	"fmt"
	"strconv"
)

type Student struct {
	Id   uint64
	Name string
	Age  uint8
}

func main() {
	s := "2"
	i, err := strconv.ParseInt(s, 2, 64)
	if err == nil {
		fmt.Println(i)
	} else {
		fmt.Println(err)
	}

}
