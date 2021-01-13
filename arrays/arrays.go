package arrays

import (
	"AlgoExpert/sorting"
	"fmt"
	"math"
)

// FirstDuplicateValue
// 2 1 5 2 3 3 4 = 2
// 2 1 5 3 3 2 4 = 3
func FirstDuplicateValue(array []int) int {
	// Write your code here.
	bigVal := false
	for _, value := range array {
		if value < 0 {
			value *= -1
		}
		if value == len(array) {
			if bigVal {
				return value
			} else {
				bigVal = true
			}
		} else {
			if array[value] < 0 {
				return value
			} else {
				array[value] = array[value] * -1
			}
		}
	}
	return -1
}


// ArrayOfProducts
// index i should have products of all elements except for i
func ArrayOfProducts(array []int) []int {
	// Write your code here.
	product := 1
	zeroCounter := 0
	for _, value := range array {
		if value == 0 {
			zeroCounter++
		} else {
			product *= value
		}
	}
	for i, value := range array {
		if zeroCounter < 1 {
			array[i] = product / value
		} else if zeroCounter > 1 {
			array[i] = 0
		} else {
			if value == 0 {
				array[i] = product
			} else {
				array[i] = 0
			}
		}
	}
	return array
}

func LongestPeak(array []int) int {
	longestPeakLength := 0
	i := 1
	for i < len(array)-1 {
		isPeak := array[i-1] < array[i] && array[i] > array[i+1]
		if !isPeak {
			i += 1
			continue
		}

		// Stretch left index
		leftIdx := i - 2
		for leftIdx >= 0 && array[leftIdx] < array[leftIdx+1] {
			leftIdx -= 1
		}

		// Stretch right index
		rightIdx := i + 2
		for rightIdx < len(array) && array[rightIdx] < array[rightIdx-1] {
			rightIdx += 1
		}

		// get current length of peak
		currentPeakLength := rightIdx - leftIdx - 1

		// update longest peak
		if currentPeakLength > longestPeakLength {
			longestPeakLength = currentPeakLength
		}
		// We know that before rightIdx, it was decreasing, jump to rightIdx
		i = rightIdx
	}
	return longestPeakLength
}



// LongestPeak
// Peak: increasing and decreasing
// 1, 4, 10, 2 is peak because 1 4 10 increase and 2 decrease
// 4, 0, 10		not peak
// 1, 2, 2, 0	not peak
// 1, 2, 3		not peak
// at least 3 requires to form peak
func LongestPeakFailed(array []int) int {
	// Write your code here.
	increasing := false
	decreasing := false
	inPeak := false
	suffixStart, suffixEnd := -1, -1
	curretStart, currentEnd := -1, -1
	if len(array) == 0 {
		return 0
	}
	for i, value := range array {
		fmt.Println("Current i:", i, "value:", value)
		if i == 0 {
			continue
		}
		// Get peak
		if increasing {
			// peak was increasing and keep increasing, update end index
			if value > array[i-1] {
				currentEnd = i
			} else if value < array[i-1] {
				// peak was increasing now its decreasing
				currentEnd = i
				increasing = false
				decreasing = true
				inPeak = true
			} else {
				// Same value, it is not peak, refresh indices to initial value
				increasing = false; decreasing = false
				curretStart = -1; currentEnd = -1
			}
		} else if decreasing {
			if value > array[i-1] {
				// was it peak decreasing? or just decreasing
				if inPeak {
					// end of peak, compare with suffix and update longest one
					if currentEnd - curretStart > suffixEnd - suffixStart {
						suffixStart = curretStart
						suffixEnd = currentEnd
					}
					// Now it is increasing
					inPeak = false
				// We do nothing if it was not in peak
				}
					decreasing = false
					increasing = true
					curretStart = i - 1
					currentEnd = i
			} else if value < array[i-1] {
				currentEnd = i
			} else {
				curretStart = -1; currentEnd = -1
				increasing = false; decreasing = false
			}
		} else {
			if value > array[i-1] {
				increasing = true
				curretStart = i-1
				currentEnd = i
			} else if value < array[i-1] {
				decreasing = true
			}
		}
		fmt.Println("curStart:", curretStart, "curEnd:",currentEnd,"sufStart:",suffixStart,"sufEnd:",suffixEnd)
		// Compare suffix with prefix
	}
	if suffixStart == -1 && suffixEnd == -1 {
		return currentEnd - curretStart + 1
	} else {
		return suffixEnd - suffixStart + 1
	}
}


