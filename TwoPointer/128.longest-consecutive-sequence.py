from typing import List
from collections import Counter


# @leet start
class Solution:
    def longestConsecutive(self, nums: List[int]) -> int:
        ret, freq = 0, Counter(nums)

        for x in freq:
            if x - 1 in freq:
                continue
            y = x + 1
            while y in freq:
                y += 1
            ret = max(ret, y - x)

        return ret


# @leet end

