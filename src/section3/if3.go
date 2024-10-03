package main

import "fmt"

func main() {
	i := 100

	if i >= 120 {
		fmt.Println("d1")
	} else if i >= 100 && i < 120 {
		fmt.Println("d2")
	} else if i < 100 && i >= 50 {
		fmt.Println("d3")
	}
}
