from typing import List

# @leet start
class Solution:
    def twoSum(self, nums: List[int], target: int) -> List[int]:
        m = {}
        for i, v in enumerate(nums):
            diff = target - v
            if diff in m:
                return [m[diff], i]
            m[v] = i
        return []
# @leet end
