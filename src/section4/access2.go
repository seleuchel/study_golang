package main

import (
	"fmt"
	checkUp "section4/lib" // 앞에 alias 별칭을 넣을 수 있다.
	_ "section4/lib2"      // _ 앞에 넣으면 사용하지 않아서 패키지가 지워지는 현상을 막을 수 있다. 일단 import 넣기
)

func main() {
	// 별칭 사용
	// 빈 식별자 사용
	fmt.Println("10보다 큰 수 : ", checkUp.CheckNum(11))

}
