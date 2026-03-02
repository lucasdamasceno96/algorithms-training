// solution_test.go
package main

import "testing"

func TestFindMax(t *testing.T) {
	tests := []struct {
		name    string
		input   []int
		want    int
		wantErr bool
	}{
		{"positive numbers", []int{1, 5, 8, 2}, 8, false},
		{"negative numbers", []int{-1, -5, -2}, -1, false},
		{"empty slice", []int{}, 0, true},
		{"single element", []int{100}, 100, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := FindMax(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("FindMax() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("FindMax() got = %v, want %v", got, tt.want)
			}
		})
	}
}
