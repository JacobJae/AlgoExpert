package recursion

import (
	"math"
)

type SpecialArray []interface{}

// Tip: Each element of `array` can either be cast
// to `SpecialArray` or `int`.
func ProductSum(array []interface{}) int {
	// Write your code here.
	return -1
}


func RecursiveDoubling(base int, exponent int) int {
	if exponent == 0 {
		return 1
	}
	doubled := base * base
	result := RecursiveDoubling(doubled, int(math.Floor(float64(exponent/2))))
	if exponent % 2 == 1 {
		result = result * base
	}
	return result
}

// a b
// c d
// [0][0] = a
// [0][1] = b
// [1][0] = c
// [1][1] = d
func FibMatrix(exponent int) [2][2]int {
	newBase := [2][2]int{}
	if exponent == 1 {
		newBase[0][0] = 1
		newBase[0][1] = 1
		newBase[1][0] = 1
		newBase[1][1] = 0
		return newBase
	} else {
		newBase = FibMatrix(exponent - 1)
		newBase2 := [2][2]int{}
		newBase2[0][0] = newBase[0][0] + newBase[1][0]
		newBase2[0][1] = newBase[0][1] + newBase[1][1]
		newBase2[1][0] = newBase[0][0]
		newBase2[1][1] = newBase[0][1]
		return newBase2
	}
}

func GetNthFib(n int) int {
	// Write your code here.
	return FibMatrix(n)[1][1]
}