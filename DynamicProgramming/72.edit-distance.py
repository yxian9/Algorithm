# @leet start
from functools import cache


class Solution:
    def minDistance(self, word1: str, word2: str) -> int:
        n1, n2 = len(word1), len(word2)

        @cache
        def dfs(i, j):
            if i == -1:  # NOTE: note 0 base is -1
                return j + 1
            if j == -1:
                return i + 1
            if word1[i] == word2[j]:
                return dfs(i - 1, j - 1)

            return min(
                dfs(i - 1, j) + 1,
                dfs(i, j - 1) + 1,
                dfs(i - 1, j - 1) + 1,
            )

        return dfs(n1 - 1, n2 - 1)


# @leet end

