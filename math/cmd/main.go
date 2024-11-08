package calculator

import (
	"math/rand/v2"
)

func Cal(val1 int, val2 int) int {
	val3 := rand.IntN(6)
	return val1 + val2 + val3
}
