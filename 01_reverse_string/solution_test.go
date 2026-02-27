package main

import "testing"

func TestReverseString(t *testing.T) {
    tests := []struct {
        name     string
        input    string
        expected string
    }{
        {"Standard string", "cloud", "duolc"},
        {"Single char", "g", "g"},
        {"Empty string", "", ""},
    }

    for _, tc := range tests {
        t.Run(tc.name, func(t *testing.T) {
            result := ReverseString(tc.input)
            if result != tc.expected {
                t.Errorf("For %s: expected %s, got %s", tc.input, tc.expected, result)
            }
        })
    }
}