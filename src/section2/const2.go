package main

import "fmt"

func main() {
	const a, b int = 1, 2
	const c, d, e = true, 0.23, "ff"

	//여러개 선언
	const (
		x, y int16 = 50, 90
		i, k       = "did", 22
	)

	fmt.Println(a, b, c, d, e)
	fmt.Println(x, y, i, k)

}
