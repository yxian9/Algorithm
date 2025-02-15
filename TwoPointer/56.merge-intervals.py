from typing import List


# @leet start
class Solution:
    def merge(self, intervals: List[List[int]]) -> List[List[int]]:
        if len(intervals) <= 1:
            return intervals
        intervals.sort()

        def overlap(a, b):
            l, r = max(a[0], b[0]), min(a[1], b[1])
            return l <= r

        res = [intervals[0]]
        for inter in intervals[1:]:
            if overlap(inter, res[-1]):
                res[-1][1] = max(res[-1][1], inter[1])
            else:
                res.append(inter)
        return res


# @leet end

