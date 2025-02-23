from typing import List
import heapq


# @leet start
class Solution:
    def kClosest(self, points: List[List[int]], k: int) -> List[List[int]]:
        res = []
        for i in range(len(points)):
            x, y = points[i]
            dis = -(x * x + y * y)
            if i < k:
                heapq.heappush(res, (dis, x, y))
            else:
                heapq.heappushpop(res, (dis, x, y))

        return [[x, y] for _, x, y in res]


# @leet end

