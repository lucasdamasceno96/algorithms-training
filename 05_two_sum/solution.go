package main

// TwoSum searches for two numbers in a slice that sum to target.
// It returns a slice containing the two indices.
func TwoSum(nums []int, target int) []int {
	// prevMap stores the value as key and its index as value.
	// make() pre-allocates the map for efficiency.
	prevMap := make(map[int]int)

	for i, n := range nums {
		diff := target - n

		// The 'comma-ok' idiom to check if the complement exists in the map
		if idx, ok := prevMap[diff]; ok {
			return []int{idx, i}
		}

		// Store the current number's index for future lookups
		prevMap[n] = i
	}

	return nil
}
