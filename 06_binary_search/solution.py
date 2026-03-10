from typing import List

def binary_search(arr: List[int], target: int)-> int:

    left: int = 0
    right: int = len(arr) - 1

    while left <= right:
        mid = left + (right - left) // 2
        if arr[mid] == target:
            return mid
        elif arr[mid] > target:
            right = mid - 1
        elif arr[mid] < target:
            left = mid + 1
        
    return -1

if __name__ == "__main__":
    # Happy path
    assert binary_search([-1, 0, 3, 5, 9, 12], 9) == 4
    
    # Target not found (simulando uma query miss)
    assert binary_search([-1, 0, 3, 5, 9, 12], 2) == -1 
    
    # Empty list edge case (cluster vazio)
    assert binary_search([], 5) == -1
    
    # Single element list
    assert binary_search([5], 5) == 0
    print("All tests passed! Production ready.")