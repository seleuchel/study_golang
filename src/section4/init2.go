package main

import (
	"fmt"
)

func init() {
	// 메인 함수가 없느 경우 다른 main 함수를 가지고 있는 src에서 이 패키지를 호출하면, 가장 먼저 실행됨
	fmt.Println("here, init")
}

// go 패키지 동작 순서
// main -> import -> import const var init -> import const var init 순서...
