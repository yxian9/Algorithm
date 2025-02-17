from typing import List


# @leet start
class Solution:
    def nextPermutation(self, nums: List[int]) -> None:
        """
        Do not return anything, modify nums in-place instead.
        """
        # find pivot
        n, p = len(nums), -1
        for i in range(n - 2, -1, -1):
            if nums[i] < nums[i + 1]:
                p = i
                break
        if p == -1:
            return nums.reverse()
        # find successor  and swap
        for i in range(n - 1, p, -1):
            if nums[i] > nums[p]:
                nums[i], nums[p] = nums[p], nums[i]
                break
        ## reverse suffix
        l, r = p + 1, n - 1
        while l < r:
            nums[l], nums[r] = nums[r], nums[l]
            l += 1
            r -= 1


# @leet end

