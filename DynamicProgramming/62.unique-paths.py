# @leet start
class Solution:
    def uniquePaths(self, m: int, n: int) -> int:
        pre = [1] * n
        cur = [1] * n

        for _ in range(1, m):
            for c in range(1, n):
                cur[c] = cur[c - 1] + pre[c]
            pre, cur = cur, pre

        return pre[-1]  # NOTE: last return is pre due to flip


# @leet end

