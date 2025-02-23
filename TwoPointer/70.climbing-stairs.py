# @leet start
class Solution:
    def climbStairs(self, n: int) -> int:
        pre, cur = 1, 1
        for _ in range(n):
            cur, pre = pre + cur, cur

        return pre


# @leet end

