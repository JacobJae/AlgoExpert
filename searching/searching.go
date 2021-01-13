package searching

import "math"

// FindThreeLargestNumbers
// Find Three Largest Numbers in array
func FindThreeLargestNumbers(array []int) []int {
	// Write your code here.
	result := []int{math.MinInt32, math.MinInt32, math.MinInt32}
	for _, i := range array {
			if result[2] < i {
				result[0] = result[1]
				result[1] = result[2]
				result[2] = i
			} else if result[1] < i {
				result[0] = result[1]
				result[1] = i
			} else if result[0] < i {
				result[0] = i
			}
	}
	return result
}

// BinarySearch
func BinarySearch(array []int, target int) int {
	return BinarySearchHelper(array,target,0,len(array)-1)
}

// BinarySearchHelper
func BinarySearchHelper(array []int, target, start, end int) int {
	index := (end + start) / 2
	if end < start {
		return -1
	}
	if array[index] == target {
		return index
	} else if target < array[index]{
		return  BinarySearchHelper(array, target, start, index-1)
	} else {
		return  BinarySearchHelper(array, target, index+1, end)
	}
}
