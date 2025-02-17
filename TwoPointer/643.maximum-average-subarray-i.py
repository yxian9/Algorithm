from typing import List
from math import inf

# @leet start
class Solution:
    def findMaxAverage(self, nums: List[int], k: int) -> float:
        cur = avg = sum(nums[:k])
        for i in range(k, len(nums)):
            cur += nums[i]
            cur -= nums[i - k]
            if cur > avg:
                avg = cur
        return float(avg) / float(k)


# @leet end

class Solution:
    def findMaxAverage(self, nums: List[int], k: int) -> float:
        total = 0
        for i, v in enumerate(nums):  # NOTE: miss last item
            total += v
            if i == k - 1:
                best = total  # NOTE: wield condition
            if i >= k:
                total -= nums[i - k]
                best = max(total, best)
        return float(best) / float(k)

class Solution:
    def findMaxAverage(self, nums: List[int], k: int) -> float:
        total = 0
        best = None
        for i, v in enumerate(nums):
            if i < k:  # NOTE: incorrect
                total += v
            else:
                if best is None or total > best:
                    best = total
                total -= nums[i - k]
                total += v
        return float(best) / float(k)
