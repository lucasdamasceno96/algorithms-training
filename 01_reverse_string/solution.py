from typing import List

def reverse_string(s: str) -> str:
    """
    Reverses a string in-place using the two-pointers technique.
    Time Complexity: O(n)
    Space Complexity: O(n) due to string immutability in Python.
    """
    # Convert string to list because Python strings are immutable
    chars: List[str] = list(s)
    
    left: int = 0
    right: int = len(chars) - 1
    
    while left < right:
        # Swap characters idiomatic way
        chars[left], chars[right] = chars[right], chars[left]
        
        # Move pointers towards the middle
        left += 1
        right -= 1
        
    return "".join(chars)

# Basic test
if __name__ == "__main__":
    # Test cases
    assert reverse_string("platform") == "mroftalp", "Failed common string"
    assert reverse_string("a") == "a", "Failed single character"
    assert reverse_string("") == "", "Failed empty string"
    
    print("01_reverse_string (Python): All tests passed!")