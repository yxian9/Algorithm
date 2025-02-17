from typing import List
from random import randint


# @leet start
class Solution:
    def __init__(self, w: List[int]):
        total = 0
        preSum = [0] * len(w)
        for i, v in enumerate(w):
            total += v
            preSum[i] = total
        self.total = total
        self.preSum = preSum

    def pickIndex(self) -> int:
        weight = randint(1, self.total)
        l, r, idx = 0, len(self.preSum) - 1, 0
        while l <= r:
            mid = (l + r) // 2
            if self.preSum[mid] >= weight:
                idx = mid
                r = mid - 1
            else:
                l = mid + 1
        return idx


# Your Solution object will be instantiated and called as such:
# obj = Solution(w)
# param_1 = obj.pickIndex()
# @leet end

