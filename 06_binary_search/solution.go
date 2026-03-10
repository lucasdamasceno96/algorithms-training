package main

// BinarySearch finds the index of a target integer in a sorted slice.
// It returns -1 if the target is not found.
func BinarySearch(arr []int, target int) int {
	left := 0
	right := len(arr) - 1

	for left <= right {
		// Calculate mid safely to avoid integer overflow
		mid := left + (right-left)/2

		if arr[mid] == target {
			return mid // Target found, route the request here
		} else if arr[mid] < target {
			left = mid + 1 // Discard the entire left partition
		} else {
			right = mid - 1 // Discard the entire right partition
		}
	}

	return -1 // Target not found in the cluster
}
