from typing import List


# @leet start
class Solution:
    def maxSubArray(self, nums: List[int]) -> int:
        best = total = nums[0]
        for i in range(1, len(nums)):
            item = nums[i]
            # if total + item > total:
            # if item + total > item:  # NOTE: chose total!!
            if total > 0:  # NOTE: chose total!!
                total += item
            else:
                total = item
            best = max(best, total)
        return best


# @leet end

