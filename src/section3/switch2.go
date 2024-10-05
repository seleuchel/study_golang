package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	switch i := rand.Intn(100); {
	case i >= 50 && i < 100:
		fmt.Println(i, "50 이상 100미만")
	case i >= 25 && i < 50:
		fmt.Println(i, "25이상 50미만")
	default:
		fmt.Println(i, "기본")
	}

}
