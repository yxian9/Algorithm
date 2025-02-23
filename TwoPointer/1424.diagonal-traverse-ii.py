from collections import defaultdict
from typing import List


# @leet start
class Solution:
    def findDiagonalOrder(self, nums: List[List[int]]) -> List[int]:
        res = []
        n = 0
        diags = defaultdict(list)
        for i in range(len(nums)):
            row = nums[i]
            for j in range(len(row)):
                diags[i + j].append(row[j])
                n = max(n, i + j)

        for i in range(n + 1):
            row = diags[i]
            row.reverse()
            res.extend(row)
        return res


# @leet end

