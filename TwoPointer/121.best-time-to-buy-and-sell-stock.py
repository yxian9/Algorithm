from typing import List


# @leet start
class Solution:
    def maxProfit(self, prices: List[int]) -> int:
        res, low = 0, prices[0]
        for v in prices:
            low = min(v, low)
            res = max(v - low, res)

        return res


# @leet end

