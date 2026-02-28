package main

import (
	"reflect"
	"testing"
)

func TestFizzBuzz(t *testing.T) {
	// Defining our test cases (The Table)
	tests := []struct {
		name     string
		input    int
		expected []string
	}{
		{"Small input", 3, []string{"1", "2", "Fizz"}},
		{"Five elements", 5, []string{"1", "2", "Fizz", "4", "Buzz"}},
		{"FizzBuzz check", 15, []string{"1", "2", "Fizz", "4", "Buzz", "Fizz", "7", "8", "Fizz", "Buzz", "11", "Fizz", "13", "14", "FizzBuzz"}},
		{"Check 20 ends with Buzz", 20, nil},
		{"Zero input", 0, []string{}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FizzBuzz(tt.input)
			// reflect.DeepEqual is used to compare slices
			if !reflect.DeepEqual(got, tt.expected) {
				t.Errorf("FizzBuzz(%d) = %v; want %v", tt.input, got, tt.expected)
			}
		})
	}
}