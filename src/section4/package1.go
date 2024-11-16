package main

// 선언
import (
	"fmt"
	"os"
)

//

func main() {
	// 패키지 : 코드 구조화(모듈) 및 재사용
	// 응집도, 결합도
	// 패키지 단위의 독립적이고 작은 단위로 개발 -> 작은 패키지로 프로그램을 작성
	// 패키지 이름 = 디렉토리이름
	// 같은 패킨지 내 -> 소스파일들은 디렉토리 명을 패키지명르로 사용한다.
	// 소문자는 private, 대문자는 public
	// fmt의 위치는 'ge env' go root 경로의 pkg 안에있다. fmt는 디렉토리이다.
	// main 패키지는 특별하다. -> 프로그램 시작점으로 인식한다. 패키지가 아님.
	var name string
	fmt.Println("이름 : ")
	fmt.Scanf("%s", &name)
	fmt.Fprintf(os.Stdout, "hi %s", name)

}
