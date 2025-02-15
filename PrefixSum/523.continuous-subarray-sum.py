from typing import List


# @leet start
class Solution:
    def checkSubarraySum(self, nums: List[int], k: int) -> bool:
        preSum = 0
        m = {0: -1}
        for i, v in enumerate(nums):
            preSum += v
            remain = preSum % k
            if remain in m:
                if i - m[remain] >= 2:
                    return True
            else:
                m[remain] = i

        return False


# @leet end

