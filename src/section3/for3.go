package main

import "fmt"

func main() {
	fmt.Println("1")

	// 레이블 접근 - 루프 흐름 제어 (Loop 사용)

	// 예제 1
Loop1:
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if i == 2 && j == 4 {
				break Loop1 // 이렇게되면 Loop1가 있는 가장 바깥으로 빠져나가게됨
			}
			fmt.Println("ok ", i, j)
		}
	}

	// 예제2
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Println("ext2 : ", i)
	}

Loop2: // loop 뒤에는 반드시 for문이 와야함. 아니면 break든 무언가 등
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if i == 1 && j == 2 {
				continue Loop2
			}
			fmt.Println("ex3  : ", i, j)
		}
	}

}
