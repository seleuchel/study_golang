package main

import "fmt"

func main() {
	// 열거형 - 상수를 사용하는 일정한 규칙에 따라 숫자를 계산 / 증가시키는 묶음  iota(0부터 시작)
	const (
		_ = iota * 2 // 생략, 일정한 규칙을 작성하면 뒤의 내용에도 반영됨.  @ 문자열은?
		A
		_ // 생략
		C
	)

	fmt.Println(A)
	fmt.Println(B)
	fmt.Println(C)

}
