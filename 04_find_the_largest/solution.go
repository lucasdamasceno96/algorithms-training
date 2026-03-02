// solution.go
package main

import (
	"errors"
)

// FindMax receives a slice of integers and returns the largest one.
// This is a linear O(n) operation.
func FindMax(nums []int) (int, error) {
	if len(nums) == 0 {
		return 0, errors.New("cannot find maximum of an empty slice")
	}

	maxVal := nums[0]

	for _, val := range nums {
		if val > maxVal {
			maxVal = val
		}
	}

	return maxVal, nil
}
