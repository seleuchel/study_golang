package main

import "fmt"

func main() {
	// 상수
	// const 사용 초기화, 한 번 선언 후 값 변경금지, 고정된 값 관리용
	// 상수는 선언과 동시에 초기화.
	const a string = "test1"
	const b = "test2"
	const c = 10 * 10
	const d = 10 // func1()  // 함수리턴을 const로 넣으면 에러남
	const e = 323.1
	const f = false

	fmt.Println("a : ", a, "b : ", b, "c : ", c, "d : ", d, "f : ", f)

}
