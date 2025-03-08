from typing import List


# @leet start
class Solution:
    def rob(self, nums: List[int]) -> int:
        @cache
        def dfs(i):
            if i < 0:
                return 0
            return max(dfs(i - 2) + nums[i], dfs(i - 1))

        return dfs(len(nums) - 1)


# @leet end

