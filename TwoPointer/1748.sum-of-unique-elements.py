from collections import Counter
from typing import List


# @leet start
class Solution:
    def sumOfUnique(self, nums: List[int]) -> int:
        s = Counter(nums)
        res = 0
        for k, v in s.items():
            if v == 1:
                res += k
        return res


# @leet end

