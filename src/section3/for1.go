package main

func main() {
	// 반복문 for
	// go에서 유일하게 제공되는 반복문
	//

	// 예제1
	for i := 0, i < 5 ,i++ { // 중괄호는 반드시 뒤에 붙어야!
		fmt.Println("ex1 : ", i)
	}

	// 에러
	/*
	while 문

	무한루프 형태
	for {
		fmt.Println("무한루프")
	}

	*/

	// 예제3 
	loc := []string{"Seoul", "busan", "incheon"} 

	for index, name := range loc{ //배열을 순회하면서 인덱스와 값을 가져온다
		fmt.Println("ex3 : ", index, name)
	}
	

}
