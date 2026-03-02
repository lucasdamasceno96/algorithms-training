from typing import List


def find_max_element(numbers: List[int]) -> int:
    """
    Scans a list of integers to find the highest value.
    Analogous to finding the peak CPU usage in a series of metrics.
    """
    if not numbers:
        raise ValueError("The list of numbers is empty")
    current_max: int = numbers[0]
    for num in numbers[1:]:
        if num > current_max:
            current_max = num
    return current_max


# -----------------------
# Inline Tests
# -----------------------
if __name__ == "__main__":
    # Test Case 1: Standard input
    assert find_max_element([1, 5, 3, 2]) == 5, "Failed on normal list [1,5,3,2]"
    
    # Test Case 2: List with negative numbers
    assert find_max_element([-10, -5, -2]) == -2, "Failed on negative numbers"
    
    # Test Case 3: Single element list
    assert find_max_element([42]) == 42, "Failed on single element list"
    
    # Test Case 4: All elements equal
    assert find_max_element([7, 7, 7, 7]) == 7, "Failed on all elements equal"
    
    # Test Case 5: Mixed negative and positive numbers
    assert find_max_element([-1, 0, 1]) == 1, "Failed on mixed negative and positive numbers"
    
    # Test Case 6: Empty list should raise ValueError
    try:
        find_max_element([])
        assert False, "Failed to raise ValueError on empty list"
    except ValueError as e:
        assert str(e) == "The list of numbers is empty", "Wrong error message for empty list"

    print("All inline tests for find_max_element passed!")