package main

import (
	"fmt"
)

func init() {
	// 메인보다 먼저
	// 패키지 로드시에 가장 먼저 호출.
	fmt.Println("init method start")
}

func main() {
	fmt.Println("start")
}
