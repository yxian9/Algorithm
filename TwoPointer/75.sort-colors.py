from typing import List


# @leet start
class Solution:
    def sortColors(self, nums: List[int]) -> None:
        """
        Do not return anything, modify nums in-place instead.
        """
        l, r = 0, len(nums) - 1
        i = 0
        while i < len(nums):
            v = nums[i]
            if v == 0 and l < i:
                nums[i], nums[l] = nums[l], nums[i]
                l += 1
                continue
            if v == 2 and i < r:
                nums[i], nums[r] = nums[r], nums[i]
                r -= 1
                continue
            i += 1


# @leet end

