package main

// BubbleSort receives a slice of integers and returns it sorted.
// It uses the in-place bubble sort logic.
func BubbleSort(numbers []int) []int {
	n := len(numbers)

	for i := 0; i < n; i++ {
		// Optimization: elements after n-i-1 are already sorted
		for j := 0; j < n-i-1; j++ {
			if numbers[j] > numbers[j+1] {
				// Swap using Go's multiple assignment syntax
				numbers[j], numbers[j+1] = numbers[j+1], numbers[j]
			}
		}
	}

	return numbers
}

func main() {
	// The main function can be empty or used for simple manual execution.
	// Testing logic resides in solution_test.go.
}
