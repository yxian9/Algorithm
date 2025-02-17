from typing import List


# @leet start
class Solution:
    def subarraySum(self, nums: List[int], k: int) -> int:
        total = res = 0
        preSum = {0: 1}
        for v in nums:
            total += v
            diff = total - k
            if diff in preSum:
                res += preSum[diff]
            preSum[total] = 1 + preSum.get(total, 0)
        return res


# @leet end

