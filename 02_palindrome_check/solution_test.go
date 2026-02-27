package main

import "testing"

func TestIsPalindrome(t *testing.T) {
    tests := []struct {
        name     string
        input    string
        expected bool
    }{
        {"Valid Palindrome", "Racecar", true},
        {"Invalid Palindrome", "DevOps", false},
        {"Single character", "a", true},
    }

    for _, tc := range tests {
        t.Run(tc.name, func(t *testing.T) {
            // Note: sua função deve se chamar IsPalindrome (I maiúsculo para ser visível)
            result := IsPalindrome(tc.input)
            if result != tc.expected {
                t.Errorf("For %s: expected %v, got %v", tc.input, tc.expected, result)
            }
        })
    }
}