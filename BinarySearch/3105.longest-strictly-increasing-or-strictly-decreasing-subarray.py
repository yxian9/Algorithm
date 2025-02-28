from typing import List


# @leet start
class Solution:
    def longestMonotonicSubarray(self, nums: List[int]) -> int:
        ret = up = down = 1

        for i in range(1, len(nums)):
            if nums[i] == nums[i - 1]:
                up = down = 1
            elif nums[i] > nums[i - 1]:
                up += 1
                down = 1
                ret = max(ret, up)
            else:
                down += 1
                up = 1
                ret = max(ret, down)

        return ret


# @leet end

