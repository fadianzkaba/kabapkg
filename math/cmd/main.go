package main

import (
	"fmt"
	"math/rand/v2"
)

func main() {
	result := cal(1, 2)
	fmt.Println(result)
}

func cal(val1 int, val2 int) int {
	val3 := rand.IntN(6)
	return val1 + val2 + val3
}
