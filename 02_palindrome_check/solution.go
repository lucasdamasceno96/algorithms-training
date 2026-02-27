package main

import (
	"fmt"
	"strings"
)

// IsPalindrome checks string consistency from both ends.
func IsPalindrome(s string) bool {
	// Normalizing to lowercase to avoid case sensitivity issues
	s = strings.ToLower(s)
	runes := []rune(s)
	left, right := 0, len(runes)-1

	for left < right {
		if runes[left] != runes[right] {
			return false
		}
		left++
		right--
	}
	return true
}

func main() {
	sample := "Level"
	result := IsPalindrome(sample)
	fmt.Printf("Word: %s | Is Palindrome: %t\n", sample, result)
}