from typing import List


# @leet start
class Solution:
    def ovlp(self, a, b):
        l, r = max(a[0], b[0]), min(a[1], b[1])
        if l <= r:
            return [l, r]
        return []

    def intervalIntersection(
        self, firstList: List[List[int]], secondList: List[List[int]]
    ) -> List[List[int]]:
        i1 = i2 = 0
        res = []
        while i1 < len(firstList) and i2 < len(secondList):
            b1, b2 = firstList[i1], secondList[i2]
            inter = self.ovlp(b1, b2)
            if inter:
                res.append(inter)
            if b1[1] <= b2[1]:
                i1 += 1
            else:
                i2 += 1

        return res


# @leet end

