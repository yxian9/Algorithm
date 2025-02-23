from typing import List


# @leet start
class Solution:
    def productExceptSelf(self, nums: List[int]) -> List[int]:
        n = len(nums)
        ret = [1] * n
        pre = post = 1

        for i in range(n):
            ret[i] *= pre
            pre *= nums[i]

        for i in range(n - 1, -1, -1):
            ret[i] *= post
            post *= nums[i]

        return ret


# @leet end

