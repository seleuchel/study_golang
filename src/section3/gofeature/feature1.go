// go 특징 1

package main

import "fmt"

func main() {
	//go는 모호하거나 애매한 문법을 금지
	//후치 연산자는 허용하나, 전치 연산자는 비허용(++i 불가)
	//증감연산 반환 값 없음!
	//포인터(pointer)를 허용한다. 그러나 연산은 비허용

	//예제
	sum, i := 0, 0

	for i <= 100 {
		// sum += i++
		sum += i
		i++

	}
	fmt.Println("ex1 : ", sum)
}
