package main

import "fmt"

func main() {
	// 중요!! : go에서만 있는 문법. 짧은 선언 - 1회성(사용 끝나면 메모리에서 비워짐)
	// 반드시 함수 안에서만 사용. 전역으로 사용 불가
	// 한번 선언 후 재할당 시 예외 발생
	// 제한된 범위의 함수내에서 사용할경우 코드 가독성을 높일 수 있음

	shortVar1 := 3
	shortVar2 := "Test"
	shortVar3 := false

	// if문, 반복문에서 사용 ex. var i int = 10
	fmt.Println("shortVar1 : ", shortVar1, "shortVar2 : ", shortVar2, "shortVar3 : ", shortVar3)

	// 진짜 독특하다
	// i는 if에서 쓰고 없어짐
	if i := 10; i < 11 {
		fmt.Println("short variable tst success!")
	}

}
