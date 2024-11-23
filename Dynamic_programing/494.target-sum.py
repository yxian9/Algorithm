from typing import List
from functools import cache


# @leet start
class Solution:
    def findTargetSumWays(self, nums: List[int], target: int) -> int:
        @cache
        def dfs(level, cur):
            if level == len(nums):
                return 1 if cur == target else 0
            l = dfs(level + 1, cur + nums[level])
            r = dfs(level + 1, cur - nums[level])
            return l + r

        return dfs(0, 0)


# @leet end
class Solution2:
    def findTargetSumWays(self, nums: List[int], target: int) -> int:
        total = sum(nums) + target
        if abs(target) > sum(nums):
            return 0
        if total % 2 == 1:
            return 0
        total = total // 2
        dp = [1] + [0] * total
        for item in nums:
            for j in range(len(dp) - 1, item - 1, -1):
                # if dp[j - item] > 0: ## no nessary, but helpful to keep for df
                dp[j] += dp[j - item]
        return dp[-1]

