def is_palindrome(text: str) -> bool:
    """Checks if palindrome using two pointers (Case insensitive)."""
    left, right = 0, len(text) - 1
    while left < right:
        if text[left].lower() != text[right].lower():
            return False
        left += 1
        right -= 1
    return True

if __name__ == "__main__":
    assert is_palindrome("Radar") is True, "Failed mixed case"
    assert is_palindrome("python") is False, "Failed non-palindrome"
    assert is_palindrome(" ") is True, "Failed space"
    
    print("02_palindrome_check (Python): All tests passed!")