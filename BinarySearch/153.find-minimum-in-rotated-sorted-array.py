from typing import List, Optional
from collections import defaultdict


# @leet start
class Solution:
    def findMin(self, nums: List[int]) -> int:
        ret = nums[0]
        l, r = 0, len(nums) - 1
        while l <= r:
            mid = l + (r - l) // 2
            ret = min(nums[mid], ret)
            if nums[l] <= nums[mid]:
                ret = min(nums[l], ret)
                l = mid + 1
            else:
                ret = min(nums[mid + 1], ret)
                r = mid - 1

        return ret


# @leet end
def findMin(nums: List[int]) -> int:
    ret = nums[0]
    l, r = 0, len(nums) - 1
    while l <= r:
        mid = l + (r - l) // 2
        ret = min(nums[mid], ret)
        if nums[l] < nums[mid]:
            ret = min(nums[l], ret)
            l = mid + 1
        else:
            ret = min(nums[mid + 1], ret)
            r = mid - 1

    return ret