// SpiralTraverse
// [1,	2,	3,	4]
// [12,	13,	14,	5]
// [11,	16,	15,	6]
// [10,	9,	8,	7]
// -> [1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16]
// Idea: Perform outside first and recurse through inside
func SpiralTraverse(array [][]int) []int {
	// Write your code here.
	result := []int{}
	dimensionRow := len(array[0])
	dimensionColoum := len(array)
	// If one of these two is 1, print all
	if dimensionColoum < 2 {
		singleArray := []int{}
		for i := 0; i < dimensionRow; i++ {
			singleArray = append(singleArray, array[0][i])
		}
		return singleArray
	} else if dimensionRow < 2 {
		singleArray := []int{}
		for i := 0; i < dimensionColoum; i++ {
			singleArray = append(singleArray, array[i][0])
		}
		return singleArray
	}
	// Get first line (left to right)
	for i := 0; i < dimensionRow; i++ {
		result = append(result, array[0][i])
	}
	// Get right side line (top to bottom)
	for i := 1; i < dimensionColoum; i++ {
		result = append(result, array[i][dimensionRow-1])
	}
	// Get bottom side line (right to left)
	for i := dimensionRow - 2; i > -1; i-- {
		result = append(result, array[dimensionColoum-1][i])
	}
	// Get left side line (bottom to top)
	for i := dimensionColoum - 2; i > 0; i-- {
		result = append(result, array[i][0])
	}
	newArray := [][]int{}
	for i := 1; i < dimensionColoum - 1; i++ {
		newArray = append(newArray, array[i][1:dimensionRow-1])
	}
	if len(newArray) < 1 || len(newArray[0]) < 1 {
		return result
	} else {
		return append(result, SpiralTraverse(newArray)...)
	}
}

// IsMonotonic
// Check if array is monotonic
func IsMonotonic(array []int) bool {
	// Write your code here.
	result := true
	nonDecrease := false
	nonIncrease := false
	for idx, _ := range array {
		// Skip first element
		if idx == 0 {
			continue
		}
		diff := array[idx - 1] - array[idx]
		if nonDecrease {
			if  diff > 0 {
				result = false
			}
		} else if nonIncrease {
			if  diff < 0 {
				result = false
			}
		} else {
			if  diff > 0 {
				nonIncrease = true
			} else if diff < 0 {
				nonDecrease = true
			}
		}
	}
	return result
}


// MoveElementToEnd
// Move target number to back off array
// Must be done in place
// other number's order doesn't matter as long as target number is in back
// 2, 1, 2, 2, 2, 3, 4, 2 -> 4, 1, 3, 2, 2, 2, 2, 2
func MoveElementToEnd(array []int, toMove int) []int {
	// Write your code here.
	targetIndex := -1
	// We shouldn't update index if it is multiple target
	multipleTarget := false
	for idx, value := range array {
		// If current value is target
		if value == toMove {
			// Check if it is multiple target
			if multipleTarget == false {
				// We only update targetIndex at first
				// Later, it will be update incrementally by 1
				if targetIndex == -1 {
					targetIndex = idx
				}
				multipleTarget = true
			}
		} else {
			multipleTarget = false
			if targetIndex > -1 {
				// Swap target and non target and increment targetIndex
				array[targetIndex] = value
				array[idx] = toMove
				targetIndex++
			}
		}
	}
	return array
}


