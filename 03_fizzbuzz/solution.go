package main

import (
	"strconv"
)

func FizzBuzz(n int) []string{
	// FizzBuzz processes numbers from 1 to n and returns their string representations
	// based on divisibility by 3 and 5.
	// Pre-allocating capacity to avoid multiple memory reallocations
	results := make([]string, 0, n)
	for i := 1; i <= n; i++ {
		if i%15 == 0 {
			results = append(results, "FizzBuzz")
		} else if i%3 == 0 {
			results = append(results, "Fizz")
		} else if i%5 == 0 {
			results = append(results, "Buzz")
		} else {
			// strconv.Itoa is the idiomatic way to convert int to string
			results = append(results, strconv.Itoa(i))
		}
	}

	return results
}

