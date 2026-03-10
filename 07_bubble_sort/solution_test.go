package main

import (
	"reflect"
	"testing"
)

func TestBubbleSort(t *testing.T) {
	// Define the test table structure
	tests := []struct {
		name     string
		input    []int
		expected []int
	}{
		{
			name:     "Standard Unsorted List",
			input:    []int{64, 34, 25, 12, 22, 11, 90},
			expected: []int{11, 12, 22, 25, 34, 64, 90},
		},
		{
			name:     "Already Sorted",
			input:    []int{1, 2, 3, 4, 5},
			expected: []int{1, 2, 3, 4, 5},
		},
		{
			name:     "Empty Slice",
			input:    []int{},
			expected: []int{},
		},
		{
			name:     "Single Element",
			input:    []int{42},
			expected: []int{42},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Executing the algorithm
			result := BubbleSort(tt.input)

			// DeepEqual is used to compare slices content
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("BubbleSort() = %v, want %v", result, tt.expected)
			}
		})
	}
}