// SmallestDifference
// pick 1 from each array and make pair, return smallest difference pair
// first array element should be in first position of return array
func SmallestDifference(array1, array2 []int) []int {
	// Write your code here.
	sortedArray1 := sorting.QuickSort(array1)
	sortedArray2 := sorting.QuickSort(array2)
	minDifference := math.MaxInt32
	result := []int{}
	arrayIndex1, arrayIndex2 := 0, 0
	for arrayIndex1 < len(array1) && arrayIndex2 < len(array2) {
		// Get difference
		difference := sortedArray1[arrayIndex1] - sortedArray2[arrayIndex2]
		// Get absolute value
		if difference < 0 {
			difference *= -1
		}
		// Compare with current minimum and update minimum difference
		if minDifference > difference {
			minDifference = difference
			result = []int{sortedArray1[arrayIndex1], sortedArray2[arrayIndex2]}
		}
		// Increment index
		// if one of index is at the end, increament other
		if arrayIndex1 > len(array1) - 2 {
			arrayIndex2++
		} else if arrayIndex2 > len(array2) - 2 {
			arrayIndex1++
		} else {
			// Now both indices are lower than max index
			diff1 := sortedArray1[arrayIndex1 + 1] - sortedArray2[arrayIndex2]
			if diff1 < 0 {
				diff1 *= -1
			}
			diff2 := sortedArray1[arrayIndex1] - sortedArray2[arrayIndex2 + 1]
			if diff2 < 0 {
				diff2 *= -1
			}
			if diff1 < diff2 {
				arrayIndex1++
			} else {
				arrayIndex2++
			}
		}
	}
	return result
}

// ThreeNumberSum
// Find 3 numbers with target sum
// eg array: [12, 3, 1, 2, -6, 5, -8, 6]
// eg target: 0
// eg answer: [[- 8, 2, 6], [-8, 3, 5], [-6, 1, 5]]
func ThreeNumberSum(array []int, target int) [][]int {
	// Write your code here.
	var result [][]int
	sortedArray := sorting.QuickSort(array)
	for i := 0; i < len(sortedArray); i++ {
		left := i + 1
		right := len(sortedArray) - 1
		for left < right {
			sum := sortedArray[i] + sortedArray[left] + sortedArray[right]
			if target == sum {
				result = append(result, []int{sortedArray[i], sortedArray[left], sortedArray[right]})
				left++
				right--
			} else if target > sum {
				left++
			} else {
				right--
			}
		}
	}
	return result
}

// IsValidSubsequence
// Check if sequence if subsequence of array
// eg array: [5, 1, 22, 25, 6, -1, 8, 10]
// eg sequence: [1, 6, -1, 10]
// eg answer: true because array contains [1, 6, -1, 10] in order
// Idea: Iterate sequentially and check
func IsValidSubsequence(array []int, sequence []int) bool {
	// Write your code here.
	var (
		arrayIndex = 0
		sequenceIndex = 0
		result = false
	)
	for sequenceIndex < len(sequence) {
		for arrayIndex < len(array) {
			//fmt.Println("Sequence: ", sequence[sequenceIndex], ", Array: ", array[arrayIndex])
			if sequence[sequenceIndex] == array[arrayIndex] {
				if sequenceIndex == len(sequence) - 1 {
					result = true
				}
				arrayIndex++
				break
			}
			arrayIndex++
		}
		sequenceIndex++
	}
	return result
}


// TwoNumberSum
// Return pair of numbers with target sum
// Idea: Use go map. map in Go is hashmap so get value takes O(1)
//		We just need to create map from array and start looking
//		for match which will take O(n) time
func TwoNumberSum(array []int, target int) []int {
	// Write your code here.
	result := []int{}
	tmp := make(map[int]int)
	for _, tmpint := range array {
		tmp[tmpint] = tmpint
	}
	for _, firstMatch := range array {
		secondMatch, found := tmp[target - firstMatch]
		if found == true && firstMatch != secondMatch {
			result = append(result, firstMatch, secondMatch)
			break
		}
	}
	return result
}