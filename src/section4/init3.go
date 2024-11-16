package main

import (
	"fmt"
	"section4/lib"
)

var num int32

func init() {
	num = 30
	fmt.Println("Main init start - main")
}

func main() {
	fmt.Println("10보다 큰수 ", lib.CheckNum(11))
}
