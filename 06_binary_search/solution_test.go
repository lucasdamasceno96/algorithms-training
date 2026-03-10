package main

import "testing"

func TestBinarySearch(t *testing.T) {
	// 1. Define the table of tests
	tests := []struct {
		name     string
		arr      []int
		target   int
		expected int
	}{
		// Happy path: Normal routing request, element exists.
		{"happy_path", []int{-1, 0, 3, 5, 9, 12}, 9, 4},

		// Target not found: Simulating a 404 Not Found or a query miss.
		{"target_not_found", []int{-1, 0, 3, 5, 9, 12}, 2, -1},

		// Empty array: Querying a cluster with zero running pods. Must not panic.
		{"empty_array", []int{}, 5, -1},

		// Single element found: A cluster with only one node, and it is the target.
		{"single_element_found", []int{5}, 5, 0},

		// Single element not found: A cluster with one node, but it's the wrong one.
		{"single_element_not_found", []int{5}, 10, -1},
	}

	// 2. Iterate over the table
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := BinarySearch(tt.arr, tt.target)
			if result != tt.expected {
				t.Errorf("got %d, want %d", result, tt.expected)
			}
		})
	}
}
