from typing import List
from collections import Counter

# @leet start
from heapq import heappush, heappushpop


class Solution:
    def topKFrequent(self, nums: List[int], k: int) -> List[int]:
        freq = Counter(nums)
        minH = []
        for v, count in freq.items():
            if len(minH) < k:
                heappush(minH, (count, v))
            else:
                heappushpop(minH, (count, v))

        return [v for _, v in minH]


# @leet end

