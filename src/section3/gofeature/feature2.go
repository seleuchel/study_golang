package main

import "fmt"

func main() {
	// 문장 끝 세미콜론 주의 ;
	// 자동으로 끝에 세미콜론 삽입
	// 두 문장을 한 문장으로 쓸 떄 ;로 구분

	// 예제 1
	for i := 0; i <= 10; i++ {
		// fmt.Println("ex1 : ", i);fmt.Println(i)

	}

	//예제 2
	for j := 10; j >= 10; j-- {
		fmt.Println("ex2 ", j)
	}
}
