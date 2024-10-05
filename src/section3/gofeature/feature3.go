// fmt 사용방법

package main

import "fmt"

func main() {
	// 코드 서식 지정
	// 코드의 일관성을 맞춰줌
	// gofmt -h : 사용법
	// gofmt -w : 원본파일에 반영 . 오.

	// 예제1
	for i := 0; i <= 100; i++ {
		fmt.Println("ex :", i)
	}
}
