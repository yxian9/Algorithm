from typing import List


# @leet start
class Solution:
    def findKthPositive(self, arr: List[int], k: int) -> int:
        l, r, idx = 0, len(arr) - 1, -1
        while l <= r:
            mid = (l + r) // 2
            if arr[mid] - mid - 1 < k:
                idx = mid
                l = mid + 1
            else:
                r = mid - 1

        return idx + k + 1


# @leet end

