from typing import List


# @leet start
class Solution:
    def maxArea(self, height: List[int]) -> int:
        l, r = 0, len(height) - 1
        ret = 0
        while l < r:
            area = min(height[l], height[r]) * (r - l)
            ret = max(area, ret)
            if height[l] < height[r]:
                l += 1
            else:
                r -= 1
        return ret


# @leet end
