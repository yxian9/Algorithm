from typing import List


# @leet start
class Solution:
    def minSubArrayLen(self, target: int, nums: List[int]) -> int:
        l = minL = total = 0
        for r in range(len(nums)):
            v = nums[r]
            total += v
            while total >= target:
                curL = r - l + 1
                if minL == 0 or curL < minL:
                    minL = curL
                d = nums[l]
                l += 1
                total -= d

        return minL


# @leet end


class Solution2:
    def minSubArrayLen(self, target: int, nums: List[int]) -> int:
        l = ret = 0
        total = 0
        for r, v in enumerate(nums):
            total += v
            while total >= target:
                if ret == 0:
                    ret = r - l + 1
                else:
                    ret = min(ret, r - l + 1)
                total -= nums[l]
                l += 1
        return ret
