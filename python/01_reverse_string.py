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
    sample: str = "platform"
    print(f"Original: {sample} | Reversed: {reverse_string(sample)}")