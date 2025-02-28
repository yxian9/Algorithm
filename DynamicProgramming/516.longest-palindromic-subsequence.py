# @leet start
from functools import cache


class Solution:
    def longestPalindromeSubseq(self, s: str) -> int:
        @cache
        def dfs(i, j):
            if i > j:
                return 0
            if i == j:
                return 1
            if s[i] == s[j]:
                return dfs(i + 1, j - 1) + 2
            else:
                return max(dfs(i, j - 1), dfs(i + 1, j))

        return dfs(0, len(s) - 1)


# @leet end
