from typing import List

def bubble_sort(numbers: List[int]) -> List[int]:
    """
    Sorts a list of integers using the Bubble Sort algorithm.
    Time Complexity: O(n^2)
    Space Complexity: O(1)
    """
    n: int = len(numbers)
    
    # Outer loop to traverse through all array elements
    for i in range(n):
        # Last i elements are already in place, so we ignore them
        for j in range(0, n - i - 1):
            # Compare the element at j with the next element at j + 1
            if numbers[j] > numbers[j + 1]:
                # Swap if the element found is greater than the next element
                numbers[j], numbers[j + 1] = numbers[j + 1], numbers[j]
                
    return numbers

if __name__ == "__main__":
    # Test Case 1: Standard unsorted list
    assert bubble_sort([64, 34, 25, 12, 22, 11, 90]) == [11, 12, 22, 25, 34, 64, 90]
    
    # Test Case 2: Already sorted list
    assert bubble_sort([1, 2, 3, 4, 5]) == [1, 2, 3, 4, 5]
    
    # Test Case 3: Empty list (Edge case)
    assert bubble_sort([]) == []
    
    # Test Case 4: List with one element
    assert bubble_sort([42]) == [42]
    
    print("All Python test cases passed successfully!")