package sorting

import (
	"math"
	"math/rand"
	"time"
)

// QuickSort
// Randomize pivot
// Randomization nullifies adverse effect of badly arranged input permutation
// Algorithm sees the input as a random permutation.
// Probability that almost all random choices are the worst is nearly 1/n!
// FACT: Randomized QuickSort will take O(n log n) time with probability very close to 1 on every input.
func QuickSort(array []int) []int {
	// Write your code here.
	if len(array) < 2 {
		return array
	}
	rand.Seed(time.Now().UnixNano())
	pivot := array[rand.Intn(len(array))]
	lessThanPivot, pivotArray, moreThanPivot := []int{}, []int{}, []int{}
	for _, value := range array {
		if value < pivot {
			lessThanPivot = append(lessThanPivot, value)
		} else if value == pivot {
			pivotArray = append(pivotArray, value)
		} else {
			moreThanPivot = append(moreThanPivot, value)
		}
	}
	return append(append(QuickSort(lessThanPivot), pivotArray...), QuickSort(moreThanPivot)...)
}

// BubbleSort
// O(n^2)
func BubbleSort(array []int) []int {
	// Write your code here.
	result := array
	for i := len(array) - 1; i >= 0; i-- {
		for j := 0; j < i; j++ {
			if result[j] > result[j+1] {
				tmp := result[j]
				result[j] = result[j+1]
				result[j+1] = tmp
			}
		}
	}
	return result
}

// InsertionSort
// O(n^2)
func InsertionSort(array []int) []int {
	// Write your code here.
	result := array
	for i := 0; i < len(array); i++ {
		for j := 0; j < i; j++ {
			if result[j] > result[i] {
				tmp := result[j]
				result[j] = result[i]
				result[i] = tmp
			}
		}
	}
	return result
}

// SelectionSort
// O(n^2)
func SelectionSort(array []int) []int {
	// Write your code here.
	result := array
	for i := 0; i < len(array); i++ {
		minimum := math.MaxInt32
		minimumIndex := -1
		for j := i; j < len(array); j++ {
			if minimum > result[j] {
				minimum = result[j]
				minimumIndex = j
			}
		}
		tmp := result[minimumIndex]
		result[minimumIndex] = result[i]
		result[i] = tmp
	}
	return result
}

