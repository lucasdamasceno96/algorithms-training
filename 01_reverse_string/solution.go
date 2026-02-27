package main

// reverseString takes a string and returns it reversed.
// It handles multi-byte characters correctly by using runes.
func ReverseString(s string) string {
	// Convert string to a slice of runes (UTF-8 safe)
	runes := []rune(s)
	
	left, right := 0, len(runes)-1
	
	for left < right {
		// Swap elements in the slice
		runes[left], runes[right] = runes[right], runes[left]
		
		// Update pointers
		left++
		right--
	}
	
	// Convert the slice of runes back to a string
	return string(runes)
}

