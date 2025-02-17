from typing import List


# @leet start
class Solution:
    def removeDuplicates(self, nums: List[int]) -> int:
        l = 1
        for r in range(1, len(nums)):
            if nums[r] == nums[r - 1]:
                continue
            nums[l] = nums[r]
            l += 1
        return l


# @leet end


class Solution:
    def removeDuplicates(self, nums: List[int]) -> int:
        l = 1
        for r in range(1, len(nums)):
            if nums[r] != nums[r - 1]:
                nums[l] = nums[r]
                l += 1
        return l


def removeDuplicates(self, nums: List[int]) -> int:
    l = 0
    for v in nums:
        if nums[l] != v:
            nums[l] = v
            l += 1
    return l

