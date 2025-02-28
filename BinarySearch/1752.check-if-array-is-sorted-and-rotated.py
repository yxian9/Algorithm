from typing import List


# @leet start
class Solution:
    def check(self, nums: List[int]) -> bool:
        n = len(nums)
        count = 0
        for i in range(1, n):
            if nums[i - 1] > nums[i]:
                count += 1
        if nums[0] < nums[n - 1]:
            count += 1
        return count <= 1


# @leet end

