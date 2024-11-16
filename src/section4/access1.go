package main

import (
	"fmt"
	"section4/lib2"
)

func main() {
	// 변수 상숭 함수 메서드 구조체 등 식별자
	// 소문자 명명은 패키지 내에서만 접근 가능
	fmt.Println("100보다 큰 수 : ", lib2.CheckNum1(101))
	fmt.Println("100보다 큰 수 : ", lib2.CheckNum2(101))
}
