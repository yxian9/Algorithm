from collections import defaultdict
from typing import List

# @leet start
from random import randint


class Solution:
    def __init__(self, nums: List[int]):
        m: defaultdict[int, list[int]] = defaultdict(list)
        for i, v in enumerate(nums):
            m[v].append(i)
        self.m = m

    def pick(self, target: int) -> int:
        idx = self.m[target]
        n = randint(0, len(idx) - 1)
        return idx[n]


# Your Solution object will be instantiated and called as such:
# obj = Solution(nums)
# param_1 = obj.pick(target)
# @leet end

