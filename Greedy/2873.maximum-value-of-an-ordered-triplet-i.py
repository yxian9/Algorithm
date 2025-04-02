from typing import List, Optional
from collections import defaultdict


# @leet start
class Solution:
    def maximumTripletValue(self, nums: List[int]) -> int:
        ret = 0
        n = len(nums)
        left = nums[0]
        for j in range(1, n):
            if left <= nums[j]:
                left = nums[j]
                continue
            for k in range(j + 1, n):
                ret = max(ret, (left - nums[j]) * nums[k])
        return ret


# @leet end

