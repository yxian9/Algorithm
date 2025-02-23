from typing import List


# @leet start
class Solution:
    def twoSum(self, numbers: List[int], target: int) -> List[int]:
        l, r = 0, len(numbers) - 1
        while l < r:
            total = numbers[l] + numbers[r]
            if total == target:
                break
            if total > target:
                r -= 1
            else:
                l += 1
        return [l + 1, r + 1]


# @leet end

