from typing import List
from typing import Optional


# @leet start
class Solution:
    def trap(self, height: List[int]) -> int:
        n = len(height)
        ret = 0
        l, r = 0, n - 1
        lmax = rmax = 0
        while l <= r:
            if height[l] < height[r]:
                lmax = max(lmax, height[l])
                ret += lmax - height[l]
                l += 1
            else:
                rmax = max(rmax, height[r])
                ret += rmax - height[r]
                r -= 1
        return ret


# @leet end
def trap2(height):
    n = len(height)
    l_highest = [0] * n
    r_highest = [0] * n
    temp = 0
    for i in range(n):
        temp = max(height[i], temp)
        l_highest[i] = temp
    temp = 0
    for i in range(n - 1, -1, -1):
        temp = max(height[i], temp)
        r_highest[i] = temp

    ret = 0
    for i in range(n):
        ret += min(l_highest[i], r_highest[i]) - height[i]
    return ret


def trap(height: List[int]) -> int:
    n = len(height)
    l_highest = [0] * n
    r_highest = [0] * n
    l_highest[0] = height[0]
    r_highest[n - 1] = height[n - 1]

    for i in range(1, n):
        l_highest[i] = max(height[i], l_highest[i - 1])
    for i in range(n - 2, -1, -1):
        r_highest[i] = max(height[i], r_highest[i + 1])

    ret = 0
    for i in range(n):
        ret += min(l_highest[i], r_highest[i]) - height[i]
    return ret
