def two_sum(nums: list[int], target: int) -> list[int]:
    """
    Finds two indices such that their values sum up to the target.
    This implementation uses a hash map to achieve O(n) time complexity.
    """
    # Map to store: { value: index }
    # Thinking of it as a localized cache for O(1) lookups.
    prev_map = {} 

    for i, n in enumerate(nums):
        diff = target - n
        
        # If the complement is in our 'cache', we are done
        if diff in prev_map:
            return [prev_map[diff], i]
        
        # Otherwise, store the current value and its index
        prev_map[n] = i
        
    # According to constraints, we'll always find a solution.
    return []

if __name__ == "__main__":
    # Test case: Standard match
    assert two_sum([2, 7, 11, 15], 9) == [0, 1]
    # Test case: Match with numbers at the end
    assert two_sum([3, 2, 4], 6) == [1, 2]
    
    print("Python Two Sum: Success!")