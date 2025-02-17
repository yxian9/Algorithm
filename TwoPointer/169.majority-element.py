from typing import List


# @leet start
class Solution:
    def majorityElement(self, nums: List[int]) -> int:
        major = nums[0]
        count = 0
        for i in nums:
            if i == major:
                count += 1
            else:
                count -= 1
            if count < 0:
                count = 0
                major = i
        return major


# @leet end

