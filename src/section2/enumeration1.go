package main

import "fmt"

func main() {
	// 열거형 - 상수를 사용하는 일정한 규칙에 따라 숫자를 계산 / 증가시키는 묶음  iota(0부터 시작)
	const (
		Jan = 1
		Feb = 2
		Mar = 3
	)

	fmt.Println(Jan)
	fmt.Println(Feb)
	fmt.Println(Mar)

}
