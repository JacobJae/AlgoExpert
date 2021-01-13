package strings

import (
	"strconv"
)

// IsPalindrome
// Check if given string is palindrome
func IsPalindrome(str string) bool {
	// Write your code here.
	result := true
	start := 0
	end := len(str) - 1
	for start < end {
		if str[start] != str[end] {
			result = false
		}
		start++
		end--
	}
	return result
}

// CaesarCipherEncryptor
// shift alphabet key times to right
func CaesarCipherEncryptor(str string, key int) string {
	// Write your code here.
	result := ""
	for _, chard := range str {
		result = result + string((int(chard) - 97 + key)%26 + 97)
	}
	return result
}

// RunLengthEncoding
// Print same character with number but not exceeding 9
func RunLengthEncoding(str string) string {
	// Write your code here.
	result := ""
	targetRune := int32(str[0])
	count := 0
	for _, curRune := range str {
		if curRune == targetRune {
			count++
		} else {
			result += strconv.Itoa(count)
			result += string(targetRune)
			targetRune = curRune
			count = 1
		}

		if count > 9 {
			result += strconv.Itoa(9)
			result += string(targetRune)
			count = 1
		}
	}
	result += strconv.Itoa(count)
	result += string(targetRune)
	return result
}