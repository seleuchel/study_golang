package main

import "fmt"

func main() {
	/// 예2ㅔ 1

	a := 30 / 15
	switch a {
	case 2, 4, 6:
		fmt.Println(a, "짝수")
	case 1, 3, 5:
		fmt.Println(a, "홀수")

	}

	// 예제2. go에서는 스위치에 break 없어도 break 된다.
	// fallthrough : 그래서 break없이 다음 스위치로 넘어가고 싶을 때 사용한다.(맨 마지막 case는 쓰지 않는다.)
	switch e := "go"; e {
	case "go":
		fmt.Println("go")
		fallthrough
	case "python":
		fmt.Println("python")
		fallthrough
	case "ruby":
		fmt.Println("ruby")
		fallthrough
	case "c++":
		fmt.Println("c++")
		fallthrough
	case "rust":
		fmt.Println("rust")
	}

}
