from typing import List
from functools import cache


# @leet start
class Solution:
    def rob(self, nums: List[int]) -> int:
        n = len(nums)
        if n == 1:
            return nums[0]

        @cache
        def dfs2(i, j):  # 1- n-1
            if i < j:
                return 0
            return max(dfs2(i - 2, j) + nums[i], dfs2(i - 1, j))

        return max(dfs2(n - 2, 0), dfs2(n - 1, 1))


# @leet end
class Solution2:
    def rob(self, nums: List[int]) -> int:
        n = len(nums)
        if n == 1:
            return nums[0]

        @cache
        def dfs1(i):  # 0- n-2
            if i < 0:
                return 0
            if i == 0:
                return nums[0]
            return max(dfs1(i - 2) + nums[i], dfs1(i - 1))

        @cache
        def dfs2(i):  # 1- n-1
            if i < 1:
                return 0
            if i == 1:
                return nums[1]
            return max(dfs2(i - 2) + nums[i], dfs2(i - 1))

        return max(dfs1(n - 2), dfs2(n - 1))
