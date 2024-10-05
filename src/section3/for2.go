package main

import "fmt"

func main() {
	sum1 := 0
	for i := 0; i <= 100; i++ {
		sum1 += i //가산
	}
	fmt.Println("ex1:", sum1)

	// 예제2

	sum2, i := 0, 0

	for i <= 100 {
		sum2 += 1
		i++

		// i = i++에 추가하면  에러남 (go에서 후치연산 i++는 반환값이 없음)
	}
	fmt.Println("ex2 :", sum2)

	// 에제3 - while사용시 이렇게 사용
	sum3, i := 0, 0

	for {
		if i > 100 {
			break
		}
		sum3 += i
		i++
	}
	fmt.Println("Ex3, ", sum3)

	//예제 4'
	for i, j := 0, 0; i <= 10; i, j = i+1, j+10 {
		fmt.Println("ex4 : ", i, j)
		//for 문에서 대입가능함
	}

	// 에러 발생 - 수치연산은 반환값이 없음(i++)
	/*
		for i, j := 0, 0 ; i <=10 ; i++, j+= 10{

		}
	*/

}
