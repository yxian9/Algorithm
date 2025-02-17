from typing import List


# @leet start
class Solution:
    def longestOnes(self, nums: List[int], k: int) -> int:
        l = res = count = 0
        for r, v in enumerate(nums):
            if v == 0:
                count += 1
            while count > k:
                if nums[l] == 0:
                    count -= 1
                l += 1
            cur_length = r - l + 1
            res = max(res, cur_length)

        return res


# @leet end

