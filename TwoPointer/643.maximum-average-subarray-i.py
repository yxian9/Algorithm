from typing import List


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

