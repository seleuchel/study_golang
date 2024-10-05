package main

import "fmt"

func main() {
	// 제어문(조건문)
	//switch 뒤 표현식 생략가능
	// case뒤 표현식(expression) 사용가능
	//자동 break 때문에 fallthrouth whswo
	// type 분기 -> 값이 아닌 변수 type으로 분기가능

	a := -7
	switch {
	case a < 0:
		fmt.Println(a, "는 음수")
	case a == 0:
		fmt.Println(a, "는 0")
	case a > 0:
		fmt.Println(a, "는 양수")

	}

	//예제
	switch b := 27; {
	case b < 0:
		fmt.Println(a, "는 음수")
	case b == 0:
		fmt.Println(a, "는 0")
	case b > 0:
		fmt.Println(a, "는 양수")

	}

	// 예제 3
	switch c := "go"; c { // switch c하고 아예 위에서 한번 선언 (c를 추가했음). 그러면 아래는 값만 들어가도됨
	case "go":
		fmt.Println("Go")
	case "java":
		fmt.Println("Java")
	default:
		fmt.Println("일치없음")
	}

	// 예제 4
	switch c := "go"; c + "lang" {
	case "golang":
		fmt.Println("1")
	case "Java":
		fmt.Println("java")
	default:
		fmt.Println("None")

	}

	// 또 다른 예제 5

	switch i, j := 20, 30; { // 스위치에서 변수 선언시 ; 뒤에 붙여줘야한다.
	case i < j:
		fmt.Println("None1")
	case i == j:
		fmt.Println("None2")
	case i > j:
		fmt.Println("None3")
	}

}
