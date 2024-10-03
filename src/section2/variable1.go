package main

import "fmt"

func main() {
	// 변수특 : 숫자 첫글자x, 대소문자구분, 특수기호사용가능(한글 변수 가능)
	// 변수 및 상수. 함수 내외 사용가능
	// 변수 선언만하고 사용 안하면 에러를 출력함..ㄴ
	// 값을 대입하지 않으면 0으로 초기화, string은 null로 초기화됨

	var a int
	var b string
	var c, d, e int
	var f, g, h int = 1, 2, 3
	var i float32 = 11.3
	var j string = "hi, sele"
	var k = 4.74
	var l = "hihi"
	var m = true

	a = 4
	b = "didi"

	fmt.Println("hello")
	fmt.Println("a : ", a)
	fmt.Println("b : ", b)
	fmt.Println("c : ", c)
	fmt.Println("d : ", d)
	fmt.Println("e : ", e)
	fmt.Println("f : ", f)
	fmt.Println("g : ", g)
	fmt.Println("h : ", h)
	fmt.Println("i : ", i)
	fmt.Println("j : ", j)
	fmt.Println("k : ", k)
	fmt.Println("l : ", l)
	fmt.Println("m : ", m)

}
