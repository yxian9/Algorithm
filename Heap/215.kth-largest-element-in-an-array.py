from typing import List, Optional
from collections import defaultdict

# @leet start
from heapq import heappush, heappushpop, heappop


class Solution:
    def findKthLargest(self, nums: List[int], k: int) -> int:
        minheap = []
        for i in range(k):
            heappush(minheap, nums[i])
        for i in range(k, len(nums)):
            heappushpop(minheap, nums[i])
        return heappop(minheap)


# @leet end

