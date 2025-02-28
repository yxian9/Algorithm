from functools import cache


# @leet start
class Solution:
    def minInsertions(self, s: str) -> int:
        @cache
        def dfs(i, j):
            if i > j or i == j:
                return 0
            if s[i] == s[j]:
                return dfs(i + 1, j - 1)
            else:
                return min(dfs(i, j - 1), dfs(i + 1, j)) + 1

        return dfs(0, len(s) - 1)


# @leet end
