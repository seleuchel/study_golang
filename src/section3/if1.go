package main

import "fmt"

func main() {
	// 조건문
	// if는 무조건 true else false (1, 0 안됨)
	// 소괄호 사용 안함

	var a int = 20
	b := 20

	if a >= 15 { // 중괄호는 반드시 if문과 동일한 줄에 쓰기!
		fmt.Println("15이상") // 중괄호 생략하면 안됨. python과 다름
	}
	if b >= 25 {
		fmt.Println("25이상")
	}

	// 에러발생!
	/*
		if c:= 1; c{  // 사용불가
			fmt.//
		}

		// 에러발생!
		if c := 40; c >= 35 {
			fmt.Println("35이상")
		}
		c += 20 //에러!


	*/

}
