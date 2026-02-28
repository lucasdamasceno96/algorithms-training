from typing import List

def fizz_buzz(n: int) -> List[str]:
    """
    Produces a list of strings from 1 to n with FizzBuzz logic.
    """
    results: List[str] = []
    
    for i in range(1, n + 1):
        if i % 15 == 0:
            results.append("FizzBuzz")
        elif i % 3 == 0:
            results.append("Fizz")
        elif i % 5 == 0:
            results.append("Buzz")
        else:
            # Fallback: conversion of the integer to string
            results.append(str(i))
    
    return results

# Unit Testing Culture
if __name__ == "__main__":
    # Test Case 1: Standard small input
    # Checking if the last element of fizz_buzz(3) is "Fizz"
    assert fizz_buzz(3) == ["1", "2", "Fizz"], "Failed on n=3"
    
    # Test Case 2: Checking FizzBuzz at 15
    result_15 = fizz_buzz(15)
    assert result_15[-1] == "FizzBuzz", "Failed on n=15 (FizzBuzz)"
    assert result_15[4] == "Buzz", "Failed on n=15 (Buzz at index 4)"
    
    # Test Case 3: Edge Case n=0
    assert fizz_buzz(0) == [], "Failed on edge case n=0"
    
    # Test Case 4: Performance/Boundary
    assert len(fizz_buzz(100)) == 100, "Failed length check for n=100"

    # Test Case 5:
    result_20 = fizz_buzz(20)
    assert result_20[-1] == "Buzz", "Failed on n=20 (Buzz at 20)"
    
    print("03_fizz_buzz (Python): All tests passed!")