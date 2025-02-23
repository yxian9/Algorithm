from typing import List


# @leet start
class Solution:
    def minEatingSpeed(self, piles: List[int], h: int) -> int:
        l, r = 1, max(piles)
        res = 0

        def t(n, k):
            return n // k if n % k == 0 else n // k + 1

        while l <= r:
            speed = (l + r) // 2
            total_hours = sum(t(pile, speed) for pile in piles)

            if total_hours > h:
                l = speed + 1
            else:
                res = speed
                r = speed - 1

        return res


# @leet end

