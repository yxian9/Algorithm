from typing import List, Optional
from collections import defaultdict


# @leet start
class Solution:
    def maximumTripletValue(self, nums: List[int]) -> int:
        n = len(nums)
        pre_max = nums[0]
        max_diff = 0
        ret = 0
        for k in range(1, n):
            ret = max(ret, max_diff * nums[k])
            pre_max = max(pre_max, nums[k])
            max_diff = max(max_diff, pre_max - nums[k])
        return ret


# @leet end

