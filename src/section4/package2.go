package main

import (
	"fmt"
	"section4/lib" // go mod init
)

func main() {
	fmt.Println("10보다 큰수 : ", lib.CheckNum(15)) // 대문자이므로 접근제어가 Public 임 ㄴ
}
